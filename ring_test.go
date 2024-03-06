package belajar_golang_standard_library

import (
	"container/ring"
	"fmt"
	"strconv"
	"testing"
)

func TestRing(t *testing.T) {
	var data *ring.Ring = ring.New(5)

	// auto
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// manually

	// data.Value = "Value 1"

	// data = data.Next()
	// data.Value = "Value 2"

	// data = data.Next()
	// data.Value = "Value 3"

	// data = data.Next()
	// data.Value = "Value 4"

	// data = data.Next()
	// data.Value = "Value 5"

	data.Do(func(value any) {
		fmt.Println(value)
	})

}
