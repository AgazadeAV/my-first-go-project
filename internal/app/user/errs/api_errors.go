package errs

import "errors"

var (
	ErrBadUUID  = errors.New("invalid UUID format")
	ErrBadJSON  = errors.New("invalid input body")
	ErrNotFound = errors.New("user not found")
)
