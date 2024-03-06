package belajar_golang_standard_library

import (
	"fmt"
	"reflect"
	"testing"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Printf("structField.Tag.Get(\"required\"): %v\n", structField.Tag.Get("required"))
		fmt.Printf("structField.Tag.Get(\"max\"): %v\n", structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func TestReflect(t *testing.T) {
	readField(Sample{"Robi"})
	readField(Person{"Robi", "Depok", "robi@email.com"})

	person := Person{
		Name:    "ada",
		Address: "ada",
		Email:   "ada",
	}

	fmt.Println(IsValid(person))
}
