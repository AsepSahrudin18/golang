package main 

import (
	"fmt"
)

var angka int = 10;
var ptr *int = &angka;

func main() {
	fmt.Println("nilai angka:", angka)
	fmt.Println("alamat angka:", &angka)
	fmt.Println("pointer ptr", ptr)
}