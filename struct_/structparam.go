package struct_

import (
	"fmt"
)

type UserParam struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func PrintUserParam(){
	//user := UserParam{ID: 1, FirstName: "John", LastName: "Doe", Email: "jon@gmail.com", IsActive: true}
	user := UserParam{1, "John","Doe","jon@gmail.com",true}
	
	displayUser1 := displayUserParam(user)
	fmt.Println(displayUser1)
}

func displayUserParam(userpar UserParam) string{
	result := fmt.Sprintf("Nama: %s %s, Email: %s", userpar.FirstName, userpar.LastName, userpar.Email)
	return result
}