package main

import "fmt"

func main() {
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
}
