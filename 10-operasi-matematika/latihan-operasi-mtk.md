# operasi matematika

singkatnya operasi matematika untuk membuat hitungan matematis seperti di bahasa pemrograman lain.

📌 Ringkasan Operator Dasar di Go
| Operator | Fungsi | Contoh |
|-----------|------------------|--------------|
| + | Penjumlahan | a + b |
| - | Pengurangan | a - b |
| _ | Perkalian | a _ b |
| / | Pembagian | a / b |
| % | Modulus (sisa) | a % b |

📌 Augmented Assignments
seperti di bahasa lain, di go juga singkatnya ada shorthand untuk cara penulisan singkat matematika.

contoh.

```
var a = 10
var b = 20

var hasil = a + b;
fmt.Println(hasil)

// dengan shorthand (cara singkat) atau disebut penulisan Augmented Assignments
var nilaiA  = 40;
nilaiA += 10;
fmt.Println(nilaiA);
```

## unary operator

Unary operator adalah operator yang hanya bekerja pada satu operand.

🔹 Jenis Unary Operator di Go:
| Operator | Fungsi | Contoh |  
|-----------|-------------------------------|-------------------|
| + | menunjukkan nilai positif | +a |
| - | Negasi (nilai negatif) | -a |
| ! | Logika NOT (kebalikan) | !true → false |
| ^ | Bitwise NOT (komplemen) | ^a |
| * | Pointer dereference | *ptr |
| & | Ambil alamat (address-of) | &a |

contoh increment dan decrement:

```
// contoh increment
var increment = 0;
increment++;
fmt.Println(increment); // output: 1

// contoh decrement.
var decrement = 0;
decrement--;
fmt.Println(decrement); // output: -1
```
