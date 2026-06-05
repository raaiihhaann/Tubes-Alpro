package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// bacaString: membaca input string dari pengguna
func bacaString(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// bacaInt: membaca input integer dari pengguna, validasi bukan angka
func bacaInt(prompt string) int {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(" Input tidak valid. Masukkan angka.")
			continue
		}
		return val
	}
}

// bacaIntDenganValidasi: membaca integer dengan batas min-max, loop + partial clear saat invalid
// headerFunc dipanggil ulang setelah clear untuk menampilkan konteks yang sudah diisi
func bacaIntDenganValidasi(prompt string, min int, max int, headerFunc func()) int {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err != nil || val < min || val > max {
			fmt.Printf(" Input tidak valid. Masukkan angka antara %d dan %d.\n", min, max)
			fmt.Print("  Tekan Enter untuk coba lagi...")
			reader.ReadString('\n')
			clearScreen()
			headerFunc()
			continue
		}
		return val
	}
}

// nextID: menghasilkan ID baru berdasarkan ID tertinggi yang ada
// Mengembalikan 0 jika sudah overflow (tidak seharusnya terjadi karena MAX_SPACES=100 < 255)
func nextID(daftar *DaftarSpace) uint8 {
	maxID := uint8(0)
	for i := 0; i < daftar.N; i++ {
		if daftar.Data[i].ID > maxID {
			maxID = daftar.Data[i].ID
		}
	}
	if maxID == 255 {
		return 0
	}
	return maxID + 1
}

// bacaID: membaca input ID (1-255) dengan validasi range aman
func bacaID(prompt string) uint8 {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err != nil || val < 1 || val > 255 {
			fmt.Println(" Input tidak valid. Masukkan angka ID (1-255).")
			continue
		}
		return uint8(val)
	}
}


func cariIndexByID(daftar *DaftarSpace, id uint8) int {
	for i := 0; i < daftar.N; i++ {
		if daftar.Data[i].ID == id {
			return i
		}
	}
	return -1
}

// ======================= WORKING SPACE CRUD =======================

// TambahWorkingSpace: menambahkan data co-working space baru
func TambahWorkingSpace(daftar *DaftarSpace) {
	if daftar.N >= MAX_SPACES {
		fmt.Println("Daftar penuh! Tidak bisa menambah data.")
		return
	}

	var ws WorkingSpace
	ws.ID = nextID(daftar)

	for {
		ws.Nama = bacaString("Nama co-working space: ")
		if ws.Nama != "" {
			break
		}
		fmt.Println(" Nama tidak boleh kosong.")
	}
	for {
		ws.Lokasi = bacaString("Lokasi: ")
		if ws.Lokasi != "" {
			break
		}
		fmt.Println(" Lokasi tidak boleh kosong.")
	}

	// Header ulang saat validasi fasilitas gagal (nama & lokasi sudah terisi)
	headerFasilitas := func() {
		fmt.Println("\n  === TAMBAH CO-WORKING SPACE ===")
		fmt.Printf("Nama co-working space: %s\n", ws.Nama)
		fmt.Printf("Lokasi: %s\n", ws.Lokasi)
	}

	jmlFas := bacaIntDenganValidasi("Jumlah fasilitas (maks 10): ", 1, 10, headerFasilitas)

	for i := 0; i < jmlFas; i++ {
		ws.Fasilitas[i] = bacaString(fmt.Sprintf("  Fasilitas %d: ", i+1))
	}
	ws.JmlFas = jmlFas
	ws.HargaSewa = bacaInt("Harga sewa per hari (Rp): ")

	daftar.Data[daftar.N] = ws
	daftar.N++
	fmt.Printf("Co-working space '%s' berhasil ditambahkan dengan ID %d.\n", ws.Nama, ws.ID)
}

// UbahWorkingSpace: mengubah data co-working space berdasarkan ID
func UbahWorkingSpace(daftar *DaftarSpace) {
	id := bacaID("Masukkan ID co-working space yang ingin diubah: ")
	idx := cariIndexByID(daftar, id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	ws := &daftar.Data[idx]
	fmt.Printf("Mengubah data: %s\n", ws.Nama)
	fmt.Println("(Kosongkan input untuk tidak mengubah)")

	nama := bacaString(fmt.Sprintf("Nama baru [%s]: ", ws.Nama))
	if nama != "" {
		ws.Nama = nama
	}

	lokasi := bacaString(fmt.Sprintf("Lokasi baru [%s]: ", ws.Lokasi))
	if lokasi != "" {
		ws.Lokasi = lokasi
	}

	ubahFas := bacaString("Ubah fasilitas? (y/n): ")
	if strings.ToLower(ubahFas) == "y" {
		headerFasilitas := func() {
			fmt.Println("\n  === UBAH CO-WORKING SPACE ===")
			fmt.Printf("Mengubah data: %s\n", ws.Nama)
		}
		jmlFas := bacaIntDenganValidasi("Jumlah fasilitas (maks 10): ", 1, 10, headerFasilitas)
		for i := 0; i < jmlFas; i++ {
			ws.Fasilitas[i] = bacaString(fmt.Sprintf("  Fasilitas %d: ", i+1))
		}
		ws.JmlFas = jmlFas
	}

	// Gunakan bacaInt biasa; 0 = tidak ubah, validasi eksplisit
	hargaBaru := bacaInt(fmt.Sprintf("Harga sewa baru [%d] (0 = tidak ubah): ", ws.HargaSewa))
	if hargaBaru > 0 {
		ws.HargaSewa = hargaBaru
	} else if hargaBaru < 0 {
		fmt.Println("  Harga tidak diubah (nilai negatif diabaikan).")
	}

	fmt.Println("Data berhasil diubah.")
}

// HapusWorkingSpace: menghapus co-working space berdasarkan ID
func HapusWorkingSpace(daftar *DaftarSpace) {
	id := bacaID("Masukkan ID co-working space yang ingin dihapus: ")
	idx := cariIndexByID(daftar, id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	nama := daftar.Data[idx].Nama
	// geser elemen ke kiri
	for i := idx; i < daftar.N-1; i++ {
		daftar.Data[i] = daftar.Data[i+1]
	}
	daftar.N--
	fmt.Printf("Co-working space '%s' berhasil dihapus.\n", nama)
}

// ======================= FEEDBACK CRUD =======================

// TambahFeedback: menambahkan ulasan dan rating untuk co-working space
func TambahFeedback(daftar *DaftarSpace) {
	id := bacaID("Masukkan ID co-working space: ")
	idx := cariIndexByID(daftar, id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	ws := &daftar.Data[idx]
	if ws.JmlFeedback >= MAX_FEEDBACK {
		fmt.Println("Feedback untuk tempat ini sudah penuh.")
		return
	}

	var fb Feedback
	fb.Ulasan = bacaString("Ulasan Anda: ")
	namaWs := ws.Nama
	ulasanInput := fb.Ulasan
	headerRating := func() {
		fmt.Printf("\n  === TAMBAH FEEDBACK: %s ===\n", namaWs)
		fmt.Printf("  Ulasan: %s\n", ulasanInput)
	}
	fb.Rating = uint8(bacaIntDenganValidasi("Rating (1-5): ", 1, 5, headerRating))

	ws.Feedback[ws.JmlFeedback] = fb
	ws.JmlFeedback++
	fmt.Println("Feedback berhasil ditambahkan.")
}

// UbahFeedback: mengubah ulasan/rating berdasarkan index feedback
func UbahFeedback(daftar *DaftarSpace) {
	id := bacaID("Masukkan ID co-working space: ")
	idx := cariIndexByID(daftar, id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	ws := &daftar.Data[idx]
	if ws.JmlFeedback == 0 {
		fmt.Println("Belum ada feedback untuk tempat ini.")
		return
	}

	fmt.Printf("Tempat: %s | Total feedback: %d\n", ws.Nama, ws.JmlFeedback)
	for i := 0; i < ws.JmlFeedback; i++ {
		fmt.Printf("  [%d] Rating %d - %s\n", i+1, ws.Feedback[i].Rating, ws.Feedback[i].Ulasan)
	}

	no := bacaInt("Nomor feedback yang ingin diubah: ") - 1
	if no < 0 || no >= ws.JmlFeedback {
		fmt.Println("Nomor tidak valid.")
		return
	}

	ws.Feedback[no].Ulasan = bacaString("Ulasan baru: ")
	namaWs2 := ws.Nama
	ulasanBaru := ws.Feedback[no].Ulasan
	headerRatingUbah := func() {
		fmt.Printf("\n  === UBAH FEEDBACK: %s ===\n", namaWs2)
		fmt.Printf("  Ulasan baru: %s\n", ulasanBaru)
	}
	ws.Feedback[no].Rating = uint8(bacaIntDenganValidasi("Rating baru (1-5): ", 1, 5, headerRatingUbah))
	fmt.Println("Feedback berhasil diubah.")
}

// HapusFeedback: menghapus feedback berdasarkan index
func HapusFeedback(daftar *DaftarSpace) {
	id := bacaID("Masukkan ID co-working space: ")
	idx := cariIndexByID(daftar, id)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	ws := &daftar.Data[idx]
	if ws.JmlFeedback == 0 {
		fmt.Println("Belum ada feedback untuk tempat ini.")
		return
	}

	fmt.Printf("Tempat: %s | Total feedback: %d\n", ws.Nama, ws.JmlFeedback)
	for i := 0; i < ws.JmlFeedback; i++ {
		fmt.Printf("  [%d] Rating %d - %s\n", i+1, ws.Feedback[i].Rating, ws.Feedback[i].Ulasan)
	}

	no := bacaInt("Nomor feedback yang ingin dihapus: ") - 1
	if no < 0 || no >= ws.JmlFeedback {
		fmt.Println("Nomor tidak valid.")
		return
	}

	for i := no; i < ws.JmlFeedback-1; i++ {
		ws.Feedback[i] = ws.Feedback[i+1]
	}
	ws.JmlFeedback--
	fmt.Println("Feedback berhasil dihapus.")
}