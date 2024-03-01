package belajar_golang_standard_library

import (
	"fmt"
	"slices"
	"testing"
)

func TestSlices(t *testing.T) {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Robi"))
	fmt.Println(slices.Index(names, "Robi"))
	fmt.Println(slices.Index(names, "Paul"))
}
