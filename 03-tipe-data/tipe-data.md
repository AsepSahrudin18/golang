# tipe data number

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
