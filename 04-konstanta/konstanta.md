### Konstanta

konstanta merupakan variabel khusus yang nilainya tidak dapat diubah setelah dideklarasikan. Biasanya digunakan untuk nilai-nilai yang sudah pasti, seperti konfigurasi, batasan, atau data tetap lainnya.

cara mendeklarasikan konstanta didefinisikan dengan kata kunci const.

```
const phi = 3.14
const aplikasi = "Belajar Golang"
```

#### konstanta bertipe dan tanpa tipe

Golang memperbolehkan dua jenis konstanta:

1. konstanta tanpa tipe:

```
const angka = 10
```

2. konstanta dengan tipe:

```
const umur int = 20
```

#### grup konstanta

kita juga bisa mendeklarasikan beberapa konstanta sekaligus dalam satu blok.

```
const (
    satu = 1
    dua  = 2
    tiga = 3
)
```

#### keyword iota (untuk nilai berurutan)

golang menyediakan keyword iota untuk membuat urutan konstanta secara otomatis. biasanya digunakan dalam enum atau deklarasi bertingkat.
