# Variable

apa itu variabel ?
variable adalah tempat untuk menyimpan sebuah nilai. Nilai ini bisa berupa angka, teks, boolean, atau tipe data lainnya.

nama variable sebaiknya deskriptif agar mudah difahami, dan mengikuti aturan yang berlaku (tidak diawali angka, tidak mengandung spasi, dan tidak mengandung simbol khusus)

dalam GO penulisan variable dapat di deklarasikan menjadi 3 cara:

- := -> digunakan untuk penulisan variable singkat, untuk tipe variable ini hanya bisa di deklarasikan di dalam fungsi (function scope), tidak bisa dideklarasikan ulang jika tidak ada variable baru di sisi kirinya. (saran gunakan ini terus terkait pendefinisian variable)
- const -> ini digunakan ketika nilai variable bersifat immutable (benar benar tidak akan berubah). contohnya untuk penulisan rumus seperti: const pi = 3.14 untuk menghitung lingkaran.
- var -> digunakan untuk mendefinisikan variable jika kebutuhan nilainya berubah dan kebutuhannya global scope, ini bersifat muttable (dapat berubah)

noted: di golang tidak boleh melakukan definisi variable yang sudah didefinisikan sebelumnya. kalau mau re-assign menggunakan tanda =

```
name := "Asep Sahrudin"
name = "nama diganti"
```

### Multiple Declaration

kita juga bisa mendeklarasikan beberapa variable sekaligus.

```
var a, b, c int =  1, 2, 3
```

atau tanpa menyebut tipe data secara explisit.

```
x, y, z := "A", "B", "C"
```

### Mengganti nilai variable

nilai variable bisa diubah selama program berjalan

```
var jumlah int = 10
jumlah = 20
```

jika menggunakan := pada variabel yang pernah dideklarasikan sebelumnya maka akan muncul error, maka gunakan = untuk mengubah nilai bukan :=

### variabel luar fungsi

kita juga bisa mendeklarasikan variabel diluar fungsi main, tapi hanya dengan var, bukan :=

```
var versi = "1.0.0"

func main() {
    fmt.Println("Versi:", versi)
}
```
