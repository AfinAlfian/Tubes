package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

const NMAX int = 1024

type kripto [NMAX]Akripto
type porto [NMAX]Aportofolio
type riwayat [NMAX]Ariwayat

type Aprofil struct {
	username string
	password string
	saldo    int
}

type Akripto struct {
	nama  string
	harga int
}

type Aportofolio struct {
	keuntungan int
	kerugian   int
}

type Ariwayat struct {
	status  string
	nkripto string
	jumlah  int
	harga   int
	tanggal string
}

// func clearScreen() {
// 	fmt.Print("\033[H\033[2J")
// }

func login(A *Aprofil) {
	fmt.Println("-----------------------------------------")
	fmt.Println("|   Selamat datang di aplikasi kripto   |")
	fmt.Println("-----------------------------------------")
	fmt.Print("Masukkan username: ")
	fmt.Scan(&A.username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&A.password)
	//clearScreen()
}

func main() {
	var p Aprofil
	var k kripto
	var b int
	var n int = 3
	k[0].nama = "Bitcoin"
	k[1].nama = "Ethereum"
	k[2].nama = "Litecoin"
	k[0].harga = 1000000000
	k[1].harga = 4555300
	k[2].harga = 87000000

	login(&p)
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("|              Menu Utama               |")
		fmt.Println("-----------------------------------------")
		fmt.Println("username:", p.username)
		fmt.Println("saldo:", p.saldo)
		fmt.Println("-----------------------------------------")
		for i := 0; i < 3; i++ {
			fmt.Println(i+1, k[i].nama, "Harga:", k[i].harga)
		}
		fmt.Println("-----------------------------------------")
		fmt.Println("1. Lihat Daftar Kripto")
		fmt.Println("2. Edit Kripto")
		fmt.Println("2. Lihat Portofolio")
		fmt.Println("3. Lihat Riwayat Transaksi")
		fmt.Println("4. Beli Kripto")
		fmt.Println("5. Jual Kripto")
		fmt.Println("6. Keluar")
		fmt.Println("-----------------------------------------")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&b)

		switch b {
		case 1:
			lihatKripto(k, n)
		case 2:
			menuKripto(&k, &n)
		case 3:
		case 4:
		case 5:
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}

}

func menuKripto(k *kripto, n *int) {
	for {
		var a int
		fmt.Println("-----------------------------------------")
		fmt.Println("|              Menu Kripto              |")
		fmt.Println("-----------------------------------------")
		fmt.Println("1. Lihat Daftar Kripto")
		fmt.Println("2. Tambah Kripto")
		fmt.Println("3. Hapus Kripto")
		fmt.Println("4. Cari Kripto")
		fmt.Println("5. Urutkan Kripto")
		fmt.Println("6. Kembali ke Menu Utama")
		fmt.Println("-----------------------------------------")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&a)
		switch a {
		case 1:
			lihatKripto(*k, *n)
		case 2:
			inputKripto(k, n)
		case 3:
			fmt.Print("Masukkan nama kripto yang ingin dihapus: ")
			var x string
			fmt.Scan(&x)
			hapusKripto(k, n, x)
		case 4:
			fmt.Print("Cari berdasarkan (harga/nama): ")
			var y string
			fmt.Scan(&y)
			if y == "harga" {
				fmt.Print("Masukkan harga kripto yang ingin dicari: ")
				var x int
				fmt.Scan(&x)
				idx := binarySearchInt(*k, *n, x)
				if idx != -1 {
					fmt.Println("Kripto ditemukan di index:", idx)
					fmt.Println("Nama:", k[idx].nama, "Harga:", k[idx].harga)
				} else {
					fmt.Println("Kripto tidak ditemukan")
				}
			} else if y == "nama" {
				fmt.Print("Masukkan nama kripto yang ingin dicari: ")
				var x string
				fmt.Scan(&x)
				idx := binarySearchStr(*k, *n, x)
				if idx != -1 {
					fmt.Println("Kripto ditemukan di index:", idx)
					fmt.Println("Nama:", k[idx].nama, "Harga:", k[idx].harga)
				} else {
					fmt.Println("Kripto tidak ditemukan")
				}
			}

		case 5:
			fmt.Print("Urutkan berdasarkan (harga/nama): ")
			var y string
			fmt.Scan(&y)
			if y == "harga" {
				SelectionSortInt(k, *n)
			} else if y == "nama" {
				SelectionSort(k, *n)
			}
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func lihatKripto(k kripto, n int) {
	fmt.Println("-----------------------------------------")
	fmt.Println("|              Daftar Kripto            |")
	fmt.Println("-----------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Println(i+1, k[i].nama, "Harga:", k[i].harga)
	}
	fmt.Println("-----------------------------------------")
}

func inputKripto(k *kripto, n *int) {
	*n = 3
	fmt.Scan(&k[*n].nama)
	fmt.Scan(&k[*n].harga)
	for k[*n].nama != "-" || k[*n].harga != 0 {
		*n++
		fmt.Print("Masukkan nama kripto: ")
		fmt.Scan(&k[*n].nama)
		fmt.Print("Masukkan harga kripto: ")
		fmt.Scan(&k[*n].harga)
	}
}

func hapusKripto(k *kripto, n *int, x string) {
	var i, j int
	for i = 0; i < *n; i++ {
		if k[i].nama == x {
			for j = i; j < *n-1; j++ {
				k[j] = k[j+1]
			}
			*n--
		}
	}
}

func binarySearchInt(A kripto, n, x int) int {
	var left, right, mid, idx int

	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < A[mid].harga {
			right = mid - 1
		} else if x > A[mid].harga {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func binarySearchStr(A kripto, n int, x string) int {
	var left, right, mid, idx int

	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < A[mid].nama {
			right = mid - 1
		} else if x > A[mid].nama {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func sequentialSearchInt(A kripto, n, x int) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if A[i].harga == x {
			idx = i
		}
	}
	return idx
}

func sequentialSearchStr(A kripto, n int, x string) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if A[i].nama == x {
			idx = i
		}
	}
	return idx
}

func SelectionSort(A *kripto, n int) {
	var i, idx, pass int
	var temp string

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].nama > A[idx].nama {
				idx = i
			}
			i++
		}
		temp = A[pass-1].nama
		A[pass-1].nama = A[idx].nama
		A[idx].nama = temp
		pass++
	}
}

func SelectionSortInt(A *kripto, n int) {
	var i, idx, pass int
	var temp int

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].harga > A[idx].harga {
				idx = i
			}
			i++
		}
		temp = A[pass-1].harga
		A[pass-1].harga = A[idx].harga
		A[idx].harga = temp
		pass++
	}
}
