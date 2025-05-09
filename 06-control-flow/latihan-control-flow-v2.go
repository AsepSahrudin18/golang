package main
import ("fmt")



// sayHello
func sayHello(name string) string {
	return "hallo selamat sore, " + name
}

func main() {
	// logic if
	nilai := 85;
	if	nilai >= 90 {
		fmt.Println("nilai A")
	} else if nilai >= 80 {
		fmt.Println("nilai B")
	} else {
		fmt.Println("nilai C")
	}

	// switch case
	hari := "senin";
	switch hari {
	case "sabtu","minggu":
		fmt.Println("weekend");
	case "senin":
		fmt.Println("awal hari")
	default:
		fmt.Println("hari biasa")
	}

	// for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	message := sayHello("Asep Sahrudin")
	fmt.Println(message)
}


