package main

import ("fmt")

func main() {

	// switch expression biasa
	// nama := "sahrudin"
	// switch nama {
	// case "sahrudin":
	// 	fmt.Println("hai " + nama + " kamu sekarang adalah manager")
	// case "andi":
	// 	fmt.Println("hai" + nama + "kamu adalah staff ahli")
	// default:
	// 	fmt.Println("namamu belum terdaftar")
	// }

	// switch statement dengan short statement
	username := "sahrudin18asep";
	switch user := len(username); user > 5 {
	case true: 
	fmt.Println("hallo sahrudin");
	case false:
		fmt.Println("hallo user");
	default:
		fmt.Println("hallo boleh kenalan?")
	}
}