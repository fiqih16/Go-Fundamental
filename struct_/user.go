package struct_

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func StructUser() {
	// Cara pertama
	//user := User{ID: 1, FirstName: "John", LastName: "Doe", Email: "jon@gmail.com", IsActive: true}
	
	// Cara kedua
	user := User{}
	user.ID = 1
	user.FirstName = "Steven"
	user.LastName = "William"
	user.Email = "steven@gmail.com"
	user.IsActive = true

	fmt.Println(user)
}