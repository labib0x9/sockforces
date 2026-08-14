package submissions

import "errors"

var (
	ErrCreateRepoFailed = errors.New("Could not clone: Name already exists on this account")
)
