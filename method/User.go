package method

import "fmt"

type UserMethod struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type GroupMethod struct {
	Name        string
	Admin       UserMethod
	Users       []UserMethod
	IsAvailable bool
}

func (user UserMethod) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func PrintMethod() {
	userMethod := UserMethod{1, "Method", "User", "Met@gmail.com", true}
	result := userMethod.display()
	fmt.Println(result)

	userMethod2 := UserMethod{2, "User", "Met", "Met2@gmail.com", true}
	fmt.Println(userMethod2.display())

	users := []UserMethod{userMethod, userMethod2}
	group := GroupMethod{"Jojon", userMethod, users, true}
	group.DisplayGroupMethod()
}

func (group GroupMethod) DisplayGroupMethod(){
	fmt.Printf("Name : %s\n", group.Name)
	fmt.Printf("Member Count : %d\n", len(group.Users))

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}