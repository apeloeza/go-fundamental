package main

import (
	"fmt"
)

func main() {
	user := User{1, "one", "one", "one@one.com", true}
	result := user.Display()

	fmt.Println(result)

user2:
	-User{2, "two", "two"}
}
