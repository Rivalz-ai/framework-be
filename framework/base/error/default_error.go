package error

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

var (
	ErrInternalServer = &Error{
		CodeField:     http.StatusInternalServerError,
		GRPCCodeField: codes.Internal,
		MessageField:  "An internal server error occurred. Please try again later",
		KeyField:      "ERROR_INTERNAL_SERVER",
	}

	ErrUnauthorized = &Error{
		CodeField:     http.StatusUnauthorized,
		GRPCCodeField: codes.Unauthenticated,
		MessageField:  "Access denied. Please authenticate with valid credentials",
		KeyField:      "ERROR_UNAUTHORIZED",
	}

	ErrBadRequest = &Error{
		CodeField:     http.StatusBadRequest,
		GRPCCodeField: codes.InvalidArgument,
		MessageField:  "The request was invalid or contained malformed parameters",
		KeyField:      "ERROR_BAD_REQUEST",
	}

	ErrNotFound = &Error{
		CodeField:     http.StatusNotFound,
		GRPCCodeField: codes.NotFound,
		MessageField:  "The requested page or resource could not be found",
		KeyField:      "ERROR_NOT_FOUND",
	}

	ErrForbidden = &Error{
		CodeField:     http.StatusForbidden,
		GRPCCodeField: codes.PermissionDenied,
		MessageField:  "Access to the requested page or resource is forbidden",
		KeyField:      "ERROR_FORBIDDEN",
	}

	ErrUnsupportedMediaType = &Error{
		CodeField:     http.StatusUnsupportedMediaType,
		GRPCCodeField: codes.InvalidArgument,
		MessageField:  "The media type of the requested resource is not supported",
		KeyField:      "ERROR_UNSUPPORTED_MEDIA_TYPE",
	}

	ErrConflict = &Error{
		CodeField:     http.StatusConflict,
		GRPCCodeField: codes.FailedPrecondition,
		MessageField:  "A conflict occurred with the current state of the resource",
		KeyField:      "ERROR_CONFLICT",
	}

	ErrTimeout = &Error{
		CodeField:     http.StatusRequestTimeout,
		GRPCCodeField: codes.DeadlineExceeded,
		MessageField:  "The request timed out",
		KeyField:      "ERROR_TIMEOUT",
	}
)
