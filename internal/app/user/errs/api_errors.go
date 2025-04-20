package errs

import "errors"

var (
	ErrBadUUID  = errors.New("Invalid UUID format")
	ErrBadJSON  = errors.New("Invalid input body")
	ErrNotFound = errors.New("User not found")
)
