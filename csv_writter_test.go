package belajar_golang_standard_library

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestCsvWritter(t *testing.T) {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"robi", "laruku", "dream"})
	_ = writer.Write([]string{"budi", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush()
}
