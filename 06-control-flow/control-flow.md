# Control Flow (if, switch, loop)

control flow adalah struktur yang mengatur urutan eksekusi kode dalam program. Dengan control flow, kita bisa membuat program mengambil keputusan (menggunakan kondisi) dan melakukan perulangan (looping). Di Golang, struktur control flow yang umum digunakan adalah: if, switch, dan for.

## 1. Struktur Kondisional (if, else if, else)

Struktur if digunakan untuk menjalankan blok hanya jika suatu kondisi bernilai true.

```
nilai := 85

if nilai >= 90 {
    fmt.Println("Nilai A")
} else if nilai >= 80 {
    fmt.Println("Nilai B")
} else {
    fmt.Println("Nilai C atau di bawahnya")
}
```

## 2. Struktur pemilihan: Switch

switch digunakan untuk mengecek satu nilai terhadap beberapa kemungkinan secara lebih rapi dibandingkan if-else

```
hari := "Senin";
	switch hari {
	case "Senin":
		fmt.Println("Awal hari!")
	case "Sabtu", "Minggu":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Hari biasa")
	}
```

catatan penting:

- Tidak perlu break, karena golang otomatis keluar setelah satu case terpenuhi.
- Bisa menuliskan beberapa nilai dalam satu case (seperti: "Sabtu", "Minggu")

## 3. Perulangan: for

Golang hanya memiliki satu kata kunci untuk perulangan, yaitu for. Tapi flexible banget, bisa dipakai sebagai while, do-while, atau for-each

1. standar (mirip di C/C++)

```
for i := 0; i < 5; i++ {
		fmt.Println(i)
}
```

2. tanpa kondisi (infinite loop)

```
for {
    fmt.Println("terus jalan")
    break
}
```

3. Dengan range
   Digunakan untuk melakukan iterasi pada array, slice, map, atau string.

```
angka := []int{1, 2, 3, 4}

for index, value := range angka {
    fmt.Println(index, value)
}
```

kalau kita butuh nilainya saja:

```
for _, value := range angka {
    fmt.Println(value)
}
```

### kesimpulan.

Control flow di golang membantu kita mengatur logika program dengan baik. Dengan if, switch , dan for, kita bisa membuat program yang cerdas dalam mengambil keputusan dan melakukan perulangan.

memahami struktur-struktur ini sangat penting sebelum masuk ke konsep yang lebih kompleks seperti function, struct, ata goroutine.
