package main 
import ("fmt")

func main() {
	fmt.Println("ini adalah string")


	lenString := "asep sahrudin";
	panjangKarakter := len(lenString)
	fmt.Println(panjangKarakter)

	getChar := "Asep Sahrudin";
	fmt.Println("Total Char: ", string(getChar[0]))

	fmt.Println(string("ngetes string"[1]))
}