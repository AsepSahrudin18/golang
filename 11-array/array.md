# array

✅ Apa itu Array di Golang?
Array adalah koleksi elemen dengan:

- Tipe data yang sama
- Ukuran jumlah array tetap (tidak bisa diubah setelah dibuat)
- Data bisa diubah namun tidak bisa bertambah.

⚠️ Catatan Penting:
Ukuran array adalah bagian dari tipe data-nya, contoh: [3]int ≠ [4]int

Gunakan slice jika kamu butuh array yang ukurannya bisa berubah.

note: untuk urutan array sama seperti di pemrograman lain yaitu mulai dari index ke 0.

rule penulisan array di go:

- array di go harus ditentukan dulu panjang arraynya (karena nilainya tidak bisa bertambah).
- harus tentukan terlebih dahulu tipe datanya dan hanya bisa menampung satu jenis tipe data.

contoh:

```
var names [3] string // angka 3 pada tanda kurung kotak adalah panjang array dan string adalah tipe datanya.
names[0] = "Asep"
names[1] = "Sahrudin"
names[2] = "Arul"
names[4] = "error ini" // jika lebih dari 3 (lebih dari jumlah array yang ditentukan) maka akan error.

fmt.Println(names[1]) // output: Sahrudin
```

contoh kalau array dapat diubah:

```
var names [3] string
names[0] = "Asep"
fmt.Println(names[0]) // output: Asep

names[0] = "Asah"
fmt.Println(names[0]) // outputnya: Asah
``
```

## Membuat array secara langsung

di Go-Lang, kita juga bisa membuat array secara langsung

contoh:

```
var employes = [6] string { "Karim", "Saepulloh", "Asep Sahrudin", "Iqbal", "Riska", "Hangga"}

atau dengan :=

employes := [6] string { "Karim", "Saepulloh", "Asep Sahrudin", "Iqbal", "Riska", "Hangga"}


atau jika ingin panjangnya otomatis. (gunakan tanda titik 3 pada panjang array)

employes := [...] string { "Karim", "Saepulloh", "Asep Sahrudin", "Iqbal", "Riska", "Hangga"}
```

note: pada array tidak ada operasi untuk menghapus data array, yang bisa dilakukan adalah mengganti nilai array dengan string kosong untuk mengkosongkan valuenya.

```
fruits := [...]string {"nanas", "apel"}
fruits[0]=""
fmt.Println(fruits[0])

note:
fruits := [...]string {"nanas", "apel"}
fmt.Println(fruits[0]="") // di golang tidak bisa melakukan langsung ganti value pada print seperti ini. fruits[0] = "" adalah assignment, bukan expression → tidak bisa langsung dijadikan argumen di Println
```
