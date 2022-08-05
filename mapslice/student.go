package mapslice

import "fmt"

func Student() {
	students := []map[string]string{
		{"name": "John", "age": "20"},
		{"name": "Jane", "age": "21"},
		{"name": "Jack", "age": "22"},
	}

	for _, student := range students {
		fmt.Println("Nama :", student["name"], ", Umur :", student["age"])
	}
}