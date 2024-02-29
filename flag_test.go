package belajar_golang_standard_library

import (
	"flag"
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "12345678", "database password")
	host := flag.String("host", "localhost", "database host")
	port := flag.Int("port", 3306, "database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Hostname", *host)
	fmt.Println("Port", *port)
}
