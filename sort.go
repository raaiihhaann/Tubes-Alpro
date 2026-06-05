package main

// ======================= SORT BY HARGA =======================

// SelectionSortByHarga: mengurutkan ascending berdasarkan HargaSewa
func SelectionSortByHarga(arr *DaftarSpace) {
	n := arr.N
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr.Data[j].HargaSewa < arr.Data[minIndex].HargaSewa {
				minIndex = j
			}
		}
		arr.Data[i], arr.Data[minIndex] = arr.Data[minIndex], arr.Data[i]
	}
}

// InsertionSortByHarga: mengurutkan ascending berdasarkan HargaSewa
func InsertionSortByHarga(arr *DaftarSpace) {
	n := arr.N
	for i := 1; i < n; i++ {
		key := arr.Data[i]
		j := i - 1
		for j >= 0 && arr.Data[j].HargaSewa > key.HargaSewa {
			arr.Data[j+1] = arr.Data[j]
			j--
		}
		arr.Data[j+1] = key
	}
}

// ======================= SORT BY RATING =======================

// hitungRataRating: menghitung rata-rata rating dari semua feedback
func hitungRataRating(ws WorkingSpace) float64 {
	if ws.JmlFeedback == 0 {
		return 0
	}
	total := 0
	for i := 0; i < ws.JmlFeedback; i++ {
		total += int(ws.Feedback[i].Rating)
	}
	return float64(total) / float64(ws.JmlFeedback)
}

// SelectionSortByRating: mengurutkan descending berdasarkan rata-rata rating
func SelectionSortByRating(arr *DaftarSpace) {
	n := arr.N
	for i := 0; i < n-1; i++ {
		maxIndex := i
		for j := i + 1; j < n; j++ {
			if hitungRataRating(arr.Data[j]) > hitungRataRating(arr.Data[maxIndex]) {
				maxIndex = j
			}
		}
		arr.Data[i], arr.Data[maxIndex] = arr.Data[maxIndex], arr.Data[i]
	}
}

// InsertionSortByRating: mengurutkan descending berdasarkan rata-rata rating
func InsertionSortByRating(arr *DaftarSpace) {
	n := arr.N
	for i := 1; i < n; i++ {
		key := arr.Data[i]
		keyRating := hitungRataRating(key)
		j := i - 1
		for j >= 0 && hitungRataRating(arr.Data[j]) < keyRating {
			arr.Data[j+1] = arr.Data[j]
			j--
		}
		arr.Data[j+1] = key
	}
}