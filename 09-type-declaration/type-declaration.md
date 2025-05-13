# type declaration

Tipe alias memungkinkan kamu memberi alias (nama lain) untuk tipe yang sudah ada. Ini tidak membuat tipe baru, hanya mengganti nama tipe tersebut.

contohnya gini:

```
type Age int   // Age sekarang adalah alias dari tipe int

atau

type NoKTP string // NoKTP sekarang adalah alias dari tipe string
```

```
type NoKTP int64

var ktpSahrudin NoKTP = 24223232523;
fmt.Println(ktpSahrudin)
```
