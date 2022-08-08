package pointer

import "fmt"

func PrintPointer() {
	var numberA int = 10
	var numberB *int = &numberA
	fmt.Println("Number A : ", numberA)
	fmt.Println("Number B : ", numberB)
	fmt.Println("Number B Value : ", *numberB)

	numberA = 20
	fmt.Println("Number A : ", numberA)
	fmt.Println("Number B : ", numberB)
}