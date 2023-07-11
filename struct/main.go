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

	//? struct sebagai parameter dari function =============================================

	user5 := User{5, "Zelda", "Lightning", "zelda.lightning@gmail.com", true}
	user6 := User{6, "Equard", "Dodi", "equard.dodi@gmail.com", true}
	user7 := User{7, "kopi", "nako", "kopinako@nako.com", true}

	displayUser1 := displayUser(user5)
	displayUser2 := displayUser(user6)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)

	fmt.Println(user7)

	//?==========================================

	user8 := User{8, "sah", "rini", "sahrini@rini.com", true}
	user9 := User{9, "ri", "ri", "riri@ri.com", true}
	users := []User{user8, user9}
	group := Group{"Gamer", user, users, true}
	displayGroup(group)

}

//? embeded struct juga bisa menjadi atribut/ field sebuah struck lain

type Group struct {
	name        string
	admin       User
	users       []User
	IsAvailable bool
}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.users))
	fmt.Println("")

	for _, user := range group.users {
		fmt.Println(user.FirstName)
	}
}

func displayUser(user8 User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user8.FirstName, user8.LastName, user8.Email)
	//%s = menampilkan string
	return result
}

//ss

