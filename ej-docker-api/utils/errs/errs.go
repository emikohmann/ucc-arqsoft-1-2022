package errs

import "net/http"

type APIError interface {
	StatusCode() int
	ErrorMessage() string
}

type apiErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (err apiErr) Error() string {
	return err.Message
}

func (err apiErr) StatusCode() int {
	return err.Status
}

func (err apiErr) ErrorMessage() string {
	return err.Message
}

func NewCustomError(status int, message string) apiErr {
	return apiErr{
		Status:  status,
		Message: message,
	}
}

func NewStatusBadRequestError(message string) apiErr {
	return NewCustomError(http.StatusBadRequest, message)
}

func NewStatusUnauthorizedError(message string) apiErr {
	return NewCustomError(http.StatusUnauthorized, message)
}

func NewStatusPaymentRequiredError(message string) apiErr {
	return NewCustomError(http.StatusPaymentRequired, message)
}

func NewStatusForbiddenError(message string) apiErr {
	return NewCustomError(http.StatusForbidden, message)
}

func NewStatusNotFoundError(message string) apiErr {
	return NewCustomError(http.StatusNotFound, message)
}

func NewStatusMethodNotAllowedError(message string) apiErr {
	return NewCustomError(http.StatusMethodNotAllowed, message)
}

func NewStatusNotAcceptableError(message string) apiErr {
	return NewCustomError(http.StatusNotAcceptable, message)
}

func NewStatusProxyAuthRequiredError(message string) apiErr {
	return NewCustomError(http.StatusProxyAuthRequired, message)
}

func NewStatusRequestTimeoutError(message string) apiErr {
	return NewCustomError(http.StatusRequestTimeout, message)
}

func NewStatusConflictError(message string) apiErr {
	return NewCustomError(http.StatusConflict, message)
}

func NewStatusGoneError(message string) apiErr {
	return NewCustomError(http.StatusGone, message)
}

func NewStatusLengthRequiredError(message string) apiErr {
	return NewCustomError(http.StatusLengthRequired, message)
}

func NewStatusPreconditionFailedError(message string) apiErr {
	return NewCustomError(http.StatusPreconditionFailed, message)
}

func NewStatusRequestEntityTooLargeError(message string) apiErr {
	return NewCustomError(http.StatusRequestEntityTooLarge, message)
}

func NewStatusRequestURITooLongError(message string) apiErr {
	return NewCustomError(http.StatusRequestURITooLong, message)
}

func NewStatusUnsupportedMediaTypeError(message string) apiErr {
	return NewCustomError(http.StatusUnsupportedMediaType, message)
}

func NewStatusRequestedRangeNotSatisfiableError(message string) apiErr {
	return NewCustomError(http.StatusRequestedRangeNotSatisfiable, message)
}

func NewStatusExpectationFailedError(message string) apiErr {
	return NewCustomError(http.StatusExpectationFailed, message)
}

func NewStatusTeapotError(message string) apiErr {
	return NewCustomError(http.StatusTeapot, message)
}

func NewStatusMisdirectedRequestError(message string) apiErr {
	return NewCustomError(http.StatusMisdirectedRequest, message)
}

func NewStatusUnprocessableEntityError(message string) apiErr {
	return NewCustomError(http.StatusUnprocessableEntity, message)
}

func NewStatusLockedError(message string) apiErr {
	return NewCustomError(http.StatusLocked, message)
}

func NewStatusFailedDependencyError(message string) apiErr {
	return NewCustomError(http.StatusFailedDependency, message)
}

func NewStatusTooEarlyError(message string) apiErr {
	return NewCustomError(http.StatusTooEarly, message)
}

func NewStatusUpgradeRequiredError(message string) apiErr {
	return NewCustomError(http.StatusUpgradeRequired, message)
}

func NewStatusPreconditionRequiredError(message string) apiErr {
	return NewCustomError(http.StatusPreconditionRequired, message)
}

func NewStatusTooManyRequestsError(message string) apiErr {
	return NewCustomError(http.StatusTooManyRequests, message)
}

func NewStatusRequestHeaderFieldsTooLargeError(message string) apiErr {
	return NewCustomError(http.StatusRequestHeaderFieldsTooLarge, message)
}

func NewStatusUnavailableForLegalReasonsError(message string) apiErr {
	return NewCustomError(http.StatusUnavailableForLegalReasons, message)
}

func NewStatusInternalServerErrorError(message string) apiErr {
	return NewCustomError(http.StatusInternalServerError, message)
}

func NewStatusNotImplementedError(message string) apiErr {
	return NewCustomError(http.StatusNotImplemented, message)
}

func NewStatusBadGatewayError(message string) apiErr {
	return NewCustomError(http.StatusBadGateway, message)
}

func NewStatusServiceUnavailableError(message string) apiErr {
	return NewCustomError(http.StatusServiceUnavailable, message)
}

func NewStatusGatewayTimeoutError(message string) apiErr {
	return NewCustomError(http.StatusGatewayTimeout, message)
}

func NewStatusHTTPVersionNotSupportedError(message string) apiErr {
	return NewCustomError(http.StatusHTTPVersionNotSupported, message)
}

func NewStatusVariantAlsoNegotiatesError(message string) apiErr {
	return NewCustomError(http.StatusVariantAlsoNegotiates, message)
}

func NewStatusInsufficientStorageError(message string) apiErr {
	return NewCustomError(http.StatusInsufficientStorage, message)
}

func NewStatusLoopDetectedError(message string) apiErr {
	return NewCustomError(http.StatusLoopDetected, message)
}

func NewStatusNotExtendedError(message string) apiErr {
	return NewCustomError(http.StatusNotExtended, message)
}

func NewStatusNetworkAuthenticationRequiredError(message string) apiErr {
	return NewCustomError(http.StatusNetworkAuthenticationRequired, message)
}
