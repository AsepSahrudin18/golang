package main
import (
	"fmt"
)

func main() {
	dataArr := [...] int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	getSlice := dataArr[:]

	fmt.Println(getSlice) // output: [0 1 2 3 4 5 6 7 8 9]
}