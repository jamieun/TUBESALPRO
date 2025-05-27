package main

import "fmt"

const MAX = 100

type CoWorking struct {
	nama      string
	lokasi    string
	fasilitas string
	harga     int
	rating    float64
}

type Pengguna struct {
	username string
	password string
}

var dataPengguna [MAX]Pengguna
var daftar [MAX]CoWorking
var jumlah int = 0

func main() {
	mainMenu()
}

func mainMenu() {
	var run bool = true
	for run == true {
		fmt.Println("\n=== MAINMENU ===")
		fmt.Println("1. LOGIN")
		fmt.Println("2. DAFTAR BARU")
		fmt.Println("3. KELUAR")
		fmt.Println("Pilih menu 1-3: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			login()
		case 2:
			daftarBaru()
		case 3:
			run = false
		}
	}
}

func daftarBaru() {
	if jumlah >= MAX {
		fmt.Println("Data sudah penuh.")
		return
	}

	var user Pengguna
	var i int
	fmt.Println("\n=== DAFTAR PENGGUNA BARU ===")
	fmt.Print("Username: ")
	fmt.Scan(&user.username)
	fmt.Print("Password: ")
	fmt.Scan(&user.password)

	for i = 0; i < jumlah; i++ {
		if dataPengguna[i].username == user.username {
			fmt.Println("Username sudah digunakan")
			return
		}
	}

	dataPengguna[jumlah] = user
	jumlah++
	fmt.Println("Pendaftaran berhasil, Silahkan login")
}

func login() {
	var user, pass string
	var berhasil bool = false
	var i int

	fmt.Println("\n=== LOGIN ===")
	fmt.Print("Username: ")
	fmt.Scan(&user)
	fmt.Print("Password: ")
	fmt.Scan(&pass)

	for i = 0; i < jumlah; i++ {
		if dataPengguna[i].username == user && dataPengguna[i].password == pass {
			berhasil = true
		}
	}

	if berhasil {
		fmt.Println("Login berhasil")
		dashBoard()
	} else {
		fmt.Println("Username atau Password salah")
	}
}

func dashBoard() {
	for {
		fmt.Println("\n=== DASHBOARD ===")
		fmt.Println("1. Tambah Co-working Space")
		fmt.Println("2. Ubah Co-working Space")
		fmt.Println("3. Hapus Co-working Space")
		fmt.Println("4. Cari Co-Working Space")
		fmt.Println("5. Urutkan Co-Working Space")
		fmt.Println("6.Kembali ke Main Menu")
		fmt.Print("Pilih 1-6: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			menuTambah()
		case 2:
			menuUbah()
		case 3:
			menuHapus()
		case 4:
			menuCari()
		case 5:
			menuUrutan()
		case 6:
			mainMenu()
		}
	}
}

func menuTambah() {
	for {
		fmt.Println("\n=== MENU TAMBAH ===")
		fmt.Println("1. Tambah Data Co-working Space")
		fmt.Println("2. Tambah Ulasan")
		fmt.Println("3. Tambah Rating")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Println("pilih menu 1-4:")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahDataCoWorking()
		case 2:
			tambahUlasan()
		case 3:
			tambahRating()
		case 4:
			dashBoard()
		}
	}
}

func tambahDataCoWorking() {
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

	daftar[jumlah] = cws
	jumlah++
	fmt.Println("Data berhasil ditambah")
	menuTambah()
}

func tambahUlasan() {
	var nama string
	var idx, i int
	var found bool = false

	fmt.Println("\n=== TAMBAH ULASAN ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
			found = true
		}
	}

	if found == false {
		fmt.Println("Co-Working Space tidak ditemukan")
		return
	}

	var ulasan string
	fmt.Print("Masukkan ulasan: ")
	fmt.Scan(&ulasan)

	fmt.Println("Ulasan berhasil ditambahkan untuk:", daftar[idx].nama)
	fmt.Println("Ulasan:", ulasan)
}

func tambahRating() {
	var nama string
	var idx, i int
	var found bool = false

	fmt.Println("\n=== TAMBAH RATING ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
			found = true
		}
	}

	if found == false {
		fmt.Println("Co-Working Space tidak ditemukan.")
		return
	}

	var rating float64
	var valid bool = false
	for valid == false {
		fmt.Print("Masukkan rating 1 - 5: ")
		fmt.Scan(&rating)

		if rating >= 1.0 && rating <= 5.0 {
			valid = true
			daftar[idx].rating = rating
			fmt.Println("Rating berhasil ditambahkan untuk:", daftar[idx].nama)
			fmt.Println("Rating: ", daftar[idx].rating)
		}
	}
}

func menuUbah() {
	fmt.Println("\n=== MENU UBAH ===")
	fmt.Println("1. Ubah Data Co-Working Space")
	fmt.Println("2. Ubah Ulasan")
	fmt.Println("3. Ubah Rating")
	fmt.Println("4. Kembali ke Dashboard")
	fmt.Println("pilih 1-4:")

	var pilih int
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		ubahCoWorking()
	case 2:
		ubahUlasan()
	case 3:
		ubahRating()
	case 4:
		dashBoard()
	}
}

func ubahCoWorking() {
	var nama string
	var idx, i int
	var found bool = false

	fmt.Println("\n=== UBAH DATA CO-WORKING SPACE ===")
	fmt.Print("Masukkan nama Co-Working Space yang ingin diubah: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
			found = true
		}
	}

	if found == false {
		fmt.Println("Co-Working Space tidak ditemukan.")
		return
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
	fmt.Println("Data berhasil diubah.")
	menuUbah()
}

func ubahUlasan() {
	var nama string
	var idx, i int
	var found bool = false

	fmt.Println("\n=== UBAH ULASAN ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
			found = true
		}
	}

	if found == false {
		fmt.Println("Co-working space tidak ditemukan.")
		return
	}
	var ulasan string
	fmt.Print("Ulasan baru: ")
	fmt.Scan(&ulasan)
	fmt.Println("Ulasan berhasil diubah untuk:", daftar[idx].nama)
}

func ubahRating() {
	var nama string
	var idx, i int
	var found bool = false

	fmt.Println("\n=== UBAH RATING ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
			found = true
		}
	}

	if found == false {
		fmt.Println("Co-Working Space tidak ditemukan.")
		return
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
		}
	}
}

func menuHapus() {
	fmt.Println("\n=== MENU HAPUS ===")
	fmt.Println("1. Hapus Data Co-working Space")
	fmt.Println("2. Hapus Ulasan")
	fmt.Println("3. Hapus Rating")
	fmt.Println("4. Kembali ke Dashboard")
	fmt.Println("pilih 1-4:")

	var pilih int
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		hapusData()
	case 2:
		hapusUlasan()
	case 3:
		hapusRating()
	case 4:
		dashBoard()
	}
}

func hapusData() {
	var nama string
	var idx, i int
	var found bool = false

	fmt.Println("\n=== HAPUS DATA CO-WORKING SPACE ===")
	fmt.Print("Masukkan nama Co-Working Space yang ingin dihapus: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
			found = true
		}
	}

	if !found {
		fmt.Println("Co-Working Space tidak ditemukan.")
		return
	}

	for i = idx; i < jumlah-1; i++ {
		daftar[i] = daftar[i+1]
	}

	jumlah--
	fmt.Println("Data berhasil dihapus")
}

func hapusUlasan() {
	var nama string
	var found bool = false

	fmt.Println("\n=== HAPUS ULASAN ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i := 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			found = true
		}
	}

	if found == false {
		fmt.Println("Co-Working Space tidak ditemukan.")
		return
	}
	fmt.Println("Ulasan berhasil dihapus.")
}

func hapusRating() {
	var nama string
	var idx, i int
	var found bool = false

	fmt.Println("\n=== HAPUS RATING ===")
	fmt.Print("Masukkan nama Co-Working Space: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			idx = i
			found = true
		}
	}

	if found == false {
		fmt.Println("Co-Working Space tidak ditemukan.")
		return
	}

	daftar[idx].rating = 0
	fmt.Println("Rating berhasil dihapus.")
}

func menuCari() {
	fmt.Println("\n=== MENU CARI ===")
	fmt.Println("1. Cari berdasarkan Nama")
	fmt.Println("2. Cari berdasarkan Lokasi")
	fmt.Println("3. Kembali ke Dashboard")
	fmt.Print("Pilih 1-3: ")

	var pilih int
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		cariBerdasarkanNama()
	case 2:
		cariBerdasarkanLokasi()
	case 3:
		dashBoard()
	}
}

func sequentialSearchNama(X string) int {
	var i int
	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == X {
			return i
		}
	}
	return -1
}

func sequentialSearchLokasi(X string) int {
	var i int
	for i = 0; i < jumlah; i++ {
		if daftar[i].lokasi == X {
			return i
		}
	}
	return -1
}

func binarySearchNama(X string) int {
	var med, kr, kn int
	var found bool = false
	kr = 0
	kn = jumlah - 1

	for kr <= kn && found == false {
		med = (kr + kn) / 2
		if daftar[med].nama < X {
			kr = med + 1
		} else if daftar[med].nama > X {
			kn = med - 1
		} else {
			found = true
		}
	}

	if found {
		return med
	} else {
		return -1
	}
}

func binarySearchLokasi(X string) int {
	var med, kr, kn int
	var found bool = false
	kr = 0
	kn = jumlah - 1

	for kr <= kn && found == false {
		med = (kr + kn) / 2
		if daftar[med].lokasi < X {
			kr = med + 1
		} else if daftar[med].lokasi > X {
			kn = med - 1
		} else {
			found = true
		}
	}

	if found {
		return med
	} else {
		return -1
	}
}

func cariBerdasarkanNama() {
	var nama string
	var idxSeq, idxBin int
	fmt.Print("Masukkan nama Co-working Space yang dicari: ")
	fmt.Scan(&nama)

	idxSeq = sequentialSearchNama(nama)
	if idxSeq == -1 {
		fmt.Println("Data tidak ditemukan (Sequential Search)")
	} else {
		fmt.Println("Data ditemukan (Sequential Search):")
		printData(daftar[idxSeq])
	}

	selectionSortNama()
	idxBin = binarySearchNama(nama)
	if idxBin == -1 {
		fmt.Println("Data tidak ditemukan (Binary Search)")
	} else {
		fmt.Println("Data ditemukan (Binary Search):")
		printData(daftar[idxBin])
	}
	menuCari()
}

func cariBerdasarkanLokasi() {
	var lokasi string
	var idxSeq, idxBin int
	fmt.Print("Masukkan lokasi Co-working Space yang dicari: ")
	fmt.Scan(&lokasi)

	idxSeq = sequentialSearchLokasi(lokasi)
	if idxSeq == -1 {
		fmt.Println("Data tidak ditemukan (Sequential Search)")
	} else {
		fmt.Println("Data ditemukan (Sequential Search):")
		printData(daftar[idxSeq])
	}

	selectionSortLokasi()
	idxBin = binarySearchLokasi(lokasi)
	if idxBin == -1 {
		fmt.Println("Data tidak ditemukan (Binary Search)")
	} else {
		fmt.Println("Data ditemukan (Binary Search):")
		printData(daftar[idxBin])
	}
	menuCari()
}

func printData(data CoWorking) {
	fmt.Println("Nama     :", data.nama)
	fmt.Println("Lokasi   :", data.lokasi)
	fmt.Println("Fasilitas:", data.fasilitas)
	fmt.Println("Harga    :", data.harga)
}

func menuUrutan() {
	fmt.Println("\n=== MENU URUTAN ===")
	fmt.Println("1. Urutkan berdasarkan Harga")
	fmt.Println("2. Urutkan berdasarkan Rating")
	fmt.Println("3. Kembali ke Dashboard")
	fmt.Print("Pilih 1-3: ")

	var pilih int
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		urutanBerdasarkanHarga()
	case 2:
		urutanBerdasarkanRating()
	case 3:
		dashBoard()
	}
}

func selectionSortNama() {
	var idx, i, j int
	var temp CoWorking
	for i = 0; i < jumlah-1; i++ {
		idx = i
		for j = i + 1; j < jumlah; j++ {
			if daftar[j].nama < daftar[idx].nama {
				idx = j
			}
		}
		temp = daftar[i]
		daftar[i] = daftar[idx]
		daftar[idx] = temp
	}
}

func selectionSortLokasi() {
	var idx, i, j int
	var temp CoWorking
	for i = 0; i < jumlah-1; i++ {
		idx = i
		for j = i + 1; j < jumlah; j++ {
			if daftar[j].lokasi < daftar[idx].lokasi {
				idx = j
			}
		}
		temp = daftar[i]
		daftar[i] = daftar[idx]
		daftar[idx] = temp
	}
}

func urutanBerdasarkanHarga() {

}

func urutanBerdasarkanRating() {

}
