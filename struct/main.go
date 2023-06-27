package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Arif"
	user.LastName = "Semangat"
	user.Email = "arif.semangat@gmail.com"
	user.IsActive = true

	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Arif"
	user2.LastName = "Semangat"
	user2.Email = "arif.semangat@gmail.com"
	user2.IsActive = true

	user3 := User{ID: 3, FirstName: "Ilham", LastName: "Bergembira", Email: "ilham.gembira@gmail.com", IsActive: true}

	user4 := User{4, "Ilham", "Hebat", "ilham.hebat@gmail.com", true}

	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)
	fmt.Println(user4)

	//? struct sebagai parameter ===================================

	user5 := User{5, "Zelda", "Lightning", "zelda.lightning@gmail.com", true}
	user6 := User{6, "Equard", "Dodi", "equard.dodi@gmail.com", true}

	displayUser1 := displayUser(user5)
	displayUser2 := displayUser(user6)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)
}

func displayUser(user5 User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user5.FirstName, user5.LastName, user5.Email)
	//%s = menampilkan string
	return result
}
