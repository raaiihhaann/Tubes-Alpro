package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// clearScreen: membersihkan layar terminal
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// ======================= DISPLAY =======================

// tampilWorkingSpace: menampilkan detail satu co-working space
func tampilWorkingSpace(ws WorkingSpace) {
	rataRating := hitungRataRating(ws)
	fmt.Println("  ─────────────────────────────────────────")
	fmt.Printf("  ID        : %d\n", ws.ID)
	fmt.Printf("  Nama      : %s\n", ws.Nama)
	fmt.Printf("  Lokasi    : %s\n", ws.Lokasi)
	fmt.Printf("  Harga     : Rp%d/hari\n", ws.HargaSewa)
	fmt.Printf("  Rating    : %.1f/5.0 (%d ulasan)\n", rataRating, ws.JmlFeedback)
	fmt.Print("  Fasilitas : ")
	for i := 0; i < ws.JmlFas; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(ws.Fasilitas[i])
	}
	fmt.Println()
}

// tampilSemuaWorkingSpace: menampilkan seluruh daftar co-working space
func tampilSemuaWorkingSpace(daftar *DaftarSpace) {
	if daftar.N == 0 {
		fmt.Println("  Belum ada data co-working space.")
		return
	}
	fmt.Printf("  Total: %d co-working space\n", daftar.N)
	for i := 0; i < daftar.N; i++ {
		tampilWorkingSpace(daftar.Data[i])
	}
	fmt.Println("  ─────────────────────────────────────────")
}

// tampilHasilSearch: menampilkan hasil pencarian berdasarkan slice index
func tampilHasilSearch(daftar *DaftarSpace, indices []int, jumlah int) {
	if jumlah == 0 {
		fmt.Println("  Tidak ditemukan hasil yang cocok.")
		return
	}
	fmt.Printf("  Ditemukan %d hasil:\n", jumlah)
	for _, idx := range indices {
		tampilWorkingSpace(daftar.Data[idx])
	}
	fmt.Println("  ─────────────────────────────────────────")
}

// tampilFeedback: menampilkan semua feedback dari satu co-working space
func tampilFeedback(daftar *DaftarSpace) {
	id := uint8(bacaInt("  Masukkan ID co-working space: "))
	idx := cariIndexByID(daftar, id)
	if idx == -1 {
		fmt.Println("  ID tidak ditemukan.")
		return
	}
	ws := daftar.Data[idx]
	fmt.Printf("  Feedback untuk: %s\n", ws.Nama)
	if ws.JmlFeedback == 0 {
		fmt.Println("  Belum ada feedback.")
		return
	}
	for i := 0; i < ws.JmlFeedback; i++ {
		rating := ws.Feedback[i].Rating
		bintang := ""
		for s := uint8(0); s < rating; s++ {
			bintang += "★"
		}
		for s := rating; s < 5; s++ {
			bintang += "☆"
		}
		fmt.Printf("  [%d] %s [%d/5]\n      \"%s\"\n", i+1, bintang, rating, ws.Feedback[i].Ulasan)
	}
}