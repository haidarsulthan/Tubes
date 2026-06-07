package main

import (
	"fmt"
)

type Lapangan struct {
	id    int
	nama  string
	tipe  string
	harga float64
}

type Penyewa struct {
	id   int
	nama string
	telp string
}

type Jadwal struct {
	id         int
	lapanganID int
	penyewaID  int
	tanggal    string
	jamMulai   int
	jamSelesai int
	harga      float64
}

var lapangans []Lapangan
var penyewas []Penyewa
var jadwals []Jadwal

var idLapangan = 1
var idPenyewa = 1
var idJadwal = 1

func inputString(prompt string) string {
	fmt.Print(prompt)
	var s string
	fmt.Scanln(&s)
	return s
}

func inputAngka(prompt string) int {
	fmt.Print(prompt)
	var n int
	fmt.Scanln(&n)
	return n
}

func inputFloat(prompt string) float64 {
	fmt.Print(prompt)
	var f float64
	fmt.Scanln(&f)
	return f
}

func garis() {
	fmt.Println("--------------------------------------------------")
}

func cariNamaLapangan(id int) string {
	for _, l := range lapangans {
		if l.id == id {
			return l.nama
		}
	}
	return "?"
}

func cariNamaPenyewa(id int) string {
	for _, p := range penyewas {
		if p.id == id {
			return p.nama
		}
	}
	return "?"
}

func lihatLapangan() {
	fmt.Println("\n=== DATA LAPANGAN ===")
	if len(lapangans) == 0 {
		fmt.Println("Belum ada data lapangan.")
		return
	}
	for _, l := range lapangans {
		fmt.Printf("ID: %d | %s | %s | Rp%.0f/jam\n", l.id, l.nama, l.tipe, l.harga)
	}
}

func tambahLapangan() {
	fmt.Println("\n=== TAMBAH LAPANGAN ===")
	nama := inputString("Nama lapangan : ")
	tipe := inputString("Tipe (indoor/outdoor) : ")
	harga := inputFloat("Harga per jam : ")

	lapangans = append(lapangans, Lapangan{idLapangan, nama, tipe, harga})
	fmt.Printf("Lapangan '%s' berhasil ditambah! (ID: %d)\n", nama, idLapangan)
	idLapangan++
}

func ubahLapangan() {
	lihatLapangan()
	id := inputAngka("\nMasukkan ID lapangan yang mau diubah: ")

	for i, l := range lapangans {
		if l.id == id {
			fmt.Printf("Data lama: %s | %s | Rp%.0f\n", l.nama, l.tipe, l.harga)
			nama := inputString("Nama baru : ")
			tipe := inputString("Tipe baru : ")
			harga := inputFloat("Harga baru : ")

			lapangans[i].nama = nama
			lapangans[i].tipe = tipe
			lapangans[i].harga = harga
			fmt.Println("Data lapangan berhasil diubah!")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func hapusLapangan() {
	lihatLapangan()
	id := inputAngka("\nMasukkan ID lapangan yang mau dihapus: ")

	for i, l := range lapangans {
		if l.id == id {
			lapangans = append(lapangans[:i], lapangans[i+1:]...)
			fmt.Printf("Lapangan '%s' berhasil dihapus!\n", l.nama)
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func lihatPenyewa() {
	fmt.Println("\n=== DATA PENYEWA ===")
	if len(penyewas) == 0 {
		fmt.Println("Belum ada data penyewa.")
		return
	}
	for _, p := range penyewas {
		fmt.Printf("ID: %d | %s | %s\n", p.id, p.nama, p.telp)
	}
}

func tambahPenyewa() {
	fmt.Println("\n=== TAMBAH PENYEWA ===")
	nama := inputString("Nama penyewa : ")
	telp := inputString("No. telepon  : ")

	penyewas = append(penyewas, Penyewa{idPenyewa, nama, telp})
	fmt.Printf("Penyewa '%s' berhasil ditambah! (ID: %d)\n", nama, idPenyewa)
	idPenyewa++
}

func ubahPenyewa() {
	lihatPenyewa()
	id := inputAngka("\nMasukkan ID penyewa yang mau diubah: ")

	for i, p := range penyewas {
		if p.id == id {
			fmt.Printf("Data lama: %s | %s\n", p.nama, p.telp)
			nama := inputString("Nama baru : ")
			telp := inputString("Telepon baru : ")

			penyewas[i].nama = nama
			penyewas[i].telp = telp
			fmt.Println("Data penyewa berhasil diubah!")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func hapusPenyewa() {
	lihatPenyewa()
	id := inputAngka("\nMasukkan ID penyewa yang mau dihapus: ")

	for i, p := range penyewas {
		if p.id == id {
			penyewas = append(penyewas[:i], penyewas[i+1:]...)
			fmt.Printf("Penyewa '%s' berhasil dihapus!\n", p.nama)
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func tambahJadwal() {
	fmt.Println("\n=== CATAT SEWA LAPANGAN ===")

	if len(lapangans) == 0 || len(penyewas) == 0 {
		fmt.Println("Pastikan data lapangan dan penyewa sudah ada dulu.")
		return
	}

	lihatLapangan()
	lapID := inputAngka("\nPilih ID lapangan: ")
	var lap *Lapangan
	for i := range lapangans {
		if lapangans[i].id == lapID {
			lap = &lapangans[i]
		}
	}
	if lap == nil {
		fmt.Println("Lapangan tidak ditemukan.")
		return
	}

	lihatPenyewa()
	pID := inputAngka("\nPilih ID penyewa: ")
	var penyewa *Penyewa
	for i := range penyewas {
		if penyewas[i].id == pID {
			penyewa = &penyewas[i]
		}
	}
	if penyewa == nil {
		fmt.Println("Penyewa tidak ditemukan.")
		return
	}

	tanggal := inputString("Tanggal (contoh: 2025-06-10): ")
	jamMulai := inputAngka("Jam mulai (0-23): ")
	jamSelesai := inputAngka("Jam selesai (0-23): ")

	if jamSelesai <= jamMulai {
		fmt.Println("Jam selesai harus lebih besar dari jam mulai!")
		return
	}

	for _, j := range jadwals {
		if j.lapanganID == lapID && j.tanggal == tanggal {
			if jamMulai < j.jamSelesai && jamSelesai > j.jamMulai {
				fmt.Printf("Jadwal bentrok dengan sewa jam %d:00-%d:00!\n", j.jamMulai, j.jamSelesai)
				return
			}
		}
	}

	durasi := jamSelesai - jamMulai
	total := float64(durasi) * lap.harga

	jadwals = append(jadwals, Jadwal{idJadwal, lapID, pID, tanggal, jamMulai, jamSelesai, total})
	fmt.Println("\nSewa berhasil dicatat!")
	fmt.Printf("Lapangan : %s\n", lap.nama)
	fmt.Printf("Penyewa  : %s\n", penyewa.nama)
	fmt.Printf("Tanggal  : %s, jam %d:00 - %d:00\n", tanggal, jamMulai, jamSelesai)
	fmt.Printf("Total    : Rp%.0f\n", total)
	idJadwal++
}

func lihatJadwal() {
	fmt.Println("\n=== JADWAL SEWA ===")
	if len(jadwals) == 0 {
		fmt.Println("Belum ada jadwal.")
		return
	}
	for _, j := range jadwals {
		fmt.Printf("ID:%d | %s | %s | %s | %d:00-%d:00 | Rp%.0f\n",
			j.id, j.tanggal, cariNamaLapangan(j.lapanganID), cariNamaPenyewa(j.penyewaID),
			j.jamMulai, j.jamSelesai, j.harga)
	}
}

func sequentialSearch(keyword string) {
	fmt.Println("\n[Sequential Search]")
	ketemu := false

	for _, p := range penyewas {
		if p.nama == keyword || p.telp == keyword {
			fmt.Printf("Ketemu! ID: %d | %s | %s\n", p.id, p.nama, p.telp)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan.")
	}
}

func binarySearch(keyword string) {
	fmt.Println("\n[Binary Search]")

	sorted := make([]Penyewa, len(penyewas))
	copy(sorted, penyewas)
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-1-i; j++ {
			if sorted[j].nama > sorted[j+1].nama {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	kiri := 0
	kanan := len(sorted) - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if sorted[tengah].nama == keyword {
			p := sorted[tengah]
			fmt.Printf("Ketemu! ID: %d | %s | %s\n", p.id, p.nama, p.telp)
			return
		} else if sorted[tengah].nama < keyword {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func menuCari() {
	fmt.Println("\n=== CARI PENYEWA ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	pilihan := inputAngka("Pilih: ")
	keyword := inputString("Masukkan keyword: ")

	if pilihan == 1 {
		sequentialSearch(keyword)
	} else if pilihan == 2 {
		binarySearch(keyword)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func selectionSortJam() {
	fmt.Println("\n[Selection Sort - urut berdasarkan jam mulai]")

	data := make([]Jadwal, len(jadwals))
	copy(data, jadwals)

	for i := 0; i < len(data)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(data); j++ {
			if data[j].jamMulai < data[minIdx].jamMulai {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}

	for _, j := range data {
		fmt.Printf("ID:%d | %s | jam %d:00-%d:00 | Rp%.0f\n",
			j.id, cariNamaLapangan(j.lapanganID), j.jamMulai, j.jamSelesai, j.harga)
	}
}

func insertionSortHarga() {
	fmt.Println("\n[Insertion Sort - urut berdasarkan harga]")

	data := make([]Jadwal, len(jadwals))
	copy(data, jadwals)

	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && data[j].harga > temp.harga {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}

	for _, j := range data {
		fmt.Printf("ID:%d | %s | jam %d:00-%d:00 | Rp%.0f\n",
			j.id, cariNamaLapangan(j.lapanganID), j.jamMulai, j.jamSelesai, j.harga)
	}
}

func menuSort() {
	fmt.Println("\n=== URUTKAN JADWAL ===")
	fmt.Println("1. Selection Sort (berdasarkan jam mulai)")
	fmt.Println("2. Insertion Sort (berdasarkan harga)")
	pilihan := inputAngka("Pilih: ")

	if len(jadwals) == 0 {
		fmt.Println("Belum ada jadwal.")
		return
	}

	if pilihan == 1 {
		selectionSortJam()
	} else if pilihan == 2 {
		insertionSortHarga()
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func tampilkanStatistik() {
	fmt.Println()
	fmt.Println("+++ FUTSAL-BOOK +++")
	fmt.Println("=== STATISTIK ===")

	if len(jadwals) == 0 {
		fmt.Println("Belum ada data transaksi.")
		return
	}

	var bulanList []string
	var totalList []float64

	for _, j := range jadwals {
		bulan := j.tanggal[:7]

		sudahAda := false
		for i, b := range bulanList {
			if b == bulan {
				totalList[i] += j.harga
				sudahAda = true
				break
			}
		}
		if !sudahAda {
			bulanList = append(bulanList, bulan)
			totalList = append(totalList, j.harga)
		}
	}

	fmt.Println("\n-- Pendapatan per Bulan --")
	totalSemua := 0.0
	for i, b := range bulanList {
		fmt.Printf("%s : Rp%.0f\n", b, totalList[i])
		totalSemua += totalList[i]
	}
	fmt.Printf("Total keseluruhan : Rp%.0f\n", totalSemua)

	var frekuensiJam [24]int
	for _, j := range jadwals {
		for jam := j.jamMulai; jam < j.jamSelesai; jam++ {
			frekuensiJam[jam]++
		}
	}

	jamTerpopuler := 0
	for i := 1; i < 24; i++ {
		if frekuensiJam[i] > frekuensiJam[jamTerpopuler] {
			jamTerpopuler = i
		}
	}

	fmt.Println("\n-- Jam Paling Sering Dipesan --")
	fmt.Printf("Jam %d:00 - %d:00 dipesan sebanyak %d kali\n",
		jamTerpopuler, jamTerpopuler+1, frekuensiJam[jamTerpopuler])

	fmt.Println("\nDetail per jam:")
	for i := 6; i < 22; i++ {
		if frekuensiJam[i] > 0 {
			fmt.Printf("Jam %02d:00 - dipesan %d kali\n", i, frekuensiJam[i])
		}
	}
}

func isiDataContoh() {
	lapangans = []Lapangan{
		{1, "LapanganA", "indoor", 100000},
		{2, "LapanganB", "outdoor", 75000},
	}
	idLapangan = 3

	penyewas = []Penyewa{
		{1, "Andi", "081111111111"},
		{2, "Budi", "082222222222"},
		{3, "Citra", "083333333333"},
	}
	idPenyewa = 4

	jadwals = []Jadwal{
		{1, 1, 1, "2025-06-01", 8, 10, 200000},
		{2, 1, 2, "2025-06-01", 14, 16, 200000},
		{3, 2, 3, "2025-06-02", 9, 11, 150000},
		{4, 1, 1, "2025-06-05", 8, 10, 200000},
	}
	idJadwal = 5
}

func main() {
	isiDataContoh()

	fmt.Println("+++ FUTSAL-BOOK +++")
	fmt.Println("Aplikasi Pemesanan Lapangan Futsal")
	garis()

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1.  Lihat lapangan")
		fmt.Println("2.  Tambah lapangan")
		fmt.Println("3.  Ubah lapangan")
		fmt.Println("4.  Hapus lapangan")
		fmt.Println("5.  Lihat penyewa")
		fmt.Println("6.  Tambah penyewa")
		fmt.Println("7.  Ubah penyewa")
		fmt.Println("8.  Hapus penyewa")
		fmt.Println("9.  Catat sewa")
		fmt.Println("10. Lihat jadwal sewa")
		fmt.Println("11. Cari penyewa")
		fmt.Println("12. Urutkan jadwal")
		fmt.Println("13. Statistik")
		fmt.Println("0.  Keluar")
		garis()

		pilihan := inputAngka("Pilih menu: ")

		switch pilihan {
		case 1:
			lihatLapangan()
		case 2:
			tambahLapangan()
		case 3:
			ubahLapangan()
		case 4:
			hapusLapangan()
		case 5:
			lihatPenyewa()
		case 6:
			tambahPenyewa()
		case 7:
			ubahPenyewa()
		case 8:
			hapusPenyewa()
		case 9:
			tambahJadwal()
		case 10:
			lihatJadwal()
		case 11:
			menuCari()
		case 12:
			menuSort()
		case 13:
			tampilkanStatistik()
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Menu tidak ada.")
		}
	}
}
