package switchcontrol

import (
	"fmt"
)

func Switch() {
	number := 5

	switch number {
	case 1:
		fmt.Println("Number 1")
	case 2:
		fmt.Println("Number 2")
	case 3:
		fmt.Println("Number 3")
	default:
		fmt.Println("Number tidak diketahui")
	}
}