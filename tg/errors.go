package tg

type BuildRequestError struct {
	message string
}

func (e *BuildRequestError) Error() string {
	return e.message
}

func NewBuildRequestError(message string) *BuildRequestError {
	return &BuildRequestError{
		message: message,
	}
}

type SendRequestError struct {
	message string
}

func (e *SendRequestError) Error() string {
	return e.message
}

func NewSendRequestError(message string) *SendRequestError {
	return &SendRequestError{
		message: message,
	}
}

type ParseResponseBodyError struct {
	message string
}

func (e *ParseResponseBodyError) Error() string {
	return e.message
}

func NewParseResponseBodyError(message string) *ParseResponseBodyError {
	return &ParseResponseBodyError{
		message: message,
	}
}

type APIError struct {
	message string
}

func (e *APIError) Error() string {
	return e.message
}

func NewAPIError(message string) *APIError {
	return &APIError{
		message: message,
	}
}
