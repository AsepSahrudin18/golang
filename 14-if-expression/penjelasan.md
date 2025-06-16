## if expression

Dalam bahasa Go (Golang), if expression adalah pernyataan logika yang digunakan untuk mengevaluasi kondisi. Jika kondisi bernilai true, maka blok kode di dalam if akan dieksekusi.

- blok if akan dieksekusi ketika kondisi if bernilai true
- kadang kita ingin melakukan program tertentu jika kondisi if bernilai false
- hal ini bisa dilakukan menggunakan else expression

```
func main() {
	nama := "Asep Sahrudin"

	if nama == "Asep Sahrudin" { // jika nama benar maka ini dieksekusi
		fmt.Println("Hallo " + nama)
	} else {
		fmt.Println("Nama Anda Salah") // jika nama salah maka ini di eksekusi
	}
}
```

### if dengan short statement

- if mendukung short statement sebelum kondisi
- hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi

```
// if dengan short statement
	username := "sahrudindaritanahgarut";
	if length := len(username); length > 5 {
		fmt.Println("username terlalu panjang")
	} else {
		fmt.Println("Nama Sudah benar")
	}
```
