package management

import (
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user)
}

type Group struct {
	Name     string
	admin    User
	users    []User
	IsActive bool
}

func (group Group) displayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("nenbyat ")
}
