package belajar_golang_standard_library

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestCsvReader(t *testing.T) {
	csvString := "robi,larukku,dream\n" +
		"budi,pratama,nugraha\n" +
		"joko,morro,diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record )
	}
}
