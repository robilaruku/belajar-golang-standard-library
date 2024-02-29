package belajar_golang_standard_library

import (
	"fmt"
	"sort"
	"testing"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int)  {
	u[i], u[j] = u[j], u[i]
}


func TestSortInterface(t *testing.T) {
	users := []User{
		{"Robi", 20},
		{"Budi", 30},
		{"Joko", 25},
		{"Adit", 36}, 
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
