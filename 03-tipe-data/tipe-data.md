# 1. tipe data number

- integer
- float

## integer

### 1. Int (Signed Integer)

Tipe data ini bisa menyimpan bilangan negatif maupun positif.

| tipe data | Nilai Minimun         | Nilai Maksimum      |
| --------- | --------------------- | ------------------- |
| int8      | -128                  | 127                 |
| int16     | -32768                | 32767               |
| int32     | -2147483648           | 2147483647          |
| int64     | - 9223372036854775808 | 9223372036854775807 |

### 2. Unsigned Int (Integer Tidak Negatif)

Tipe data ini hanya bisa menyimpan bilangan nol atau positif saja (tidak negatif).

| tipe data | Nilai Minimun | Nilai Maksimum       |
| --------- | ------------- | -------------------- |
| uint8     | 0             | 255                  |
| uint16    | 0             | 65535                |
| uint32    | 0             | 429496729            |
| uint64    | 0             | 18446744073709551615 |

noted: gunakan tipe data sesuai kebutuhan supaya tidak membebani memory.

Kegunaan Perbedaan Ini:

| Tipe            | Kapan Digunakan                                                                 |
| --------------- | ------------------------------------------------------------------------------- |
| `int`, `intN`   | Saat nilai bisa **negatif**, misalnya suhu, perubahan saldo, perbedaan nilai    |
| `uint`, `uintN` | Saat nilai **tidak mungkin negatif**, misalnya ID, jumlah barang, panjang array |

### contoh untuk membangun software akunting bisa gunakan int atau int64

# Catatan:

Jangan gunakan float (float32, float64) untuk nilai uang karena bisa menyebabkan pembulatan yang tidak akurat.

## float

✅ Kapan Float Digunakan?
float32 atau float64 cocok digunakan untuk nilai-nilai yang mengandung desimal dan tidak butuh presisi absolut, seperti:

| Kasus Penggunaan                   | Kenapa Cocok Pakai Float            |
| ---------------------------------- | ----------------------------------- |
| Perhitungan suhu (`36.6`°C)        | Desimal, dan tidak kritis presisi   |
| Data sensor (misal tekanan, jarak) | Nilai fluktuatif, toleransi presisi |
| Posisi koordinat GPS (`-6.1751`)   | Presisi cukup, bukan keuangan       |
| Statistik (misal rata-rata `3.75`) | Desimal bisa berubah-ubah           |

❌ Kapan Tidak Dianjurkan?
Untuk uang atau nilai keuangan, jangan pakai float, karena:

Float bisa menyimpan hasil seperti: 0.1 + 0.2 = 0.30000000000000004

Ini bisa bikin pembulatan atau perhitungan jadi kacau.

## Zero Value

jika sebuah variable di definisikan tanpa nilai awal, maka golang akan memberikan nilai default (zero value) tergantung pada tipenya.

| tipe data | Zero Value |
| --------- | ---------- |
| int       | 0          |
| float     | 0.0        |
| bool      | false      |
| string    | ""(kosong) |

noted: di golang tidak bisa melakukan konversi otomatis.
jadi ketika ingin menjumlahkan bilangan berbeda tipe data harus di konversi terlebih dahulu.

contohnya gini:
di bahasa javascript bisa gini:

```
let result = "5" + 2;  // hasilnya "52", 2 dikonversi otomatis ke string
```

di python.

```
a = 10
b = 3.5
c = a + b  # Python otomatis mengubah 10 jadi 10.0, hasil c = 13.5
```

❌ Di Go Tidak Ada Konversi Otomatis
Go sangat strict (ketat). Kamu tidak bisa menjumlahkan int dan float64 langsung misalnya.

Contoh error di Go:

```
var a int = 5
var b float64 = 6.7
c := a + b  // ❌ error: mismatched types int and float64
```

# 2. Tipe data String

adalah tipe data karakter yang biasa kita sebut teks.

function dalam string

- len("string") -> digunakan untuk menghitung jumlah karakter
- "string"[0] -> mengambil posisi dalam string (perlu diingat bahwa di golang string yang diambil ini bentuknya adalah byte, jadi ketika ingin mengambil karakter tertentu harus di konversi dulu)

contoh len:

```
fmt.Println(len("asep sahrudin")) // outputnya: 13
```

contoh ambil posisi string

```
fmt.Println("asep sahrudin"[0]) // outputnya: 65

note: 65 itu outputnya dalam bentuk byte jadi ketika ingin outputnya a untuk karakter pertama maka harus konversi tipe datanya dulu ke string

fmt.Println(string("asep sahrudin"[0])) // outputnya: a
```

# 3. mengecek limit tipe data.

di Go, kita bisa mengecek batas maksimum dan minimum dari tipe data (seperti int, int8, int16, int32, int64) dengan bantuan package math.

contoh untuk mengecek limit int.

```
fmt.Println("int8 :", math.MinInt8, "sampai", math.MaxInt8)
fmt.Println("int16 :", math.MinInt16, "sampai", math.MaxInt16)
fmt.Println("int32 :", math.MinInt32, "sampai", math.MaxInt32)
fmt.Println("int64 :", math.MinInt64, "sampai", math.MaxInt64)
```

contoh untuk mengecek limit uint.

```
fmt.Println("cek limit tipe data uint")
fmt.Println("uint8 :", 0, "sampai", math.MaxUint8)
fmt.Println("uint16 :", 0, "sampai", math.MaxUint16)
fmt.Println("uint32 :", 0, "sampai", math.MaxUint32)
fmt.Println("uint64 :", 0, "sampai",uint64(math.MaxUint64)) // harus pakai unit64 biar aman.
```

✅ Kapan Pakai int?
Gunakan int jika:

- Nilai bisa positif atau negatif
- Melibatkan pengurangan atau perbandingan negatif
- Umum untuk index, counter, perhitungan matematis
- Ingin portabilitas (default int menyesuaikan arsitektur 32/64-bit)

contoh:

```
saldo := 1000;
saldo -= 1200; // bisa jadi negatif.

output: -200
```
