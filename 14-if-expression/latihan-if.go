package main

import ("fmt")

func main() {

	// if biasa
	// nama := "Asep Sahrudin"
	// if nama == "Asep Sahrudin" { // jika nama benar maka ini dieksekusi
	// 	fmt.Println("Hallo " + nama)
	// } else {
	// 	fmt.Println("Nama Anda Salah") // jika nama salah maka ini di eksekusi
	// }

	// if dengan short statement
	username := "12345";
	if length := len(username); length > 5 {
		fmt.Println("username terlalu panjang")
	} else {
		fmt.Println("Nama Sudah benar")
	}

}