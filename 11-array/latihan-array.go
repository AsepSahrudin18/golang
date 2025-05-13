package main
import (
	"fmt"
)

func main() {
	var names [3] string
	names[0] = "Asep"
	fmt.Println(names[0]) // output: Asep

	names[0] = "Asah"
	fmt.Println(names[0]) // outputnya: Asah


	employes := [...]string { "Karim", "Saepulloh", "Asep Sahrudin", "Iqbal", "Riska", "Hangga"}

	fmt.Println(employes)



	fruits := [...]string {"nanas", "apel"}
	fruits[0]=""
	fmt.Println(fruits[0])
}