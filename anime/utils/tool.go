package utils

import (
	"errors"
	"fmt"
	"log"
)

func ErrHappened(format string, a ...any) error {
	errMsg := fmt.Sprintf(format, a...)
	log.Println(errMsg)
	return errors.New(errMsg)
}
