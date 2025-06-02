package main

import "fmt"

const MAX = 100
const NMAX = 50

type CoWorking struct {
	nama      string
	lokasi    string
	fasilitas string
	harga     int
	rating    float64
	ulasan    string
}

type Pengguna struct {
	username string
	password string
}

var dataPengguna [NMAX]Pengguna
var daftar [MAX]CoWorking
var jumData int = 0
var jumUser int = 0

func main() {
	var dataPengguna [NMAX]Pengguna
	var daftar [MAX]CoWorking
	var jumUser int
	var jumData int

	mainMenu(&dataPengguna, &jumUser, &daftar, &jumData)
}

func mainMenu(dataPengguna *[NMAX]Pengguna, jumUser *int, daftar *[MAX]CoWorking, jumData *int) {
	var run bool = true
	for run == true {
		var isLogin bool = false
		fmt.Println("\n=== MAIN MENU ===")
		fmt.Println("1. LOGIN")
		fmt.Println("2. DAFTAR BARU")
		fmt.Println("3. KELUAR")
		fmt.Print("Pilih menu 1-3: ")

		var pilih int
		fmt.Scan(&pilih)

		if pilih == 1 {
			var user, pass string
			fmt.Println("\n=== LOGIN ===")
			fmt.Print("Username: ")
			fmt.Scan(&user)
			fmt.Print("Password: ")
			fmt.Scan(&pass)
			isLogin = login(*dataPengguna, *jumUser, user, pass)
		} else if pilih == 2 {
			isLogin = daftarBaru(dataPengguna, jumUser)
		} else if pilih == 3 {
			fmt.Println("Terima kasih")
			run = false
		} else {
			fmt.Println("Pilihan tidak valid")
		}
		for isLogin {
			dashBoard(daftar, jumData)
			var pilihan int
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				menuTambah(daftar, jumData)
			case 2:
				menuUbah(daftar, jumData)
			case 3:
				menuHapus(daftar, jumData)
			case 4:
				menuCari(daftar, *jumData)
			case 5:
				menuUrutan(daftar, jumData)
			case 6:
				tampilkanSemuaData(daftar, jumData)
			case 7:
				isLogin = false
			}
		}
	}
}

func daftarBaru(dataPengguna *[NMAX]Pengguna, jumlah *int) bool {
	var i int
	var valid bool = true

	if *jumlah >= MAX {
		fmt.Println("Data pengguna sudah penuh")
	}

	var user Pengguna
	fmt.Println("\n=== DAFTAR PENGGUNA BARU ===")
	fmt.Print("Username: ")
	fmt.Scan(&user.username)
	fmt.Print("Password: ")
	fmt.Scan(&user.password)

	for i = 0; i < *jumlah; i++ {
		if dataPengguna[i].username == user.username {
			valid = false
			fmt.Println("Username sudah digunakan.")
		}
	}

	if valid == true {
		dataPengguna[*jumlah] = user
		*jumlah = *jumlah + 1
		fmt.Println("Pendaftaran berhasil, selamat datang")
		return valid
	} else {
		return valid
	}
}

func login(dataPengguna [NMAX]Pengguna, jumlah int, username string, password string) bool {
	var i int
	var berhasil bool = false

	for i = 0; i < jumlah; i++ {
		if dataPengguna[i].username == username && dataPengguna[i].password == password {
			berhasil = true
		}
	}

	if berhasil {
		fmt.Println("Login berhasil")
		return berhasil
	} else {
		fmt.Println("Username atau Password salah")
		return berhasil
	}
}

func dashBoard(daftar *[MAX]CoWorking, jumlah *int) {
	fmt.Println("\n===== DASHBOARD =====")
	fmt.Println("1. Tambah Data Co-Working Space")
	fmt.Println("2. Ubah Data Co-Working Space")
	fmt.Println("3. Hapus Data Co-Working Space")
	fmt.Println("4. Cari Data Co-Working Space")
	fmt.Println("5. Urutkan Data Co-Working Space")
	fmt.Println("6. Tampilkan semua Co-Working Space")
	fmt.Println("7. Kembali ke Main Menu")
	fmt.Print("Pilih menu 1 - 7: ")
}

func menuTambah(daftar *[MAX]CoWorking, jumlah *int) {
	var valid bool = true
	for valid == true {
		fmt.Println("\n=== MENU TAMBAH ===")
		fmt.Println("1. Tambah Data Co-working Space")
		fmt.Println("2. Tambah Ulasan")
		fmt.Println("3. Tambah Rating")
		fmt.Println("4. Kembali ke Dashboard")
		fmt.Print("Pilih menu 1-4: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahCoWorking(daftar, jumlah)
		case 2:
			tambahUlasan(daftar, *jumlah)
		case 3:
			tambahRating(daftar, *jumlah)
		case 4:
			valid = false
		}
	}
}

func tambahCoWorking(daftar *[MAX]CoWorking, jumlah *int) {
	var cws CoWorking

	fmt.Println("\n=== TAMBAH DATA CO-WORKING SPACE ===")
	fmt.Print("Nama: ")
	fmt.Scan(&cws.nama)
	fmt.Print("Lokasi: ")
	fmt.Scan(&cws.lokasi)
	fmt.Print("Fasilitas: ")
	fmt.Scan(&cws.fasilitas)
	fmt.Print("Harga sewa: ")
	fmt.Scan(&cws.harga)

	daftar[*jumlah] = cws
	*jumlah = *jumlah + 1
	fmt.Println("Data berhasil ditambah")
}

func tambahUlasan(daftar *[MAX]CoWorking, jumlah int) {
	var i int
	var nama string
	var idx int = -1

	fmt.Println("\n=== TAMBAH ULASAN ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Print("Masukkan ulasan: ")
		fmt.Scan(&daftar[idx].ulasan)
		fmt.Println("Ulasan berhasil ditambahkan untuk:", daftar[idx].nama)
		fmt.Println("Ulasan:", daftar[idx].ulasan)
	}
}

func tambahRating(daftar *[MAX]CoWorking, jumlah int) {
	var nama string
	var idx, i int
	var rating float64
	var valid bool = false

	idx = -1
	fmt.Println("\n=== TAMBAH RATING ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		for valid == false {
			fmt.Print("Masukkan rating 1 - 5: ")
			fmt.Scan(&rating)
			if rating >= 1.0 && rating <= 5.0 {
				daftar[idx].rating = rating
				fmt.Println("Rating berhasil ditambahkan untuk:", daftar[idx].nama)
				fmt.Println("Rating:", daftar[idx].rating)
				valid = true
			} else {
				fmt.Println("Rating tidak valid")
			}
		}
	}
}

func menuUbah(daftar *[MAX]CoWorking, jumlah *int) {
	var valid bool = true
	for valid == true {
		fmt.Println("\n=== MENU UBAH ===")
		fmt.Println("1. Ubah Data Co-Working Space")
		fmt.Println("2. Ubah Ulasan")
		fmt.Println("3. Ubah Rating")
		fmt.Println("4. Kembali ke Dashboard")
		fmt.Print("Pilih 1-4: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			ubahCoWorking(daftar, *jumlah)
		case 2:
			ubahUlasan(daftar, *jumlah)
		case 3:
			ubahRating(daftar, *jumlah)
		case 4:
			valid = false
		}
	}
}

func ubahCoWorking(daftar *[MAX]CoWorking, jumlah int) {
	var nama string
	var idx, i int
	idx = -1

	fmt.Println("\n=== UBAH DATA CO-WORKING SPACE ===")
	fmt.Print("Masukkan nama Co-Working Space yang ingin diubah: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
	}

	fmt.Println("Data sebelum:")
	fmt.Println("Nama: ", daftar[idx].nama)
	fmt.Println("Lokasi: ", daftar[idx].lokasi)
	fmt.Println("Fasilitas: ", daftar[idx].fasilitas)
	fmt.Println("Harga: ", daftar[idx].harga)

	fmt.Println("\nMasukkan data baru:")
	fmt.Print("Nama baru: ")
	fmt.Scan(&daftar[idx].nama)
	fmt.Print("Lokasi baru: ")
	fmt.Scan(&daftar[idx].lokasi)
	fmt.Print("Fasilitas baru: ")
	fmt.Scan(&daftar[idx].fasilitas)
	fmt.Print("Harga baru: ")
	fmt.Scan(&daftar[idx].harga)
	fmt.Println("Data berhasil diubah")
}

func ubahUlasan(daftar *[MAX]CoWorking, jumlah int) {
	var nama string
	var idx, i int
	idx = -1

	fmt.Println("\n=== UBAH ULASAN ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	}

	fmt.Print("Ulasan baru: ")
	fmt.Scan(&daftar[idx].ulasan)
	fmt.Println("Ulasan berhasil diubah untuk:", daftar[idx].nama)
	fmt.Println("Ulasan:", daftar[idx].ulasan)
}

func ubahRating(daftar *[MAX]CoWorking, jumlah int) {
	var nama string
	var idx, i int
	idx = -1

	fmt.Println("\n=== UBAH RATING ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	}

	var rating float64
	var valid bool = false
	for valid == false {
		fmt.Print("Masukkan rating baru 1 - 5: ")
		fmt.Scan(&rating)

		if rating >= 1.0 && rating <= 5.0 {
			valid = true
			daftar[idx].rating = rating
			fmt.Println("Rating baru berhasil ditambahkan untuk:", daftar[idx].nama)
			fmt.Println("Rating: ", daftar[idx].rating)
		} else {
			fmt.Println("Rating tidak valid")
		}
	}
}

func menuHapus(daftar *[MAX]CoWorking, jumlah *int) {
	var valid bool = true
	for valid == true {
		fmt.Println("\n=== MENU HAPUS ===")
		fmt.Println("1. Hapus Data Co-working Space")
		fmt.Println("2. Hapus Ulasan")
		fmt.Println("3. Hapus Rating")
		fmt.Println("4. Kembali ke Dashboard")
		fmt.Print("Pilih 1-4: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			hapusCoWorking(daftar, jumlah)
		case 2:
			hapusUlasan(daftar, jumlah)
		case 3:
			hapusRating(daftar, jumlah)
		case 4:
			valid = false
		}
	}
}

func hapusCoWorking(daftar *[MAX]CoWorking, jumlah *int) {
	var nama string
	var idx, i int
	idx = -1

	fmt.Println("\n=== HAPUS DATA CO-WORKING SPACE ===")
	fmt.Print("Masukkan nama Co-Working Space yang ingin dihapus: ")
	fmt.Scan(&nama)

	for i = 0; i < *jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
	}

	for i := idx; i < *jumlah-1; i++ {
		daftar[i] = daftar[i+1]
	}

	*jumlah = *jumlah - 1
	fmt.Println("Data berhasil dihapus.")
}

func hapusUlasan(daftar *[MAX]CoWorking, jumlah *int) {
	var i int
	var nama string
	var found bool = false

	fmt.Println("\n=== HAPUS ULASAN ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < *jumlah; i++ {
		if daftar[i].nama == nama {
			daftar[i].ulasan = ""
			found = true
		}
	}

	if found {
		fmt.Println("Ulasan berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func hapusRating(daftar *[MAX]CoWorking, jumlah *int) {
	var i int
	var nama string
	var found bool = false

	fmt.Println("\n=== HAPUS RATING ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < *jumlah; i++ {
		if daftar[i].nama == nama {
			daftar[i].rating = 0
			found = true
		}
	}

	if found {
		fmt.Println("Rating berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func menuCari(daftar *[MAX]CoWorking, jumlah int) {
	var valid bool = true
	var pilih int
	var nama, lokasi, fasilitas string

	for valid == true {
		fmt.Println("\n=== MENU CARI ===")
		fmt.Println("1. Cari berdasarkan Nama")
		fmt.Println("2. Cari berdasarkan Lokasi")
		fmt.Println("3. Cari berdasarkan Fasilitas")
		fmt.Println("4. Kembali ke Dashboard")
		fmt.Print("Pilih 1-4: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Masukkan nama Co-working Space yang dicari: ")
			fmt.Scan(&nama)
			if cariBerdasarkanNama(daftar, jumlah, nama) == false {
				fmt.Println("Data tidak ditemukan.")
			}
		case 2:
			fmt.Print("Masukkan lokasi Co-working Space yang dicari: ")
			fmt.Scan(&lokasi)
			if cariBerdasarkanLokasi(daftar, jumlah, lokasi) == false {
				fmt.Println("Data tidak ditemukan.")
			}
		case 3:
			fmt.Print("Masukkan fasilitas Co-working Space yang dicari: ")
			fmt.Scan(&fasilitas)
			if cariBerdasarkanFasilitas(daftar, jumlah, fasilitas) == false {
				fmt.Println("Data tidak ditemukan")
			}
		case 4:
			valid = false
		}
	}
}

func cariBerdasarkanNama(daftar *[MAX]CoWorking, jumlah int, nama string) bool {
	var i int
	var found bool = false
	fmt.Println("\n=== PENCARIAN BERDASARKAN NAMA ===")
	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			printData(daftar[i])
			fmt.Println("-----------------------------")
			found = true
		}
	}
	return found
}

func cariBerdasarkanLokasi(daftar *[MAX]CoWorking, jumlah int, lokasi string) bool {
	var i int
	var found bool = false
	fmt.Println("\n=== PENCARIAN BERDASARKAN LOKASI ===")
	for i = 0; i < jumlah; i++ {
		if daftar[i].lokasi == lokasi {
			printData(daftar[i])
			fmt.Println("-----------------------------")
			found = true
		}
	}
	return found
}

func cariBerdasarkanFasilitas(daftar *[MAX]CoWorking, jumlah int, fasilitas string) bool {
	var i int
	var found bool = false
	fmt.Println("\n=== PENCARIAN BERDASARKAN FASILITAS ===")
	for i = 0; i < jumlah; i++ {
		if daftar[i].fasilitas == fasilitas {
			printData(daftar[i])
			fmt.Println("-----------------------------")
			found = true
		}
	}
	return found
}

func printData(A CoWorking) {
	fmt.Println("Nama      :", A.nama)
	fmt.Println("Lokasi    :", A.lokasi)
	fmt.Println("Fasilitas :", A.fasilitas)
	fmt.Println("Harga Sewa:", A.harga)
	fmt.Println("Ulasan    :", A.ulasan)
	fmt.Println("Rating    :", A.rating)
}

func menuUrutan(daftar *[MAX]CoWorking, jumlah *int) {
	var valid bool = true
	var pilih int
	for valid == true {
		fmt.Println("\n=== MENU URUTAN ===")
		fmt.Println("1. Urutkan berdasarkan harga termurah")
		fmt.Println("2. Urutkan berdasarkan harga termahal")
		fmt.Println("3. Urutkan berdasarkan rating terendah")
		fmt.Println("4. Urutkan berdasarkan rating tertinggi")
		fmt.Println("5. Kembali ke DashBoard")
		fmt.Print("Pilih 1-5: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			urutanBerdasarkanHargaTermurah(daftar, jumlah)
		case 2:
			urutanBerdasarkanHargaTermahal(daftar, jumlah)
		case 3:
			urutanBerdasarkanRatingTerendah(daftar, jumlah)
		case 4:
			urutanBerdasarkanRatingTertinggi(daftar, jumlah)
		case 5:
			valid = false
		}
	}
}

func urutanBerdasarkanHargaTermurah(daftar *[MAX]CoWorking, jumlah *int) {
	var i, j, idx int
	var temp CoWorking

	for i = 0; i < *jumlah-1; i++ {
		idx = i
		for j = i + 1; j < *jumlah; j++ {
			if daftar[j].harga < daftar[idx].harga {
				idx = j
			}
		}
		temp = daftar[i]
		daftar[i] = daftar[idx]
		daftar[idx] = temp
	}

	fmt.Println("\nData diurutkan berdasarkan Harga termurah - termahal:")
	for i = 0; i < *jumlah; i++ {
		printData(daftar[i])
		fmt.Println("-----------------------------")
	}
}

func urutanBerdasarkanHargaTermahal(daftar *[MAX]CoWorking, jumlah *int) {
	var i, j, idx int
	var temp CoWorking

	for i = 0; i < *jumlah-1; i++ {
		idx = i
		for j = i + 1; j < *jumlah; j++ {
			if daftar[j].harga > daftar[idx].harga {
				idx = j
			}
		}
		temp = daftar[i]
		daftar[i] = daftar[idx]
		daftar[idx] = temp
	}

	fmt.Println("\nData diurutkan berdasarkan Harga termahal - termurah:")
	for i = 0; i < *jumlah; i++ {
		printData(daftar[i])
		fmt.Println("-----------------------------")
	}
}

func urutanBerdasarkanRatingTerendah(daftar *[MAX]CoWorking, jumlah *int) {
	var i, j int
	var temp CoWorking

	for i = 1; i < *jumlah; i++ {
		temp = daftar[i]
		j = i - 1
		for j >= 0 && daftar[j].rating > temp.rating {
			daftar[j+1] = daftar[j]
			j = j - 1
		}
		daftar[j+1] = temp
	}

	fmt.Println("\nData diurutkan berdasarkan Rating terendah - tertinggi:")
	for i = 0; i < *jumlah; i++ {
		printData(daftar[i])
		fmt.Println("-----------------------------")
	}
}

func urutanBerdasarkanRatingTertinggi(daftar *[MAX]CoWorking, jumlah *int) {
	var i, j int
	var temp CoWorking

	for i = 1; i < *jumlah; i++ {
		temp = daftar[i]
		j = i - 1
		for j >= 0 && daftar[j].rating < temp.rating {
			daftar[j+1] = daftar[j]
			j = j - 1
		}
		daftar[j+1] = temp
	}

	fmt.Println("\nData diurutkan berdasarkan Rating tertinggi - terendah:")
	for i = 0; i < *jumlah; i++ {
		printData(daftar[i])
		fmt.Println("-----------------------------")
	}
}

func tampilkanSemuaData(daftar *[MAX]CoWorking, jumlah *int) {
	var i int
	fmt.Println("\n=== DAFTAR SEMUA CO-WORKING SPACE ===")
	for i = 0; i < *jumlah; i++ {
		fmt.Println("Nama      :", daftar[i].nama)
		fmt.Println("Lokasi    :", daftar[i].lokasi)
		fmt.Println("Fasilitas :", daftar[i].fasilitas)
		fmt.Println("Harga Sewa:", daftar[i].harga)
		fmt.Println("Ulasan    :", daftar[i].ulasan)
		fmt.Println("Rating    :", daftar[i].rating)
		fmt.Println("-----------------------------")
	}
}
