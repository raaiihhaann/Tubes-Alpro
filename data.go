package main

// inisialisasiData: mengisi daftar dengan data dummy awal untuk demonstrasi
func inisialisasiData(daftar *DaftarSpace) {
	daftar.Data[0] = WorkingSpace{
		ID:    1,
		Nama:  "SpaceHub Jakarta",
		Lokasi: "Jakarta Selatan",
		Fasilitas: [10]string{"WiFi", "Meeting Room", "Private Desk", "Cafeteria"},
		JmlFas:    4,
		HargaSewa: 150000,
		Feedback: [MAX_FEEDBACK]Feedback{
			{Ulasan: "Nyaman dan bersih", Rating: 5},
			{Ulasan: "WiFi cepat", Rating: 4},
		},
		JmlFeedback: 2,
	}
	daftar.Data[1] = WorkingSpace{
		ID:    2,
		Nama:  "KerjaBersama Bandung",
		Lokasi: "Bandung",
		Fasilitas: [10]string{"WiFi", "Locker", "Pantry"},
		JmlFas:    3,
		HargaSewa: 80000,
		Feedback: [MAX_FEEDBACK]Feedback{
			{Ulasan: "Harga terjangkau", Rating: 4},
		},
		JmlFeedback: 1,
	}
	daftar.Data[2] = WorkingSpace{
		ID:    3,
		Nama:  "WorkNest Surabaya",
		Lokasi: "Surabaya",
		Fasilitas: [10]string{"WiFi", "Meeting Room", "Standing Desk", "Gym"},
		JmlFas:    4,
		HargaSewa: 120000,
		Feedback: [MAX_FEEDBACK]Feedback{
			{Ulasan: "Fasilitas lengkap", Rating: 5},
			{Ulasan: "Agak ramai", Rating: 3},
		},
		JmlFeedback: 2,
	}
	daftar.N = 3
}