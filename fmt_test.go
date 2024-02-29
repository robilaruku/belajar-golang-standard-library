package belajar_golang_standard_library

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {
	fmt.Println("Hello, World!!")

	firstName := "Robi"
	lastName := "Laruku"

	fmt.Printf("Hello, '%s %s', '%s'' !!\n", firstName, lastName, "wkwkwkwk")
}
