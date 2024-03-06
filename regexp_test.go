package belajar_golang_standard_library

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	regex := regexp.MustCompile(`r([a-z])([a-z])i`)

	fmt.Println(regex.MatchString("robi"))
	fmt.Println(regex.MatchString("royi"))
	fmt.Println(regex.MatchString("roKi"))

	fmt.Println(regex.FindAllString("robi roki rodi reyi redi reDi", 10))
}