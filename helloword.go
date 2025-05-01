// untuk memulai menulis go:
// ketik: go mod init example.com/hello
package main
import ("fmt")

// func main() {
// 	fmt.Println("Hello Word");
// }

// dalam go jika kita menekan enter maka dianggap akhir dari sebuah statment dan akan menambahkan tanda titik koma (;) secara implisit (tidak muncul dalam kode sumber)
// contoh:
func main() 
{
	fmt.Println("Hello Word");
}

// jalankan kode tersebut maka apa yang akan terjadi?
// error: syntax error: unexpected semicolon or newline before