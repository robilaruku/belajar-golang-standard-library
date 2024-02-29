package belajar_golang_standard_library

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrConv(t *testing.T) {
	result, err := strconv.ParseBool("false")

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)

	fmt.Println(binary)

	stringInt := strconv.Itoa(999)

	fmt.Println(stringInt)
}
