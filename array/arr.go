package array

import (
	"fmt"
)

func Arrayone() {

	// cara pertama
	// var languages [3]string
	// languages[0] = "Go"
	// languages[1] = "Python"
	// languages[2] = "Java"

	// cara kedua
	// languages := [3]string{"Go", "Python", "Java"}

	// cara ketiga
	languages := [...]string{"Go", "Python", "Java"}

	fmt.Println(languages)
	length := len(languages)
	fmt.Println("Isi data array :",length)

}

func Arraytwo(){
	languages := [...]string{"PHP", "JS", "Kotlin", "Dart"}

	for index, lang := range languages {
		fmt.Println("index :", index, "lang :", lang)
	}
}