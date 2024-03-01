package belajar_golang_standard_library

import (
	"fmt"
	"path"
	"testing"
)

func TestPath(t *testing.T) {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))
}