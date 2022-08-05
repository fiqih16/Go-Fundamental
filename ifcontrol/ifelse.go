package ifcontrol

import (
	"fmt"
)

func Age() {
	age := 20
	if age > 20 {
		fmt.Println("Lebih dari 20 tahun")
	} else {
		fmt.Println("Kurang dari 20 tahun")
	}
}

func Score(){
	score := 75
	var grade string

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "E"
	}
	fmt.Println(grade)
}