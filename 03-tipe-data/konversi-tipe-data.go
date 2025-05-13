package main 
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("cek limit tipe data int")
	fmt.Println("int8 :", math.MinInt8, "sampai", math.MaxInt8)
	fmt.Println("int16 :", math.MinInt16, "sampai", math.MaxInt16)
	fmt.Println("int32 :", math.MinInt32, "sampai", math.MaxInt32)
	fmt.Println("int64 :", math.MinInt64, "sampai", math.MaxInt64)


	fmt.Println("cek limit tipe data uint")
	fmt.Println("uint8 :", 0, "sampai", math.MaxUint8)
	fmt.Println("uint16 :", 0, "sampai", math.MaxUint16)
	fmt.Println("uint32 :", 0, "sampai", math.MaxUint32)
	fmt.Println("uint64 :", 0, "sampai",uint64(math.MaxUint64)) // harus pakai unit64 biar aman.

	saldo := 1000;
	saldo -= 1200; // bisa jadi negatif.

	fmt.Println(saldo);

	var cekSaldo uint64 = 12000;
	cekSaldo -= 10000;
	fmt.Println(cekSaldo)
}

