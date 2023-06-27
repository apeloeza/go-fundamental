package main

import "fmt"

func main() {
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
}
