package belajar_golang_standard_library

import (
	"fmt"
	"strings"
	"testing"
)

func TestStr(t *testing.T) {
	fmt.Println(strings.Contains("Robi Laruku", "Robi"))
	fmt.Println(strings.Split("Robi Laruku", " "))
	fmt.Println(strings.ToLower("Robi Laruku"))
	fmt.Println(strings.ToUpper("Robi Laruku"))
	fmt.Println(strings.Trim("    Robi Laruku    ", " "))
	fmt.Println(strings.ReplaceAll("Robi Laruku Dream", "Robi", "Budi"))
}
