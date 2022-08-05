package perulanganfor

import (
	"fmt"
)

func Forone() {
	for i := 1; i < 5; i++ {
		fmt.Println("Cara pertama menggunakan for :", i)
	}
}

func Fortwo(){
	number := 1
	for number <= 5 {
		fmt.Println("Cara kedua menggunakan for :", number)
		number++
	}
}

func Forthree(){
	title := "Cara ketiga"
	for _, letter := range title {
		fmt.Println("letter :", string(letter))
	}
}


func Findeven(){
	data := "Golang the best language"
	// Menampilkan angka genap
	for index, letter := range data {
		if index % 2 == 0 {
			fmt.Println("Menampilkan angka genap, index :", index, "letter :", string(letter))
		}
	}
	// Menampilkan huruf vokal A, E, I, O, U
	for index, letter := range data {
		letterString := string(letter)
		// menggunakan if
		// if letterString == "a" || letterString == "e" || letterString == "i" || letterString == "o" || letterString == "u" {
		// 	fmt.Println("Menampilkan huruf vokal, index :", index, "letter :", string(letter))
		// }
		// menggunakan switch
		switch letterString {
		case "a", "e", "i", "o", "u":
			fmt.Println("Menampilkan huruf vokal, index :", index, "letter :", string(letter))
		} 
	}
}