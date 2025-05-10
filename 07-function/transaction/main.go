package main
import (
	"fmt"
	"strings"
	"github.com/dustin/go-humanize"
)

// aplikasi pencatatan pengeluaran
/** 
- category: pendapatan, pengeluaran, total
*/

var (
	gaji = 15000000;
	pengeluaran = 3500000;
);

func transaction(nominalPendapatan, nominalPengeluaran int) (totalNominal int) {
	totalNominal = nominalPendapatan - nominalPengeluaran;
	return;
}

func main() {
	hitungPengeluaran := transaction(gaji, pengeluaran);

	// number format
	formatNumPendapatan 	:= humanize.Comma(int64(gaji));
	formatNumPengeluaran 	:= humanize.Comma(int64(pengeluaran));
	formatNumTotalNilai		:= humanize.Comma(int64(hitungPengeluaran));

	// replace , comma to .
	pendapatanId 	:= strings.ReplaceAll(formatNumPendapatan, ",",".");
	pengeluaranId 	:= strings.ReplaceAll(formatNumPengeluaran, ",",".");
	totalNilaiId	:= strings.ReplaceAll(formatNumTotalNilai, ",",".");

	// info
	totalGaji 			:= "Pendapatan: " + "Rp. " + pendapatanId;
	totalPengeluaran 	:= "Pengeluaran: " + "Rp. " + pengeluaranId;
	total 				:= "Total Nominal: " + "Rp. " + totalNilaiId;
	fmt.Println(totalGaji);
	fmt.Println(totalPengeluaran);
	fmt.Println(total);
}