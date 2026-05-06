package main

import "fmt"

type CryptoAsset struct {
	Name              string
	Difficulty        float64
	EstimatedReward   float64
	Algorithm         string
	EnergyConsumption float64
	Mined             bool
}

type MiningSession struct {
	AssetName          string
	Duration           float64
	Reward             float64
	EnergyUsed         float64
	ComputingPowerUsed float64
}

var masterData []CryptoAsset // inisialisasi slice
var miningSessions []MiningSession
var computingPower float64

func initData() {
	masterData = []CryptoAsset{
		{"Bitcoin", 100.0, 0.005, "SHA-256", 1500.0, false},
		{"Zcash", 10.0, 0.08, "Equihash", 450.0, false},
		{"Ethereum", 50.0, 0.02, "Ethash", 900.0, false},
		{"Litecoin", 25.0, 0.12, "Scrypt", 650.0, false},
	}
}

func main() {
	initData()

	for {
		fmt.Println("\n=== CRYPTO MINING SIMULATION ===")
		fmt.Println("1. Kelola aset kripto")
		fmt.Println("2. Simulasi mining")
		fmt.Println("3. Cari aset")
		fmt.Println("4. Urutkan aset")
		fmt.Println("5. Laporan mining")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			kelolaAset()
		case 2:
			simulasiMining()
		case 3:
			menuCariAset()
		case 4:
			menuUrutkanAset()
		case 5:
			tampilkanLaporan()
		case 0:
			fmt.Println("\nTerima kasih sudah menggunakan aplikasi simulasi mining.")
			return
		default:
			fmt.Println("\n[ERROR] Pilihan menu tidak valid!")
		}
	}
}

func kelolaAset() {
	for {
		fmt.Println("\n--- KELOLA ASET KRIPTO ---")
		tampilkanSemuaAset()
		fmt.Println("\n1. Tambah aset baru")
		fmt.Println("2. Edit aset")
		fmt.Println("3. Hapus aset")
		fmt.Println("4. Tandai aset untuk ditambang")
		fmt.Println("5. Batalkan aset dari daftar tambang")
		fmt.Println("0. Kembali ke menu utama")
		fmt.Print("Pilih aksi: ")

		var aksi int
		fmt.Scanln(&aksi)

		switch aksi {
		case 1:
			tambahAset()
		case 2:
			editAset()
		case 3:
			hapusAset()
		case 4:
			ubahStatusMining(true)
		case 5:
			ubahStatusMining(false)
		case 0:
			return
		default:
			fmt.Println("\n[ERROR] Input tidak dikenali!")
		}
	}
}

func tampilkanSemuaAset() {
	if len(masterData) == 0 {
		fmt.Println("\nBelum ada aset kripto.")
		return
	}

	fmt.Println("\n--- DAFTAR ASET KRIPTO ---")
	for i, aset := range masterData {
		tampilkanRingkasanAset(i, aset)
	}
}

func tampilkanRingkasanAset(index int, aset CryptoAsset) {
	status := "belum dipilih"
	if aset.Mined {
		status = "dipilih untuk mining"
	}

	fmt.Printf("%d. %s | Difficulty: %.2f | Reward: %.6f | Algoritma: %s | Energi: %.2f kWh | %s\n",
		index+1,
		aset.Name,
		aset.Difficulty,
		aset.EstimatedReward,
		aset.Algorithm,
		aset.EnergyConsumption,
		status,
	)
}

func tambahAset() {
	var aset CryptoAsset

	fmt.Println("\n--- TAMBAH ASET BARU ---")
	fmt.Print("Nama aset: ")
	fmt.Scanln(&aset.Name)
	fmt.Print("Mining difficulty: ")
	fmt.Scanln(&aset.Difficulty)
	fmt.Print("Estimasi reward: ")
	fmt.Scanln(&aset.EstimatedReward)
	fmt.Print("Tipe algoritma: ")
	fmt.Scanln(&aset.Algorithm)
	fmt.Print("Estimasi konsumsi energi (kWh): ")
	fmt.Scanln(&aset.EnergyConsumption)

	if !validasiAset(aset) {
		return
	}

	masterData = append(masterData, aset)
	fmt.Printf("\n[SUKSES] Aset %s berhasil ditambahkan.\n", aset.Name)
}

func editAset() {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset yang bisa diedit.")
		return
	}

	fmt.Print("\nMasukkan nomor aset yang ingin diedit: ")
	var nomor int
	fmt.Scanln(&nomor)

	index := nomor - 1
	if !validasiIndex(index) {
		return
	}

	asetBaru := masterData[index]
	fmt.Printf("Nama aset baru (%s): ", asetBaru.Name)
	fmt.Scanln(&asetBaru.Name)
	fmt.Print("Mining difficulty baru: ")
	fmt.Scanln(&asetBaru.Difficulty)
	fmt.Print("Estimasi reward baru: ")
	fmt.Scanln(&asetBaru.EstimatedReward)
	fmt.Print("Tipe algoritma baru: ")
	fmt.Scanln(&asetBaru.Algorithm)
	fmt.Print("Estimasi konsumsi energi baru (kWh): ")
	fmt.Scanln(&asetBaru.EnergyConsumption)

	if !validasiAset(asetBaru) {
		return
	}

	masterData[index] = asetBaru
	fmt.Printf("\n[SUKSES] Aset nomor %d berhasil diedit.\n", nomor)
}

func hapusAset() {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset yang bisa dihapus.")
		return
	}

	fmt.Print("\nMasukkan nomor aset yang ingin dihapus: ")
	var nomor int
	fmt.Scanln(&nomor)

	index := nomor - 1
	if !validasiIndex(index) {
		return
	}

	namaAset := masterData[index].Name
	masterData = append(masterData[:index], masterData[index+1:]...)
	fmt.Printf("\n[SUKSES] Aset %s berhasil dihapus.\n", namaAset)
}

func ubahStatusMining(dipilih bool) {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset yang tersedia.")
		return
	}

	fmt.Print("Masukkan nomor aset: ")
	var nomor int
	fmt.Scanln(&nomor)

	index := nomor - 1
	if !validasiIndex(index) {
		return
	}

	masterData[index].Mined = dipilih
	if dipilih {
		fmt.Printf("\n[SUKSES] %s berhasil ditambahkan ke daftar aset ditambang!\n", masterData[index].Name)
	} else {
		fmt.Printf("\n[SUKSES] %s berhasil dihapus dari daftar aset ditambang!\n", masterData[index].Name)
	}
}

func validasiAset(aset CryptoAsset) bool {
	if teksKosong(aset.Name) || teksKosong(aset.Algorithm) {
		fmt.Println("\n[ERROR] Nama aset dan algoritma tidak boleh kosong.")
		return false
	}
	if aset.Difficulty <= 0 || aset.EstimatedReward <= 0 || aset.EnergyConsumption <= 0 {
		fmt.Println("\n[ERROR] Difficulty, reward, dan konsumsi energi harus lebih dari 0.")
		return false
	}
	return true
}

func validasiIndex(index int) bool {
	if index < 0 || index >= len(masterData) {
		fmt.Println("\n[ERROR] Nomor aset tidak valid!")
		return false
	}
	return true
}

func simulasiMining() {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset untuk disimulasikan.")
		return
	}

	fmt.Print("\nMasukkan computational power pengguna (hashrate): ")
	fmt.Scanln(&computingPower)

	if computingPower <= 0 {
		fmt.Println("\n[ERROR] Computational power harus lebih dari 0.")
		return
	}

	adaAsetDipilih := false
	fmt.Println("\n--- HASIL SIMULASI MINING ---")

	for _, aset := range masterData {
		if !aset.Mined {
			continue
		}

		adaAsetDipilih = true

		var algoMultiplier float64
		switch aset.Algorithm {
		case "SHA-256":
			algoMultiplier = 2.0
		case "Ethash":
			algoMultiplier = 1.5
		case "Scrypt":
			algoMultiplier = 1.2
		default:
			algoMultiplier = 1.0
		}

		// Rumus sederhana:
		// Semakin tinggi difficulty, semakin lama waktu mining.
		// Semakin tinggi computingPower, semakin cepat waktu mining.
		duration := (aset.Difficulty * algoMultiplier) / computingPower

		// Reward dibuat menurun jika difficulty tinggi agar simulasi terasa realistis.
		// Angka 100 digunakan sebagai skala sederhana untuk pemula.
		reward := aset.EstimatedReward * (computingPower / (aset.Difficulty + 100))

		// Energi dihitung dari konsumsi energi aset dikali durasi mining.
		energyUsed := aset.EnergyConsumption * duration

		session := MiningSession{
			AssetName:          aset.Name,
			Duration:           duration,
			Reward:             reward,
			EnergyUsed:         energyUsed,
			ComputingPowerUsed: computingPower,
		}
		miningSessions = append(miningSessions, session)

		fmt.Printf("\nAset: %s\n", aset.Name)
		fmt.Printf("Algoritma: %s\n", aset.Algorithm)
		fmt.Printf("Estimasi durasi mining: %.2f jam\n", duration)
		fmt.Printf("Computational power digunakan: %.2f hashrate\n", computingPower)
		fmt.Printf("Estimasi energi digunakan: %.2f kWh\n", energyUsed)
		fmt.Printf("Estimasi reward diterima: %.6f\n", reward)
	}

	if !adaAsetDipilih {
		fmt.Println("\n[INFO] Belum ada aset yang dipilih untuk mining.")
	}
}

func menuCariAset() {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset untuk dicari.")
		return
	}

	fmt.Println("\n--- CARI ASET ---")
	fmt.Println("1. Sequential search")
	fmt.Println("2. Binary search")
	fmt.Print("Pilih metode pencarian: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	fmt.Print("Masukkan nama aset: ")
	var keyword string
	fmt.Scanln(&keyword)

	switch pilihan {
	case 1:
		index := sequentialSearch(keyword)
		tampilkanHasilPencarian(index)
	case 2:
		salinanData := salinData(masterData)
		insertionSortByName(salinanData)

		// Binary search memiliki kompleksitas O(log n), tetapi hanya benar
		// jika data sudah terurut. Karena itu data disalin lalu diurutkan dulu.
		index := binarySearch(salinanData, keyword)
		if index == -1 {
			fmt.Println("\n[INFO] Aset tidak ditemukan.")
		} else {
			fmt.Println("\n[SUKSES] Aset ditemukan pada data yang sudah diurutkan berdasarkan nama:")
			tampilkanRingkasanAset(index, salinanData[index])
		}
	default:
		fmt.Println("\n[ERROR] Metode pencarian tidak valid!")
	}
}

func sequentialSearch(keyword string) int {
	// Sequential search mengecek data satu per satu dari awal sampai akhir.
	// Kompleksitas waktunya O(n) karena pada kasus terburuk semua data dicek.
	for i, aset := range masterData {
		if samaTanpaHurufBesarKecil(aset.Name, keyword) {
			return i
		}
	}
	return -1
}

func binarySearch(data []CryptoAsset, keyword string) int {
	kiri := 0
	kanan := len(data) - 1
	keyword = ubahKeHurufKecil(keyword)

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		namaAset := ubahKeHurufKecil(data[tengah].Name)

		if namaAset == keyword {
			return tengah
		}
		if namaAset < keyword {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func tampilkanHasilPencarian(index int) {
	if index == -1 {
		fmt.Println("\n[INFO] Aset tidak ditemukan.")
		return
	}

	fmt.Println("\n[SUKSES] Aset ditemukan:")
	tampilkanRingkasanAset(index, masterData[index])
}

func menuUrutkanAset() {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset untuk diurutkan.")
		return
	}

	fmt.Println("\n--- URUTKAN ASET ---")
	fmt.Println("1. Selection sort berdasarkan difficulty")
	fmt.Println("2. Selection sort berdasarkan reward")
	fmt.Println("3. Insertion sort berdasarkan difficulty")
	fmt.Println("4. Insertion sort berdasarkan reward")
	fmt.Print("Pilih metode pengurutan: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		selectionSortByDifficulty(masterData)
	case 2:
		selectionSortByReward(masterData)
	case 3:
		insertionSortByDifficulty(masterData)
	case 4:
		insertionSortByReward(masterData)
	default:
		fmt.Println("\n[ERROR] Metode pengurutan tidak valid!")
		return
	}

	fmt.Println("\n[SUKSES] Data aset berhasil diurutkan.")
	tampilkanSemuaAset()
}

func selectionSortByDifficulty(data []CryptoAsset) {
	// Selection sort mencari nilai paling kecil, lalu menukarnya ke posisi awal.
	// Proses ini diulang untuk setiap posisi. Kompleksitas waktunya O(n^2).
	for i := 0; i < len(data)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j].Difficulty < data[minIndex].Difficulty {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}

func selectionSortByReward(data []CryptoAsset) {
	// Selection sort tetap O(n^2), di sini pembandingnya adalah reward terbesar.
	for i := 0; i < len(data)-1; i++ {
		maxIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j].EstimatedReward > data[maxIndex].EstimatedReward {
				maxIndex = j
			}
		}
		data[i], data[maxIndex] = data[maxIndex], data[i]
	}
}

func insertionSortByDifficulty(data []CryptoAsset) {
	// Insertion sort menyisipkan setiap data ke posisi yang tepat
	// pada bagian kiri slice yang sudah dianggap terurut. Kompleksitas O(n^2).
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j].Difficulty > key.Difficulty {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func insertionSortByReward(data []CryptoAsset) {
	// Insertion sort tetap O(n^2), di sini data disisipkan berdasarkan reward terbesar.
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j].EstimatedReward < key.EstimatedReward {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func insertionSortByName(data []CryptoAsset) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && ubahKeHurufKecil(data[j].Name) > ubahKeHurufKecil(key.Name) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func salinData(data []CryptoAsset) []CryptoAsset {
	hasil := make([]CryptoAsset, len(data))
	copy(hasil, data)
	return hasil
}

func teksKosong(teks string) bool {
	for _, huruf := range teks {
		if huruf != ' ' && huruf != '\t' && huruf != '\n' && huruf != '\r' {
			return false
		}
	}
	return true
}

func samaTanpaHurufBesarKecil(teks1 string, teks2 string) bool {
	return ubahKeHurufKecil(teks1) == ubahKeHurufKecil(teks2)
}

func ubahKeHurufKecil(teks string) string {
	hasil := []rune(teks)
	for i, huruf := range hasil {
		if huruf >= 'A' && huruf <= 'Z' {
			hasil[i] = huruf + 32
		}
	}
	return string(hasil)
}

func tampilkanLaporan() {
	fmt.Println("\n--- LAPORAN MINING ---")
	if len(miningSessions) == 0 {
		fmt.Println("[INFO] Belum ada sesi mining. Jalankan simulasi mining terlebih dahulu.")
		return
	}

	var totalReward float64
	var totalDuration float64
	var totalEnergy float64
	mostProfitable := miningSessions[0]

	for _, session := range miningSessions {
		totalReward += session.Reward
		totalDuration += session.Duration
		totalEnergy += session.EnergyUsed

		if session.Reward > mostProfitable.Reward {
			mostProfitable = session
		}
	}

	averageDuration := totalDuration / float64(len(miningSessions))

	fmt.Printf("Total sesi mining: %d\n", len(miningSessions))
	fmt.Printf("Total estimasi reward: %.6f\n", totalReward)
	fmt.Printf("Rata-rata durasi mining: %.2f jam\n", averageDuration)
	fmt.Printf("Total konsumsi energi: %.2f kWh\n", totalEnergy)
	fmt.Printf("Aset paling menguntungkan: %s dengan reward %.6f\n", mostProfitable.AssetName, mostProfitable.Reward)
}
