package utils

import "fmt"

type AppError struct {
	ErrorCode    int
	ErrorMessage string
}

func (appErr *AppError) Error() string {
	return fmt.Sprintf("%v - %v", appErr.ErrorCode, appErr.ErrorMessage)
}
