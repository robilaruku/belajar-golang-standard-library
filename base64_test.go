package belajar_golang_standard_library

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	value := "Robi Irhamni"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
