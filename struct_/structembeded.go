package struct_

import "fmt"

type UserExample struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name string
	Admin UserExample
	Users []UserExample
	IsAvailable bool
}

func StructUserGroup() {
	user1 := UserExample{ 1, "Steven", "William", "email@gmail.com", true }
	user2 := UserExample{ 2, "Julio", "Stafanus", "email2@gmail.com", true }

	users := []UserExample{ user1, user2 }
	group := Group{ "Joni", user1, users, true }
	displayGroup(group)
}

func displayGroup(group Group){
	fmt.Printf("Name : %s\n", group.Name)
	fmt.Printf("Member Count : %d\n", len(group.Users))

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}