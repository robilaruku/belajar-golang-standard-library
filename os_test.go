package belajar_golang_standard_library

import (
	"fmt"
	"os"
	"testing"
)

func TestOsArgs(t *testing.T) {
	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
