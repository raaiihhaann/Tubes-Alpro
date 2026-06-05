package main

import "fmt"

// ======================= MENU HANDLERS =======================

func menuCRUDSpace(daftar *DaftarSpace) {
	for {
		clearScreen()
		fmt.Println("\n  === KELOLA CO-WORKING SPACE ===")
		fmt.Println("  1. Tambah co-working space")
		fmt.Println("  2. Ubah co-working space")
		fmt.Println("  3. Hapus co-working space")
		fmt.Println("  4. Lihat semua co-working space")
		fmt.Println("  0. Kembali")
		pilihan := bacaInt("  Pilihan: ")

		switch pilihan {
		case 1:
			TambahWorkingSpace(daftar)
		case 2:
			UbahWorkingSpace(daftar)
		case 3:
			HapusWorkingSpace(daftar)
		case 4:
			tampilSemuaWorkingSpace(daftar)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 0:
			return
		default:
			fmt.Println("  Pilihan tidak valid.")
		}
	}
}

func menuCRUDFeedback(daftar *DaftarSpace) {
	for {
		clearScreen()
		fmt.Println("\n  === KELOLA FEEDBACK ===")
		fmt.Println("  1. Tambah feedback")
		fmt.Println("  2. Ubah feedback")
		fmt.Println("  3. Hapus feedback")
		fmt.Println("  4. Lihat feedback suatu tempat")
		fmt.Println("  0. Kembali")
		pilihan := bacaInt("  Pilihan: ")

		switch pilihan {
		case 1:
			TambahFeedback(daftar)
		case 2:
			UbahFeedback(daftar)
		case 3:
			HapusFeedback(daftar)
		case 4:
			tampilFeedback(daftar)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 0:
			return
		default:
			fmt.Println("  Pilihan tidak valid.")
		}
	}
}

func menuPencarian(daftar *DaftarSpace) {
	for {
		clearScreen()
		fmt.Println("\n  === PENCARIAN ===")
		fmt.Println("  1. Sequential Search berdasarkan nama")
		fmt.Println("  2. Sequential Search berdasarkan lokasi")
		fmt.Println("  3. Sequential Search berdasarkan fasilitas")
		fmt.Println("  4. Binary Search berdasarkan nama (exact)")
		fmt.Println("  5. Binary Search berdasarkan lokasi (exact)")
		fmt.Println("  0. Kembali")
		pilihan := bacaInt("  Pilihan: ")

		switch pilihan {
		case 1:
			keyword := bacaString("  Kata kunci nama: ")
			indices, jml := SeqSearchByNama(daftar, keyword)
			tampilHasilSearch(daftar, indices, jml)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 2:
			keyword := bacaString("  Kata kunci lokasi: ")
			indices, jml := SeqSearchByLokasi(daftar, keyword)
			tampilHasilSearch(daftar, indices, jml)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 3:
			keyword := bacaString("  Nama fasilitas (mis: WiFi, meeting room): ")
			indices, jml := SeqSearchByFasilitas(daftar, keyword)
			tampilHasilSearch(daftar, indices, jml)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 4:
			target := bacaString("  Nama tepat (exact): ")
			salinan := *daftar
			SortByNamaForBinary(&salinan)
			idx := BinarySearchByNama(&salinan, target)
			if idx == -1 {
				fmt.Println("  Tidak ditemukan.")
			} else {
				tampilWorkingSpace(salinan.Data[idx])
			}
			bacaString("\n  Tekan Enter untuk kembali...")
		case 5:
			target := bacaString("  Lokasi tepat (exact): ")
			salinan := *daftar
			SortByLokasiForBinary(&salinan)
			idx := BinarySearchByLokasi(&salinan, target)
			if idx == -1 {
				fmt.Println("  Tidak ditemukan.")
			} else {
				tampilWorkingSpace(salinan.Data[idx])
			}
			bacaString("\n  Tekan Enter untuk kembali...")
		case 0:
			return
		default:
			fmt.Println("  Pilihan tidak valid.")
		}
	}
}

func menuPengurutan(daftar *DaftarSpace) {
	for {
		clearScreen()
		fmt.Println("\n  === PENGURUTAN ===")
		fmt.Println("  1. Selection Sort: harga sewa (termurah ke termahal)")
		fmt.Println("  2. Insertion Sort: harga sewa (termurah ke termahal)")
		fmt.Println("  3. Selection Sort: rating tertinggi")
		fmt.Println("  4. Insertion Sort: rating tertinggi")
		fmt.Println("  0. Kembali")
		pilihan := bacaInt("  Pilihan: ")

		switch pilihan {
		case 1:
			salinan := *daftar
			SelectionSortByHarga(&salinan)
			fmt.Println("  [Selection Sort] Diurutkan berdasarkan harga (termurah ke termahal):")
			tampilSemuaWorkingSpace(&salinan)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 2:
			salinan := *daftar
			InsertionSortByHarga(&salinan)
			fmt.Println("  [Insertion Sort] Diurutkan berdasarkan harga (termurah ke termahal):")
			tampilSemuaWorkingSpace(&salinan)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 3:
			salinan := *daftar
			SelectionSortByRating(&salinan)
			fmt.Println("  [Selection Sort] Diurutkan berdasarkan rating tertinggi:")
			tampilSemuaWorkingSpace(&salinan)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 4:
			salinan := *daftar
			InsertionSortByRating(&salinan)
			fmt.Println("  [Insertion Sort] Diurutkan berdasarkan rating tertinggi:")
			tampilSemuaWorkingSpace(&salinan)
			bacaString("\n  Tekan Enter untuk kembali...")
		case 0:
			return
		default:
			fmt.Println("  Pilihan tidak valid.")
		}
	}
}