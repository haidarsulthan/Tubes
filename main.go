package main
import "fmt"

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
			tipe := inputString("Tipe baru(Indoor/Outdoor) : ")
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

func dataPenyewa() {
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
	dataPenyewa()
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
	dataPenyewa()
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
	fmt.Println("\n=== TAMBAH SEWA LAPANGAN ===")

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

func main() {

	fmt.Println("Aplikasi Pemesanan Lapangan Futsal")
	fmt.Println("--------------------------------------------------")

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("--------------------------------------------------")
		fmt.Println("1.  Tambah lapangan")
		fmt.Println("2.  Ubah lapangan")
		fmt.Println("3.  Hapus lapangan")
		fmt.Println("4.  Tambah penyewa")
		fmt.Println("5.  Ubah penyewa")
		fmt.Println("6.  Hapus penyewa")
		fmt.Println("7.  Tambah sewa")
		fmt.Println("8.  Lihat jadwal sewa")
		fmt.Println("0.  Keluar")
		fmt.Println("--------------------------------------------------")

		pilihan := inputAngka("Pilih menu: ")

		switch pilihan {
		case 1:
			tambahLapangan()
		case 2:
			ubahLapangan()
		case 3:
			hapusLapangan()
		case 4:
			tambahPenyewa()
		case 5:
			ubahPenyewa()
		case 6:
			hapusPenyewa()
		case 7:
			tambahJadwal()
		case 8:
			lihatJadwal()
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Menu tidak ada.")
		}
	}
}
