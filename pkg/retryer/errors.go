package retryer

import "errors"

var (
	ErrMaxStepsReached = errors.New("maximum retry steps reached")
)
