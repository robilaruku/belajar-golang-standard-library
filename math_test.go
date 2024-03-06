package belajar_golang_standard_library

import (
	"fmt"
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Max(10, 11))
	fmt.Println(math.Min(10, 11))
}
