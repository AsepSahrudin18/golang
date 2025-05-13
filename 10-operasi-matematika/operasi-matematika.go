package main
import (
	"fmt"
)

func main() {
	var a = 10
	var b = 20

	var hasil = a + b;
	fmt.Println(hasil)

	// dengan shorthand (cara singkat) atau disebut penulisan Augmented Assignments
	var nilaiA  = 40;
	nilaiA += 10;
	fmt.Println(nilaiA);


	// contoh increment
	var increment = 0;
	increment++;
	fmt.Println(increment);

	// contoh decrement.
	var decrement = 0;
	decrement--;
	fmt.Println(decrement);
}