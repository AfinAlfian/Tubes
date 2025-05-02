package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

const NMAX int = 1024

type tabInt [NMAX]int
type tabStr [NMAX]string

type profil struct {
	username string
	password string
	saldo    int
}

type kripto struct {
	nama  tabStr
	harga tabInt
}

type portofolio struct {
	keuntungan tabInt
	kerugian   tabInt
}

type riwayat struct {
	status  tabStr
	nkripto tabStr
	jumlah  tabInt
	harga   tabInt
	tanggal tabStr
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func login(profil *profil) {
	fmt.Println("-----------------------------------------")
	fmt.Println("|   Selamat datang di aplikasi kripto   |")
	fmt.Println("-----------------------------------------")
	fmt.Print("Masukkan username: ")
	fmt.Scan(&profil.username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&profil.password)
	clearScreen()
}

func main() {
	var p profil
	var k kripto
	var n int
	k.nama[0] = "Bitcoin"
	k.nama[1] = "Ethereum"
	k.nama[2] = "Litecoin"
	k.harga[0] = 1000000000
	k.harga[1] = 230000
	k.harga[2] = 87000000

	login(&p)
	fmt.Println("-----------------------------------------")
	fmt.Println("|              Menu Utama               |")
	fmt.Println("-----------------------------------------")
	fmt.Println("username:", p.username)
	fmt.Println("saldo:", p.saldo)
	fmt.Println("-----------------------------------------")
	for i := 0; i < 3; i++ {
		fmt.Println(i+1, k.nama[i], "Harga:", k.harga[i])
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
	fmt.Scan(&n)

}
