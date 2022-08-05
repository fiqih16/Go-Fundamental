package function

import (
	"errors"
	"fmt"
)

func Printresultkuis() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total) 

	result, err := calculates(10, 2, "=")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func calculates(num1, num2 int, operation string) (int, error) {
	var result int
	var errorResult error
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
			result = num1 / num2
	default:
		errorResult = errors.New("Invalid operation")
	}
	return result, errorResult

}