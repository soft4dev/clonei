package errors

import "fmt"

type CustomError struct {
	Err       error
	ShowUsage bool
}

func (e *CustomError) Error() string {
	return e.Err.Error()
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func NewCustomError(msg string, args ...any) *CustomError {
	return &CustomError{
		Err:       fmt.Errorf(msg, args...),
		ShowUsage: true,
	}
}

func NewCustomErrorWithoutUsage(msg string, args ...any) *CustomError {
	return &CustomError{
		Err:       fmt.Errorf(msg, args...),
		ShowUsage: false,
	}
}

func ShouldShowUsage(err error) bool {
	if cmdErr, ok := err.(*CustomError); ok {
		return cmdErr.ShowUsage
	}
	return true
}
