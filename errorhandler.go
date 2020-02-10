package errorhandler

type ErrorHandler struct {
	Err error
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

