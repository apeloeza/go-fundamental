package main

import (
	"fmt"
	"fundamental/calculation"
)

func main() {
	fmt.Println("go Growing better")
	//sentence := TestAja()
	//println(sentence)  // ==> go run main.go entity.go
	result := calculation.Tambah(8, 12)
	fmt.Println(result)
	result2 := calculation.Perkalian(2, 2)
	fmt.Println(result2)
	var nama string = "Arif"
	var umur int = 25
	NomorVaforit := 14
	fmt.Println(nama, umur, NomorVaforit)
	age := 9
	var grade string
	if age < 10 {
		grade = "e"
	} else if age <= 8 {
		grade = "d"
	} else if age <= 5 {
		grade = "c"
	} else {
		grade = "null"
	}
	fmt.Println(grade)
	age2 := 2
	switch age2 {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	}
	for i := 1; i <= 3; i++ {
		fmt.Println("saya belajar dong", i)
	}
	i := 1
	for i <= 3 {
		fmt.Println("saya sedang belajar go", i)
		i++
	}
	title := "arif"
	for index, letter := range title {
		if index%2 == 0 { //modulo
			fmt.Println("index: ", index, " letter : ", string(letter))
		}
	}
	for index, letter := range title {
		letterString := string(letter)
		// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		// 	fmt.Println("index: ", index, " letter : ", string(letter))
		// }
		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index :", index, " letter :", string(letter))
		}
	}
	/*array = satu kumpulan variabel dengan tipedata yang sama
	cara 1
	var names [5]string
	names[0] = "golang"
	names[1] = "sangat"
	names[2] = "mudah"
	names[3] = "dan"
	names[4] = "ringan"
	*/
	//cara2
	/*names := [5]string{
		"golang",
		"sangat",
		"mudah",
		"dan",
		"ringan",
	}
	fmt.Println(names)
	length := len(names)
	fmt.Println(length)
	*/
	//cara 3
	names := [...]string{
		"golang",
		"sangat",
		"mudah",
		"dan",
		"ringan",
	}
	for index, lang := range names {
		fmt.Println("Index : ", index, " names : ", lang)
	}
	//slice
	//mirip dengan array
	var gamingConsoles []string
	//cara 1
	gamingConsoles = append(gamingConsoles, "Playstation4")
	gamingConsoles = append(gamingConsoles, "Nintendo switch")
	gamingConsoles = append(gamingConsoles, "xbox ")
	//cara 2
	fmt.Println(gamingConsoles)

}
