package app_errors

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

var EntityNotFound = gorm.ErrRecordNotFound

type CommonError struct {
	ErrorCode    int    `json:"code"`
	HttpCode     int    `json:"-"`
	ErrorMessage string `json:"message"`
	ErrorTrace   error  `json:"-"`
}

func (err *CommonError) Error() string {
	return fmt.Sprintf("CommonError: %+v. Trace: %+v", err.ErrorMessage, err.ErrorTrace.Error())
}

func NewError(err error, message string) *CommonError {
	cerr := &CommonError{
		ErrorMessage: message,
		ErrorTrace:   err,
	}

	if errors.As(err, &EntityNotFound) {
		cerr.HttpCode = http.StatusNotFound
	}

	return cerr
}
