package errs

import "errors"

type MyError error

var (
	UNKNOWN_ROUTER_TYPE_ERROR  MyError = errors.New("Unknown router type")
	UNKNOWN_SERVICE_TYPE_ERROR MyError = errors.New("Unknown service type")

	STORE_CONNECTION_IS_DEFINED_ERROR         = errors.New("Store connection is defined")
	STORE_CONNECTION_UNDEFINED_ERROR          = errors.New("Store connection undefined")
	STORE_MIGRATION_ERROR                     = errors.New("Migration error")
	STORE_MODEL_EXSITS_ERROR          MyError = errors.New("Model alredy exsist")
	STORE_UNREADY_ERROR               MyError = errors.New("Store unready")
)
