package heimdallr

import "errors"

var (
	//ErrFormatError is returned if the ID number doesn't match the required regex format
	ErrFormatError = errors.New("ID Number incorrectly formatted")
)
