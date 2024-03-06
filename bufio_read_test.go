package belajar_golang_standard_library

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestBufioRead(t *testing.T) {
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
