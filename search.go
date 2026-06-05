package main

import "strings"

// ======================= SEQUENTIAL SEARCH =======================

// SeqSearchByNama: mencari semua co-working space yang mengandung kata kunci pada nama
// Mengembalikan array index yang cocok
func SeqSearchByNama(daftar *DaftarSpace, keyword string) ([]int, int) {
	hasil := make([]int, 0)
	keyword = strings.ToLower(keyword)
	for i := 0; i < daftar.N; i++ {
		if strings.Contains(strings.ToLower(daftar.Data[i].Nama), keyword) {
			hasil = append(hasil, i)
		}
	}
	return hasil, len(hasil)
}

// SeqSearchByLokasi: mencari semua co-working space berdasarkan lokasi
func SeqSearchByLokasi(daftar *DaftarSpace, keyword string) ([]int, int) {
	hasil := make([]int, 0)
	keyword = strings.ToLower(keyword)
	for i := 0; i < daftar.N; i++ {
		if strings.Contains(strings.ToLower(daftar.Data[i].Lokasi), keyword) {
			hasil = append(hasil, i)
		}
	}
	return hasil, len(hasil)
}

// SeqSearchByFasilitas: mencari co-working space yang memiliki fasilitas tertentu
func SeqSearchByFasilitas(daftar *DaftarSpace, fasilitas string) ([]int, int) {
	hasil := make([]int, 0)
	fasilitas = strings.ToLower(fasilitas)
	for i := 0; i < daftar.N; i++ {
		for f := 0; f < daftar.Data[i].JmlFas; f++ {
			if strings.Contains(strings.ToLower(daftar.Data[i].Fasilitas[f]), fasilitas) {
				hasil = append(hasil, i)
				break
			}
		}
	}
	return hasil, len(hasil)
}

// ======================= BINARY SEARCH =======================

// BinarySearchByNama: mencari co-working space berdasarkan nama exact match (case-insensitive)
// PENTING: data harus sudah diurutkan berdasarkan Nama secara ascending sebelum memanggil fungsi ini
// Mengembalikan index pertama yang ditemukan, atau -1 jika tidak ada
func BinarySearchByNama(daftar *DaftarSpace, target string) int {
	target = strings.ToLower(target)
	low := 0
	high := daftar.N - 1

	for low <= high {
		mid := (low + high) / 2
		midNama := strings.ToLower(daftar.Data[mid].Nama)
		if midNama == target {
			return mid
		} else if midNama < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// BinarySearchByLokasi: mencari co-working space berdasarkan lokasi exact match (case-insensitive)
// PENTING: data harus sudah diurutkan berdasarkan Lokasi secara ascending sebelum memanggil fungsi ini
func BinarySearchByLokasi(daftar *DaftarSpace, target string) int {
	target = strings.ToLower(target)
	low := 0
	high := daftar.N - 1

	for low <= high {
		mid := (low + high) / 2
		midLokasi := strings.ToLower(daftar.Data[mid].Lokasi)
		if midLokasi == target {
			return mid
		} else if midLokasi < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// SortByNamaForBinary: mengurutkan data berdasarkan Nama (ascending) untuk keperluan binary search
func SortByNamaForBinary(daftar *DaftarSpace) {
	n := daftar.N
	for i := 1; i < n; i++ {
		key := daftar.Data[i]
		j := i - 1
		for j >= 0 && strings.ToLower(daftar.Data[j].Nama) > strings.ToLower(key.Nama) {
			daftar.Data[j+1] = daftar.Data[j]
			j--
		}
		daftar.Data[j+1] = key
	}
}

// SortByLokasiForBinary: mengurutkan data berdasarkan Lokasi (ascending) untuk keperluan binary search
func SortByLokasiForBinary(daftar *DaftarSpace) {
	n := daftar.N
	for i := 1; i < n; i++ {
		key := daftar.Data[i]
		j := i - 1
		for j >= 0 && strings.ToLower(daftar.Data[j].Lokasi) > strings.ToLower(key.Lokasi) {
			daftar.Data[j+1] = daftar.Data[j]
			j--
		}
		daftar.Data[j+1] = key
	}
}