package main

import "fmt"

func main() {
	printMyResult("saya ")
	printMyResult("ingin ")
	printMyResult("belajar ")

	sentence2 := printMyResult2("saya sedang")
	fmt.Println((sentence2))
	/*<OUTPUT> := //?<PARAMETER> ("saya sedang")  <== INI FUNCTION
	fmt.Println((sentence2)) */ //? <== memanggil isi dari output

	result := add(10, 20)
	result2 := less(20, 10)
	fmt.Println(result, result2)

	luas, keliling := calculate(10, 2)
	fmt.Println(luas, keliling)
	luas2, _ := calculate2(2, 3) //jika hanya butuh luas jangan dihapus tapi diganti dengan "-"
	fmt.Println(luas2)
	//fmt.Println(keliling2)

	luas3, keliling3 := calculate3(13, 22)
	fmt.Println(luas3)
	fmt.Println(keliling3)
}

func printMyResult(sentence string) {
	fmt.Println(sentence)
}

func printMyResult2(sentence2 string) string {
	newSentence := sentence2 + " belajar golang"
	return newSentence
}

//?function lebih lanjut
func add(number int, numberTwo int) int {
	result := number + numberTwo
	return result
}
func less(numberThree, numberFour int) int {
	return numberThree - numberFour
}

//? function dengan multiple return
func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}
func calculate2(panjang2, lebar2 int) (int, int) {
	luas2 := panjang2 * lebar2
	keliling2 := 2 * (panjang2 + lebar2)
	return luas2, keliling2
}

//?function dengan predevined return value
func calculate3(panjang3 int, lebar3 int) (luas3 int, keliling3 int) {
	luas3 = panjang3 * lebar3
	keliling3 = 2 * (panjang3 * lebar3)
	return
}

//parameter/input
//proses
//nilai/return value

//input
//proses
//output

//sebuah function bisa mengembalikan sebuah nilai
