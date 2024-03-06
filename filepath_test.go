package belajar_golang_standard_library

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFilepath(t *testing.T) {
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}