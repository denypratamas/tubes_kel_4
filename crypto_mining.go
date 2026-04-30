package main

import (
	"fmt"
)

type CryptoAsset struct {
	Name       string
	Difficulty float64
	Mined      bool
}

var masterData []CryptoAsset // inisialisasi slice
var computingPower float64

func initData() {
	masterData = []CryptoAsset{
		{"Bitcoin", 100.0, false},
		{"Zcash", 10.0, false},
		{"Ethereum", 50.0, false},
		{"Litecoin", 25.0, false},
	}
}

func kelolaAset() {
	for {
		fmt.Println("\n--- DAFTAR ASET KRIPTO ---")
		for i, aset := range masterData {
			status := "belum ditambahkan"
			if aset.Mined {
				status = "sudah ditambahkan"
			}
			fmt.Printf("%d. %s - %.0f hashrate - %s\n", i+1, aset.Name, aset.Difficulty, status)
		}

		fmt.Println("\nKetik '+' untuk menambah aset dan '-' untuk hapus aset, atau '0' untuk kembali ke menu utama:")
		var aksi string
		fmt.Scan(&aksi)

		if aksi == "0" {
			break
		}
		if aksi == "+" || aksi == "-" {
			fmt.Print("Masukkan nomor aset: ")
			var nomor int
			fmt.Scanln(&nomor)

			index := nomor - 1
			if index >= 0 && index < len(masterData) {
				if aksi == "+" {
					masterData[index].Mined = true
					fmt.Printf("\n[SUKSES] %s berhasil ditambahkan ke daftar aset ditambang!\n", masterData[index].Name)
				} else if aksi == "-" {
					masterData[index].Mined = false
					fmt.Printf("\n[SUKSES] %s berhasil dihapus dari daftar aset ditambang!\n", masterData[index].Name)
				} else {
					fmt.Println("\n[ERROR] Nomor aset tidak valid!")
				}
			} else {
				fmt.Println("\n[ERROR] Input tidak dikenali!")
			}
		}
	}

}
