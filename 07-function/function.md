# function

function merupakan bagian penting dalam pemrograman golang, fungsi membantu kita memisahkan kode ke dalam blok-blok kecil yang bisa digunakan kembali, membuat program jadi lebih rapi, mudah dibaca, dan mudah dikelola.

### cara membuat fungsi di golang.

di golang untuk mendeklarasikan fungsi dengan keyword func, diikuti dengan nama fungsi, daftar parameter, dan tipe nilai yang dikembalikan (jika ada).

#### 1. fungsi sederhana tanpa paramter dan tanpa return.

```
func sapa() {
	fmt.Println("Hallo dari fungsi!");
}
```

#### 2. function dengan return bernama

```
func bagi(a, b int) (hasil int) {
	hasil = a / b;
	return // cukup gunakan return tanpa argument
}
```

Kapan sebaiknya digunakan?
Cocok untuk fungsi pendek atau yang nilainya jelas.

Tidak direkomendasikan kalau membuat kode jadi tidak eksplisit (misalnya: banyak variabel, bisa bikin bingung).

#### 3. function dengan paramter

```
func hello(name string) {
	fmt.Println("hai apa kabar? ",name)
}
```

#### 4. function dengan return parameter

```
func question(name string) string {
	return "Bagaimana kabar kamu hari ini? " + name
}
```

#### 5. function dengan beberapa parameter

// note: pendefinisan tipe data sekali seperti ini jika casenya tipe data paramter sama

```
func tambah(a, b int) int {
	return a + b;
}
```

#### 6. jika paramter lebih dari satu dan tipe datanya berbeda

```
func info(name string, umur int) (string, int) {
	return name, umur;
}
```

#### 7. function dengan multiple return value

```
func hitung(a int, b int) (int, int) {
	return a + b, a * b;
}
```

#### 8. function dengan return bernama, kita juga bisa melakukan return dengan memberi nama nilai yang dikembalikan.

```
func bagi(a, b int) (hasil int) {
	hasil = a / b;
	return // cukup gunakan return tanpa argument
}
```
