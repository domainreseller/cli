package dna

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ErrorKind is a coarse, stable classification of API failures. Programs and
// AI agents can branch on it without parsing messages. The CLI maps each kind
// to a documented exit code.
type ErrorKind string

const (
	KindAuth        ErrorKind = "auth"         // bad/missing credentials, unknown or suspended reseller
	KindForbidden   ErrorKind = "forbidden"    // authenticated but not allowed
	KindNotFound    ErrorKind = "not_found"    // resource does not exist
	KindValidation  ErrorKind = "validation"   // the API rejected the input or a business rule failed
	KindConflict    ErrorKind = "conflict"     // state conflict
	KindRateLimited ErrorKind = "rate_limited" // too many requests (retry after RetryAfter)
	KindServer      ErrorKind = "server"       // 5xx; may be transient
	KindUnknown     ErrorKind = "unknown"
)

// APIError is returned for every non-2xx response.
type APIError struct {
	StatusCode int `json:"status"`
	// Code is the API's error code when it sends one (for example
	// "Dna.DomainService:Domain:1018"), otherwise a synthetic code such as
	// INVALID_API_KEY, INVALID_RESELLER, RATE_LIMITED or HTTP_404.
	Code             string            `json:"code"`
	Message          string            `json:"message"`
	Details          string            `json:"details,omitempty"`
	ValidationErrors []ValidationError `json:"validationErrors,omitempty"`
	CorrelationID    string            `json:"correlationId,omitempty"`
	RetryAfter       time.Duration     `json:"-"`
	Body             []byte            `json:"-"`
}

// ValidationError is one field-level validation failure.
type ValidationError struct {
	Message string   `json:"message"`
	Members []string `json:"members,omitempty"`
}

func (e *APIError) Error() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d %s: %s", e.StatusCode, e.Code, e.Message)
	for _, v := range e.ValidationErrors {
		sb.WriteString("; ")
		if len(v.Members) > 0 {
			sb.WriteString(strings.Join(v.Members, ",") + ": ")
		}
		sb.WriteString(v.Message)
	}
	return sb.String()
}

// internalErrorCode is the API's generic "An internal error has occurred"
// code. The API sometimes sends it with a 404 status, which must not be read
// as "resource does not exist".
const internalErrorCode = "Dna.DomainService:General:10001"

// Kind classifies the error.
func (e *APIError) Kind() ErrorKind {
	switch {
	case e.Code == internalErrorCode:
		return KindServer
	case e.StatusCode == http.StatusUnauthorized, e.StatusCode == http.StatusFound, e.StatusCode == http.StatusMovedPermanently:
		return KindAuth
	case e.StatusCode == http.StatusForbidden:
		return KindForbidden
	case e.StatusCode == http.StatusNotFound:
		return KindNotFound
	case e.StatusCode == http.StatusConflict:
		return KindConflict
	case e.StatusCode == http.StatusTooManyRequests:
		return KindRateLimited
	case e.StatusCode == http.StatusBadRequest, e.StatusCode == http.StatusUnprocessableEntity:
		return KindValidation
	case e.StatusCode >= 500:
		return KindServer
	}
	return KindUnknown
}

// Retryable reports whether repeating the same request later may succeed.
func (e *APIError) Retryable() bool {
	k := e.Kind()
	return k == KindRateLimited || (k == KindServer && e.StatusCode >= 500 && e.StatusCode != http.StatusNotImplemented)
}

// Hint returns a short, actionable suggestion for the error ("" if none).
func (e *APIError) Hint() string {
	switch e.Code {
	case "INVALID_RESELLER":
		return "The reseller ID is unknown, inactive or suspended in this environment. Check DNA_RESELLER_ID (a UUID from the panel's API settings) and that you are using the right environment (--env production|ote)."
	case "INVALID_API_KEY":
		return "Production and OTE use different API keys. Make sure the key belongs to the selected environment (--env), or run `dna auth login` again."
	case "MISSING_API_KEY", "MISSING_RESELLER_ID":
		return "Run `dna auth login`, or set DNA_RESELLER_ID and DNA_API_KEY."
	case "SUBRESELLER_NOT_ALLOWED":
		return "Sub-reseller accounts cannot use the API. Use the main reseller's credentials."
	case "RATE_LIMITED":
		return "The gateway allows 5 requests/second per key (2/s for domain search). Slow down and retry."
	}
	switch e.Kind() {
	case KindNotFound:
		return "Check the spelling and the environment (--env production|ote): each environment has its own domains. `dna domain list` shows what this account holds."
	case KindValidation:
		if len(e.ValidationErrors) > 0 {
			return "Fix the listed fields and retry. `dna schema <operation>` shows every field and its constraints."
		}
	case KindServer:
		if e.CorrelationID != "" {
			return "Server-side failure. If it persists, contact support and quote correlation id " + e.CorrelationID + "."
		}
	}
	return ""
}

// IsKind reports whether err is an *APIError of the given kind.
func IsKind(err error, kind ErrorKind) bool {
	var ae *APIError
	return errors.As(err, &ae) && ae.Kind() == kind
}

// IsNotFound reports whether err is a 404 from the API.
func IsNotFound(err error) bool { return IsKind(err, KindNotFound) }

// IsAuth reports whether err is an authentication failure.
func IsAuth(err error) bool { return IsKind(err, KindAuth) }

// NetworkError wraps transport failures (DNS, TLS, timeouts, resets).
type NetworkError struct {
	Op  string
	Err error
}

func (e *NetworkError) Error() string { return fmt.Sprintf("%s: %v", e.Op, e.Err) }
func (e *NetworkError) Unwrap() error { return e.Err }

// Timeout reports whether the failure was a timeout.
func (e *NetworkError) Timeout() bool {
	var ne net.Error
	return errors.As(e.Err, &ne) && ne.Timeout()
}

// abpBoilerplate matches the generic .NET exception message the API sometimes
// returns instead of a useful one.
var abpBoilerplate = regexp.MustCompile(`(?i)exception of type '[^']*' was thrown`)

type abpError struct {
	Error *struct {
		Code             string            `json:"code"`
		Message          string            `json:"message"`
		Details          string            `json:"details"`
		ValidationErrors []ValidationError `json:"validationErrors"`
	} `json:"error"`
	// Some endpoints answer with a flat object.
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func newAPIError(resp *Response) *APIError {
	e := &APIError{StatusCode: resp.StatusCode, CorrelationID: resp.CorrelationID, Body: resp.Body}
	body := strings.TrimSpace(string(resp.Body))

	switch resp.StatusCode {
	case http.StatusFound, http.StatusMovedPermanently:
		// The multi-reseller middleware redirects to /Error/Suspend when the
		// __reseller value is unknown, inactive or suspended.
		e.Code = "INVALID_RESELLER"
		e.Message = "authentication failed: the reseller ID is unknown, inactive or suspended"
		return e
	case http.StatusTooManyRequests:
		e.Code = "RATE_LIMITED"
		e.Message = "rate limit exceeded"
		if s := resp.Header.Get("Retry-After"); s != "" {
			if n, err := strconv.ParseFloat(s, 64); err == nil {
				e.RetryAfter = time.Duration(n * float64(time.Second))
				e.Message += fmt.Sprintf(" (retry after %ss)", s)
			}
		}
		return e
	}

	var ab abpError
	if body != "" && (strings.HasPrefix(body, "{") || strings.HasPrefix(body, "[")) && json.Unmarshal(resp.Body, &ab) == nil {
		if ab.Error != nil {
			e.Code, e.Message, e.Details, e.ValidationErrors = ab.Error.Code, ab.Error.Message, ab.Error.Details, ab.Error.ValidationErrors
		} else {
			e.Code, e.Message, e.Details = ab.Code, ab.Message, ab.Details
		}
	} else if body != "" {
		e.Message = body
	}

	if resp.StatusCode == http.StatusUnauthorized && e.Code == "" {
		lower := strings.ToLower(e.Message)
		switch {
		case strings.Contains(lower, "api key is missing"):
			e.Code = "MISSING_API_KEY"
		case strings.Contains(lower, "resellerid is missing"):
			e.Code = "MISSING_RESELLER_ID"
		case strings.Contains(lower, "invalid api key"):
			e.Code = "INVALID_API_KEY"
		case strings.Contains(lower, "authorisation"), strings.Contains(lower, "authorization"):
			e.Code = "SUBRESELLER_NOT_ALLOWED"
		default:
			e.Code = "UNAUTHORIZED"
		}
	}

	// Replace boilerplate or code-echo messages with the details field.
	if abpBoilerplate.MatchString(e.Message) || (e.Message != "" && e.Message == e.Code) {
		d := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(e.Details), "Dna "))
		if d != "" && !abpBoilerplate.MatchString(d) && d != e.Code {
			e.Message = d
		} else {
			e.Message = "the operation could not be completed (error " + e.Code + ")"
		}
	}
	if e.Code == "" {
		e.Code = "HTTP_" + strconv.Itoa(resp.StatusCode)
	}
	if strings.TrimSpace(e.Message) == "" {
		e.Message = strings.ToLower(http.StatusText(resp.StatusCode))
		if e.Message == "" {
			e.Message = "unexpected status"
		}
	}
	return e
}
