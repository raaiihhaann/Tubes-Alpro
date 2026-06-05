package main

const MAX_SPACES = 100
const MAX_FEEDBACK = 50

type Feedback struct {
	Ulasan string
	Rating uint8
}

type WorkingSpace struct {
	ID        uint8
	Nama      string
	Lokasi    string
	Fasilitas [10]string
	JmlFas    int
	HargaSewa int
	Feedback  [MAX_FEEDBACK]Feedback
	JmlFeedback int
}

type DaftarSpace struct {
	Data [MAX_SPACES]WorkingSpace
	N    int
}

type DaftarFasilitas struct {
	Items [10]string
	N     int
}