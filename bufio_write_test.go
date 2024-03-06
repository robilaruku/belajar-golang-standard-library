package belajar_golang_standard_library

import (
	"bufio"
	"os"
	"testing"
)

func TestBufioWrite(t *testing.T) {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("Selamat Belajar\n")
	_ = writer.Flush()
}
