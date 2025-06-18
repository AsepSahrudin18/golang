## switch expression

digunakan untuk mengecek isi dalam konidisi dalam satu variable saja.

dalam swicth case di go juga mendukung short statement:

```
    username := "sahrudin18asep";
	switch user := len(username); user > 5 {
	case true:
	fmt.Println("hallo sahrudin");
	case false:
		fmt.Println("hallo user");
	default:
		fmt.Println("hallo boleh kenalan?")
	}
```
