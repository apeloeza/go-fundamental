package main

import (
	"fmt"
	"fundamental/calculation"
)

func main() {
	//? perkenalan di catatan
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
	//? if Control ==================================
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
	/* if
	if else
	else if
	if bersarang */
	//? switch Control ==============================
	age2 := 2
	switch age2 {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	}

	fmt.Println("perulangan1")

	//? perulangan (for) ==================================================
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
	//? Array =========================================================
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
	//? slice =============================================
	//mirip dengan array
	var gamingConsoles []string
	//cara 1
	gamingConsoles = append(gamingConsoles, "Playstation4")
	gamingConsoles = append(gamingConsoles, "Nintendo switch")
	gamingConsoles = append(gamingConsoles, "xbox ")

	fmt.Println(gamingConsoles)
	//cara 2
	gamingConsoles2 := []string{"playstation2", "Nintendo OLD", "Xbox 1"}
	fmt.Println(gamingConsoles2)

	cetak := calculation.TestAja()
	namaSaya := calculation.Namaku()
	//fmt.Println(TestAja) 	//salah
	fmt.Println(cetak, namaSaya)

	for _, console := range gamingConsoles2 {
		fmt.Println(console)
	}

	//? MAP =============================================

	//?cara 1
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["Rudy"] = 9
	// myMap["sedang"] = 10
	// myMap["belajar"] = 11
	// myMap["golang"] = 9

	// fmt.Println(myMap)
	//? cara 2
	// myMap := map[string]string{"andi": "sedang mandi", "bersih": "sangat baik"}

	// fmt.Println(myMap)
	//? cara 3 for menggunakan array
	myMap := map[string]string{
		"jogi":       "sangat jujur",
		"rani":       "ratu kelas",
		"javascript": "hype",
	}

	for key, value := range myMap {
		fmt.Println("key : ", key, "value: ", value)
	}
	//? cara 4 melakukan delete data MAP
	fmt.Println("=========")
	delete(myMap, "rani")

	for key, value := range myMap {
		fmt.Println("key : ", key, "value: ", value)
	}
	//? mencari apa ada value nya atau tidak
	value := myMap["Angular"]
	fmt.Println(value, "-")

	//? ,emcari value dengan friendly
	value, isAvailabe := myMap["Python"]
	fmt.Println(isAvailabe)
	fmt.Println(value)

	//?  SLICE MAP ========================
	students := []map[string]string{
		{"name": " Agung", "score": "A"},
		{"name": " Indra", "score": "B"},
		{"name": " Badang", "score": "C"},
	}
	//fmt.Println(students)
	//? mencetak nilai satu satu
	for _, student := range students {
		fmt.Println(student["name"], " - ", student["score"])
	}

	//? ===============================
}
