package main

import "fmt"

func main() {
	var daftar DaftarSpace
	inisialisasiData(&daftar)

	for {
		clearScreen()
		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║    CO-WORKING SPACE MANAGEMENT SYSTEM    ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  1. Kelola Co-Working Space              ║")
		fmt.Println("║  2. Kelola Feedback & Rating             ║")
		fmt.Println("║  3. Pencarian                            ║")
		fmt.Println("║  4. Pengurutan / Tampil Daftar           ║")
		fmt.Println("║  5. Tampil Semua Co-Working Space        ║")
		fmt.Println("║  0. Keluar                               ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		pilihan := bacaInt("  Pilihan: ")

		switch pilihan {
		case 1:
			menuCRUDSpace(&daftar)
		case 2:
			menuCRUDFeedback(&daftar)
		case 3:
			menuPencarian(&daftar)
		case 4:
			menuPengurutan(&daftar)
		case 5:
			tampilSemuaWorkingSpace(&daftar)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 0:
			fmt.Println("  Terima kasih. Program selesai.")
			return
		default:
			fmt.Println("  Pilihan tidak valid.")
		}
	}
}