package main

import "fmt"

func main() {
	//hitung rata-rata
	//? solusi pertama
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var total int
	for _, score := range scores {
		total = total + score
	}

	length := len(scores)
	average := float64(total) / float64(length)
	fmt.Println(average)
	//? solusi kedua

	scores2 := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int
	for _, score := range scores2 {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println(goodScores)
}
