package belajar_golang_standard_library

import (
	"errors"
	"fmt"
	"testing"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Robi" {
		return NotFoundError
	}

	// sukses

	return nil
}

func TestErrorPkg(t *testing.T) {
	err := GetById("Robi")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
			
		}
	}
}
