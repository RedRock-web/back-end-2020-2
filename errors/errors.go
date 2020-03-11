package errors

import (
	"errors"
	"fmt"
)

func CheckError(err error, msg string) {
	if err != nil {
		errors.New(msg)
		fmt.Println(err)
	}
}
