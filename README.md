# Dokumentasi Tugas Besar
## Aplikasi Manajemen dan Review Co-Working Space

---

## 1. Deskripsi Aplikasi

Aplikasi **Co-Working Space Management System** adalah program berbasis CLI (Command-Line Interface) yang dibangun menggunakan bahasa pemrograman Go. Aplikasi ini dirancang untuk membantu dua kelompok pengguna utama:

- **Pemilik co-working space** — untuk mengelola data tempat kerja mereka, mencakup informasi nama, lokasi, fasilitas, dan harga sewa.
- **Pekerja remote / freelancer** — untuk mencari, membandingkan, dan memberikan ulasan terhadap co-working space yang pernah mereka kunjungi.

Fitur utama aplikasi meliputi:

- **CRUD Co-Working Space**: tambah, ubah, dan hapus data co-working space beserta fasilitasnya.
- **CRUD Feedback & Rating**: tambah, ubah, dan hapus ulasan serta rating (skala 1–5) untuk setiap tempat.
- **Pencarian**: Sequential Search (berdasarkan nama, lokasi, fasilitas) dan Binary Search (berdasarkan nama atau lokasi secara exact match).
- **Pengurutan**: Selection Sort dan Insertion Sort berdasarkan harga sewa (ascending) atau rating rata-rata (descending).
- **Tampil data**: menampilkan semua co-working space beserta detail fasilitas, harga, dan statistik rating.

Struktur data utama yang digunakan adalah array statis (`[MAX_SPACES]WorkingSpace`) yang dibungkus dalam struct `DaftarSpace`, tanpa penggunaan fitur-fitur Go tingkat lanjut seperti goroutine, generics, atau package eksternal.

---

## 2. Penjelasan Alur Program

### 2.1 Inisialisasi

Saat program dijalankan, fungsi `main()` pada `main.go` memanggil `inisialisasiData()` untuk mengisi array dengan 3 data dummy awal sebagai demonstrasi. Setelah itu, program memasuki loop utama yang menampilkan menu utama.

### 2.2 Menu Utama

Program menampilkan 5 pilihan menu utama yang masing-masing memanggil fungsi menu yang sesuai:

```
1. Kelola Co-Working Space      → menuCRUDSpace()
2. Kelola Feedback & Rating     → menuCRUDFeedback()
3. Pencarian                    → menuPencarian()
4. Pengurutan / Tampil Daftar   → menuPengurutan()
5. Tampil Semua Co-Working Space → tampilSemuaWorkingSpace()
0. Keluar
```

Setiap menu membentuk sub-loop tersendiri yang akan kembali ke menu utama saat pengguna memilih opsi `0`.

### 2.3 Alur CRUD Co-Working Space

1. Pengguna memilih submenu (tambah / ubah / hapus / lihat).
2. Untuk **tambah**: program membaca input nama, lokasi, jumlah fasilitas, nama tiap fasilitas, dan harga sewa. ID di-generate otomatis oleh `nextID()`.
3. Untuk **ubah**: program mencari data berdasarkan ID menggunakan `cariIndexByID()`, lalu memperbarui field yang diisi pengguna (field kosong tidak diubah).
4. Untuk **hapus**: program mencari data berdasarkan ID, lalu menggeser elemen array ke kiri untuk menutup celah.

### 2.4 Alur CRUD Feedback

1. Pengguna memasukkan ID co-working space yang dituju.
2. Program mencari index-nya menggunakan `cariIndexByID()`.
3. Untuk **tambah**: program membaca ulasan (string) dan rating (integer 1–5), lalu menyimpannya ke dalam array `Feedback` di dalam struct `WorkingSpace`.
4. Untuk **ubah/hapus**: program menampilkan daftar feedback yang ada, pengguna memilih nomor feedback, dan program memperbarui atau menggeser array feedback.

### 2.5 Alur Pencarian

- **Sequential Search**: program iterasi seluruh array dan mengumpulkan index yang cocok menggunakan `strings.Contains()` (case-insensitive). Mendukung partial match.
- **Binary Search**: program menyalin array ke variabel `salinan`, mengurutkan salinan tersebut berdasarkan field yang relevan (nama atau lokasi), lalu menjalankan binary search pada salinan. Data asli tidak dimodifikasi.

### 2.6 Alur Pengurutan

Saat pengguna memilih opsi pengurutan, program menyalin array ke variabel `salinan`, menjalankan algoritma sort yang dipilih (Selection/Insertion Sort berdasarkan harga atau rating) pada salinan tersebut, lalu menampilkan hasilnya. Data asli tetap tidak berubah.

---

## 3. Penjelasan Subprogram dan Interaksinya

Program diorganisasi secara modular dalam beberapa file Go, masing-masing dengan tanggung jawab yang spesifik.

### 3.1 `models.go` — Definisi Tipe Data

Mendefinisikan seluruh struct dan konstanta yang digunakan di seluruh program. Tidak mengandung logika apapun.

| Tipe | Keterangan |
|------|------------|
| `Feedback` | Menyimpan satu ulasan: `Ulasan` (string) dan `Rating` (uint8, 1–5) |
| `WorkingSpace` | Data satu co-working space: ID, nama, lokasi, array fasilitas, harga sewa, dan array feedback |
| `DaftarSpace` | Wrapper array statis `[MAX_SPACES]WorkingSpace` beserta jumlah data aktif (`N`) |

Semua file lain bergantung pada tipe-tipe ini. `models.go` adalah fondasi dari seluruh program.

---

### 3.2 `data.go` — Inisialisasi Data

| Subprogram | Jenis | Keterangan |
|------------|-------|------------|
| `inisialisasiData(daftar *DaftarSpace)` | Prosedur | Mengisi `daftar` dengan 3 data dummy awal |

Dipanggil sekali oleh `main()` di awal program. Tidak dipanggil dari tempat lain.

---

### 3.3 `crud.go` — Input, Validasi, dan Operasi Data

File ini memuat dua kelompok subprogram: **helper input** dan **operasi CRUD**.

**Helper Input:**

| Subprogram | Jenis | Keterangan |
|------------|-------|------------|
| `bacaString(prompt)` | Fungsi | Membaca satu baris input string |
| `bacaInt(prompt)` | Fungsi | Membaca integer, loop ulang jika bukan angka |
| `bacaIntDenganValidasi(prompt, min, max, headerFunc)` | Fungsi | Membaca integer dalam rentang min–max; memanggil `headerFunc()` untuk mencetak ulang konteks jika input gagal |
| `bacaID(prompt)` | Fungsi | Membaca ID valid (uint8, 1–255) |
| `nextID(daftar)` | Fungsi | Menghitung ID baru dengan mencari nilai maksimum ID yang ada lalu menambah 1 |
| `cariIndexByID(daftar, id)` | Fungsi | Mencari index di array berdasarkan ID; mengembalikan -1 jika tidak ditemukan |

**Operasi CRUD Co-Working Space:**

| Subprogram | Jenis | Keterangan |
|------------|-------|------------|
| `TambahWorkingSpace(daftar)` | Prosedur | Membaca input, membentuk struct `WorkingSpace`, menyimpan ke array |
| `UbahWorkingSpace(daftar)` | Prosedur | Mencari data by ID, memperbarui field yang tidak kosong |
| `HapusWorkingSpace(daftar)` | Prosedur | Mencari data by ID, menggeser array ke kiri, mengurangi `N` |

**Operasi CRUD Feedback:**

| Subprogram | Jenis | Keterangan |
|------------|-------|------------|
| `TambahFeedback(daftar)` | Prosedur | Mencari space by ID, menambahkan `Feedback` baru ke array feedback-nya |
| `UbahFeedback(daftar)` | Prosedur | Menampilkan daftar feedback, memperbarui feedback yang dipilih |
| `HapusFeedback(daftar)` | Prosedur | Menghapus feedback pada index tertentu dengan pergeseran array |

Semua subprogram CRUD bergantung pada helper input (`bacaInt`, `bacaString`, `bacaID`, `bacaIntDenganValidasi`) dan `cariIndexByID()`.

---

### 3.4 `sort.go` — Algoritma Pengurutan

| Subprogram | Jenis | Keterangan |
|------------|-------|------------|
| `hitungRataRating(ws WorkingSpace)` | Fungsi | Menghitung rata-rata rating dari semua feedback; mengembalikan 0 jika belum ada feedback |
| `SelectionSortByHarga(arr)` | Prosedur | Selection Sort ascending berdasarkan `HargaSewa` |
| `InsertionSortByHarga(arr)` | Prosedur | Insertion Sort ascending berdasarkan `HargaSewa` |
| `SelectionSortByRating(arr)` | Prosedur | Selection Sort descending berdasarkan rata-rata rating (memanggil `hitungRataRating`) |
| `InsertionSortByRating(arr)` | Prosedur | Insertion Sort descending berdasarkan rata-rata rating (memanggil `hitungRataRating`) |

Semua fungsi sort beroperasi pada pointer `*DaftarSpace`. Dipanggil oleh `menuPengurutan()` di `menu.go`, selalu pada salinan data (`salinan := *daftar`) sehingga data asli tidak termodifikasi.

---

### 3.5 `search.go` — Algoritma Pencarian

| Subprogram | Jenis | Keterangan |
|------------|-------|------------|
| `SeqSearchByNama(daftar, keyword)` | Fungsi | Sequential search partial match pada field `Nama` |
| `SeqSearchByLokasi(daftar, keyword)` | Fungsi | Sequential search partial match pada field `Lokasi` |
| `SeqSearchByFasilitas(daftar, fasilitas)` | Fungsi | Sequential search partial match pada array `Fasilitas` tiap space |
| `BinarySearchByNama(daftar, target)` | Fungsi | Binary search exact match pada `Nama`; data **harus** sudah terurut ascending |
| `BinarySearchByLokasi(daftar, target)` | Fungsi | Binary search exact match pada `Lokasi`; data **harus** sudah terurut ascending |
| `SortByNamaForBinary(daftar)` | Prosedur | Insertion sort berdasarkan `Nama` (ascending) sebagai prasyarat Binary Search |
| `SortByLokasiForBinary(daftar)` | Prosedur | Insertion sort berdasarkan `Lokasi` (ascending) sebagai prasyarat Binary Search |

Sequential search mengembalikan `([]int, int)` — slice index hasil dan jumlahnya — untuk mendukung pencarian multi-hasil. Binary search mengembalikan satu `int` index (atau -1). Sebelum binary search dipanggil, `menu.go` selalu memanggil fungsi sort yang sesuai pada salinan data.

---

### 3.6 `display.go` — Tampilan dan Output

| Subprogram | Jenis | Keterangan |
|------------|-------|------------|
| `clearScreen()` | Prosedur | Membersihkan layar terminal (cross-platform: Windows/Unix) |
| `tampilWorkingSpace(ws WorkingSpace)` | Prosedur | Mencetak detail satu co-working space termasuk rata-rata rating (memanggil `hitungRataRating`) |
| `tampilSemuaWorkingSpace(daftar)` | Prosedur | Iterasi seluruh array dan memanggil `tampilWorkingSpace` tiap elemen |
| `tampilHasilSearch(daftar, indices, jumlah)` | Prosedur | Menampilkan hasil pencarian berdasarkan slice index yang dikembalikan oleh fungsi search |
| `tampilFeedback(daftar)` | Prosedur | Membaca ID, mencari space-nya, lalu menampilkan semua feedback dengan visualisasi bintang (★☆) |

`tampilWorkingSpace` bergantung pada `hitungRataRating` dari `sort.go`. `tampilFeedback` menggunakan `bacaInt` dan `cariIndexByID` dari `crud.go`.

---

### 3.7 `menu.go` — Handler Menu (Orkestrator)

File ini berfungsi sebagai lapisan penghubung antara antarmuka pengguna dan logika bisnis.

| Subprogram | Jenis | Keterangan |
|------------|-------|------------|
| `menuCRUDSpace(daftar)` | Prosedur | Sub-loop menu untuk operasi tambah/ubah/hapus/lihat co-working space |
| `menuCRUDFeedback(daftar)` | Prosedur | Sub-loop menu untuk operasi tambah/ubah/hapus/lihat feedback |
| `menuPencarian(daftar)` | Prosedur | Sub-loop menu pencarian; memanggil fungsi search dan display sesuai pilihan |
| `menuPengurutan(daftar)` | Prosedur | Sub-loop menu pengurutan; menyalin data, memanggil fungsi sort, lalu menampilkan hasil |

Semua fungsi menu dipanggil dari `main()` di `main.go`. Mereka tidak mengandung logika bisnis sendiri, melainkan mendelegasikan ke subprogram di `crud.go`, `sort.go`, `search.go`, dan `display.go`.

---

## 4. Diagram Interaksi Antar Subprogram

```
main.go
  └── main()
        ├── inisialisasiData()                  [data.go]
        ├── menuCRUDSpace()                     [menu.go]
        │     ├── TambahWorkingSpace()           [crud.go]
        │     │     ├── bacaString/bacaInt/bacaIntDenganValidasi()
        │     │     └── nextID()
        │     ├── UbahWorkingSpace()             [crud.go]
        │     │     └── cariIndexByID()
        │     ├── HapusWorkingSpace()            [crud.go]
        │     │     └── cariIndexByID()
        │     └── tampilSemuaWorkingSpace()      [display.go]
        │           └── tampilWorkingSpace()
        │                 └── hitungRataRating() [sort.go]
        │
        ├── menuCRUDFeedback()                  [menu.go]
        │     ├── TambahFeedback()              [crud.go]
        │     ├── UbahFeedback()                [crud.go]
        │     ├── HapusFeedback()               [crud.go]
        │     └── tampilFeedback()              [display.go]
        │
        ├── menuPencarian()                     [menu.go]
        │     ├── SeqSearchByNama/Lokasi/Fasilitas()  [search.go]
        │     ├── SortByNamaForBinary / SortByLokasiForBinary()
        │     ├── BinarySearchByNama / BinarySearchByLokasi()
        │     └── tampilHasilSearch()           [display.go]
        │
        ├── menuPengurutan()                    [menu.go]
        │     ├── SelectionSort/InsertionSortByHarga()   [sort.go]
        │     ├── SelectionSort/InsertionSortByRating()  [sort.go]
        │     │     └── hitungRataRating()
        │     └── tampilSemuaWorkingSpace()     [display.go]
        │
        └── tampilSemuaWorkingSpace()           [display.go]
```

---

## 5. Implementasi Ketentuan Tubes

| No | Ketentuan | Implementasi |
|----|-----------|--------------|
| 1 | Algoritma terstruktur | Semua logika menggunakan sequence, selection (switch/if), dan iteration (for) |
| 2 | Konsep modular | Program dibagi ke 8 file dengan fungsi dan prosedur yang terdefinisi jelas |
| 3 | Array dan struct/record | `DaftarSpace` (array statis), `WorkingSpace`, `Feedback` (struct) |
| 4 | Data dibaca dari piranti masukan | Semua input dibaca via `bacaString`, `bacaInt`, `bacaID` dari stdin |
| 5 | Sorting dan searching | Selection Sort, Insertion Sort, Sequential Search, Binary Search |
| 6 | Bahasa Go, fitur dasar | Hanya menggunakan array, struct, fungsi, pointer, dan package standar |
| 7 | Dokumentasi subprogram | Dokumen ini |
