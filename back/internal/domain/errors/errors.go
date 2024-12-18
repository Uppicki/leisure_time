package errs

import "errors"

type MyError error

var (
	UNKNOWN_ROUTER_TYPE_ERROR MyError = errors.New("Unknown router type")
)
