package main
import (
	"fmt"
)

func main() {
	days := [...] string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daysSlice1 := days[5:];

	// ubah slice value 
	daysSlice1[0] = "Sabtu Baru";
	fmt.Println(days) // output: [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu] // ketika mengubah slice maka yang akan berubah adalah isi dari arraynya.
}