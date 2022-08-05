package function

import "fmt"

func Printresult() {
	sentence := myResult("Hello World, ")
	fmt.Println(sentence)

	result := add(10, 20)
	fmt.Println(result)

	luas, keliling := calculate(10, 2)
	fmt.Println(luas)
	fmt.Println(keliling)
}

func myResult(sentence string) string {
	newSentence := sentence + "saya sedang belajar golang"
	return newSentence
}

func add(num1 int, num2 int) int {
	return num1 + num2
	
}

func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}