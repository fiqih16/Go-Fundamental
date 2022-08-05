package slice

import "fmt"

func Sliceone() {
	// cara pertama
	var gamingConsoles []string
	gamingConsoles = append(gamingConsoles, "Playstation4")
	gamingConsoles = append(gamingConsoles, "XboxOne")
	gamingConsoles = append(gamingConsoles, "NintendoSwitch")

	// cara kedua
	//gamingConsoles := []string{"Playstation4", "XboxOne", "NintendoSwitch"}

	for _, console:= range gamingConsoles {
		fmt.Println(console)
	}

	//fmt.Println(gamingConsoles)
	lenght := len(gamingConsoles)
	fmt.Println("Isi data array :", lenght)
}