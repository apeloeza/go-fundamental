package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	admin       User
	Users       []User
	IsAvailable bool
}

func main() {
	user := User{1, "one", "one", "one.one@ocbc.com", true}
	user2 := User{2, "two", "two", "two.two@ocbc.com", true}
	Users := []User{user, user2}
	group := Group{"DrawingTim", user, Users, true}
	// displayGroup(group)
	group.displayGroup()
}

func (group Group) displayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("users name :")
	for _, user := range group.Users {
		fmt.Print(user.FirstName)
	}

}
