package main

import "fmt"

func main() {
	//hitung rata-rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var total int
	for _, score := range scores { //
		total = total + score
	}
	length := len(scores)
	average := float64(total) / float64(length)

	fmt.Println(average)

	//hitung nilai terbesar
	scores2 := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int

	for _, score2 := range scores2 {
		if score2 >= 90 {
			goodScores = append(goodScores, score2)
		}
	}
	fmt.Println(goodScores)
}
