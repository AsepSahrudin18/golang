package main

import ("fmt")


// function tanpa return tanpa parameter
func sapa() {
	fmt.Println("Hallo dari fungsi!");
}

// function dengan parameter
func hello(name string) {
	fmt.Println("hai apa kabar? ",name)
}

// function dengan return parameter
func question(name string) string {
	return "Bagaimana kabar kamu hari ini? " + name
}

// function dengan beberapa parameter
// note: pendefinisan tipe data sekali seperti ini jika casenya tipe data paramter sama
func tambah(a, b int) int {
	return a + b;
}

// jika paramter lebih dari satu dan tipe datanya berbeda
func info(name string, umur int) (string, int) {
	return name, umur;
}

// function dengan multiple return value
func hitung(a int, b int) (int, int) {
	return a + b, a * b;
}

// function dengan return bernama, kita juga bisa melakukan return dengan memberi nama nilai yang dikembalikan.
func bagi(a, b int) (hasil int) {
	hasil = a / b;
	return // cukup gunakan return tanpa argument
}

func main() {
	// ini untuk memanggil fungsi yang lain yaa. untuk fungsi yang dipelajari yang diatas"nya. jadi fahami yang atasnya. (child nya) ini parentnya.
	sapa();
	hello("Asep Sahrudin");

	message :=question("Asep Sahrudin");
	fmt.Println(message);

	tambahAngka := tambah(10, 20);
	fmt.Println(tambahAngka);

	getName, getUmur := info("Asep Sahrudin", 25)
	fmt.Println(getName, getUmur);

	// gunakan underscore _ jika salah satu variable tidak digunakan
	getNameAgain,_ := info("Asep Sahrudin", 25)
	fmt.Println(getNameAgain);

	getTambah, getKPerkalian := hitung(25, 25);
	fmt.Println(getTambah);
	fmt.Println(getKPerkalian);

	getBagi := bagi(10, 3);
	fmt.Println(getBagi);
}