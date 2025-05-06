package main
import ("fmt")

func main() {
	nilai := 85;

	if nilai >= 90 {
		fmt.Println("nilai A")
	} else if nilai >= 80 {
		fmt.Println("nilai B")
	} else {
		fmt.Println("nilai C atau dibawahnya")
	}


	hari := "Senin";
	switch hari {
	case "Senin":
		fmt.Println("Awal hari!")
	case "Sabtu", "Minggu":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Hari biasa")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}