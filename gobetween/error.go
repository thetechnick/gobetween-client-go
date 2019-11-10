package gobetween

type apiError struct {
	err  string
	body []byte
}

func (e apiError) Body() []byte {
	return e.body
}

func (e apiError) Error() string {
	return e.err + ": " + string(e.body)
}

type NotFoundError struct{ apiError }

func newNotFoundError(body []byte) NotFoundError {
	return NotFoundError{
		apiError{
			err:  "not found",
			body: body,
		},
	}
}

func IsNotFoundError(err error) bool {
	_, ok := err.(NotFoundError)
	return ok
}

type BadRequestError struct{ apiError }

func newBadRequestError(body []byte) BadRequestError {
	return BadRequestError{
		apiError{
			err:  "bad request",
			body: body,
		},
	}
}

func IsBadRequestError(err error) bool {
	_, ok := err.(BadRequestError)
	return ok
}

type ConflictError struct{ apiError }

func newConflictError(body []byte) BadRequestError {
	return BadRequestError{
		apiError{
			err:  "conflict",
			body: body,
		},
	}
}

func IsConflictError(err error) bool {
	_, ok := err.(ConflictError)
	return ok
}

type InternalError struct{ apiError }

func newInternalError(body []byte) BadRequestError {
	return BadRequestError{
		apiError{
			err:  "internal",
			body: body,
		},
	}
}

func IsInternalError(err error) bool {
	_, ok := err.(InternalError)
	return ok
}

type UnauthorizedError struct{ apiError }

func newUnauthorizedError(body []byte) UnauthorizedError {
	return UnauthorizedError{
		apiError{
			err:  "unauthorized",
			body: body,
		},
	}
}

func IsUnauthorizedError(err error) bool {
	_, ok := err.(UnauthorizedError)
	return ok
}
