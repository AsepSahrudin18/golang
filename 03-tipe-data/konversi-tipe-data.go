package main 
import ("fmt")

func main() {
	var nilai8 int8 = 127;
	fmt.Println(nilai8); // untuk int8 maksimum 127 jika lebih maka akan error: cannot use .. (untyped int constant) as int8 value in variable declaration (overflows)
}