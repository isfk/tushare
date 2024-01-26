package tushare

import "fmt"

type TushareError struct {
	Code    int
	Message string
}

func NewTushareError(code int, message string) *TushareError {
	return &TushareError{Code: code, Message: message}
}

func (tse *TushareError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", tse.Code, tse.Message)
}
