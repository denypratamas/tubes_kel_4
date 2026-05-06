package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

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
var inputScanner = bufio.NewScanner(os.Stdin)

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

		pilihan := bacaInt("Pilih menu: ")

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
		fmt.Println("0. Kembali ke menu utama")

		aksi := bacaInt("Pilih aksi: ")

		switch aksi {
		case 1:
			tambahAset()
		case 2:
			editAset()
		case 3:
			hapusAset()
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
	status := "belum pernah ditambang"
	if aset.Mined {
		status = "pernah ditambang"
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
	fmt.Println("\n--- TAMBAH ASET BARU ---")

	aset := CryptoAsset{
		Name:              bacaStringWajib("Nama aset: "),
		Difficulty:        bacaFloatPositif("Mining difficulty: "),
		EstimatedReward:   bacaFloatPositif("Estimasi reward: "),
		Algorithm:         bacaStringWajib("Tipe algoritma: "),
		EnergyConsumption: bacaFloatPositif("Estimasi konsumsi energi (kWh): "),
		Mined:             false,
	}

	masterData = append(masterData, aset)
	fmt.Printf("\n[SUKSES] Aset %s berhasil ditambahkan.\n", aset.Name)
}

func editAset() {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset yang bisa diedit.")
		return
	}

	nomor := bacaInt("\nMasukkan nomor aset yang ingin diedit: ")
	index := nomor - 1
	if !validasiIndex(index) {
		return
	}

	asetBaru := masterData[index]
	fmt.Println("\nMasukkan data baru untuk aset ini.")
	asetBaru.Name = bacaStringWajib("Nama aset baru: ")
	asetBaru.Difficulty = bacaFloatPositif("Mining difficulty baru: ")
	asetBaru.EstimatedReward = bacaFloatPositif("Estimasi reward baru: ")
	asetBaru.Algorithm = bacaStringWajib("Tipe algoritma baru: ")
	asetBaru.EnergyConsumption = bacaFloatPositif("Estimasi konsumsi energi baru (kWh): ")

	masterData[index] = asetBaru
	fmt.Printf("\n[SUKSES] Aset nomor %d berhasil diedit.\n", nomor)
}

func hapusAset() {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset yang bisa dihapus.")
		return
	}

	nomor := bacaInt("\nMasukkan nomor aset yang ingin dihapus: ")
	index := nomor - 1
	if !validasiIndex(index) {
		return
	}

	namaAset := masterData[index].Name
	masterData = append(masterData[:index], masterData[index+1:]...)
	fmt.Printf("\n[SUKSES] Aset %s berhasil dihapus.\n", namaAset)
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

	fmt.Println("\n--- SIMULASI MINING ---")
	tampilkanSemuaAset()

	assetIndex := pilihAsetUntukMining()
	if assetIndex == -1 {
		fmt.Println("\n[ERROR] Aset tidak ditemukan.")
		return
	}

	computingPower = bacaFloatPositif("Masukkan computational power pengguna (hashrate): ")
	aset := masterData[assetIndex]
	algoMultiplier := hitungAlgoMultiplier(aset.Algorithm)

	// Rumus sederhana:
	// Semakin tinggi difficulty, semakin lama estimasi durasi mining.
	// Semakin tinggi computingPower, semakin cepat estimasi durasi mining.
	// Algo multiplier membuat algoritma yang lebih berat terasa lebih sulit.
	duration := (aset.Difficulty * algoMultiplier) / computingPower

	// Reward dibuat menurun jika difficulty tinggi agar simulasi terasa realistis.
	// Angka 100 digunakan sebagai skala sederhana untuk pemula.
	reward := aset.EstimatedReward * (computingPower / (aset.Difficulty + 100))

	// Energi dihitung dari konsumsi energi aset dikali durasi mining.
	energyUsed := aset.EnergyConsumption * duration
	delay := hitungDelayMining(aset, computingPower, algoMultiplier)

	fmt.Printf("\nSedang menambang %s...\n", aset.Name)
	time.Sleep(delay)

	session := MiningSession{
		AssetName:          aset.Name,
		Duration:           duration,
		Reward:             reward,
		EnergyUsed:         energyUsed,
		ComputingPowerUsed: computingPower,
	}
	miningSessions = append(miningSessions, session)
	masterData[assetIndex].Mined = true

	fmt.Println("\n--- HASIL SIMULASI MINING ---")
	fmt.Printf("Aset: %s\n", aset.Name)
	fmt.Printf("Algoritma: %s\n", aset.Algorithm)
	fmt.Printf("Estimasi durasi mining: %.2f jam\n", duration)
	fmt.Printf("Computational power digunakan: %.2f hashrate\n", computingPower)
	fmt.Printf("Estimasi energi digunakan: %.2f kWh\n", energyUsed)
	fmt.Printf("Estimasi reward diterima: %.6f\n", reward)
}

func pilihAsetUntukMining() int {
	input := bacaStringWajib("\nMasukkan nomor atau nama aset yang ingin ditambang: ")

	nomor, err := strconv.Atoi(input)
	if err == nil {
		index := nomor - 1
		if validasiIndex(index) {
			return index
		}
		return -1
	}

	return sequentialSearch(input)
}

func hitungAlgoMultiplier(algorithm string) float64 {
	switch ubahKeHurufKecil(algorithm) {
	case "sha-256":
		return 2.0
	case "ethash":
		return 1.5
	case "scrypt":
		return 1.2
	case "equihash":
		return 1.1
	default:
		return 1.0
	}
}

func hitungDelayMining(aset CryptoAsset, power float64, algoMultiplier float64) time.Duration {
	// Delay mengikuti rumus dasar: (difficulty * algoMultiplier) / computingPower.
	// Nilainya kemudian diskalakan dan dibatasi agar simulasi terasa nyata,
	// tetapi program tetap nyaman dipakai dan tidak berhenti terlalu lama.
	rawDelay := (aset.Difficulty * algoMultiplier) / power
	delaySeconds := rawDelay * 1.5

	if delaySeconds < 1 {
		delaySeconds = 1
	}
	if delaySeconds > 7 {
		delaySeconds = 7
	}

	return time.Duration(delaySeconds * float64(time.Second))
}

func menuCariAset() {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset untuk dicari.")
		return
	}

	fmt.Println("\n--- CARI ASET ---")
	keyword := bacaStringWajib("Masukkan nama aset: ")

	salinanData := salinData(masterData)
	insertionSortByName(salinanData)

	// Binary search memiliki kompleksitas O(log n), tetapi hanya benar
	// jika data sudah terurut. Karena itu data disalin lalu diurutkan dulu.
	index := binarySearch(salinanData, keyword)
	if index == -1 {
		fmt.Println("\n[INFO] Aset tidak ditemukan.")
		return
	}

	fmt.Println("\n[SUKSES] Aset ditemukan:")
	tampilkanRingkasanAset(index, salinanData[index])
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

func menuUrutkanAset() {
	if len(masterData) == 0 {
		fmt.Println("\n[ERROR] Belum ada aset untuk diurutkan.")
		return
	}

	fmt.Println("\n--- URUTKAN ASET ---")
	fmt.Println("1. Urutkan berdasarkan Difficulty")
	fmt.Println("2. Urutkan berdasarkan Reward")

	pilihan := bacaInt("Pilih menu pengurutan: ")

	switch pilihan {
	case 1:
		selectionSortByDifficulty(masterData)
		fmt.Println("\n[SUKSES] Data aset diurutkan berdasarkan difficulty dengan Selection Sort.")
	case 2:
		insertionSortByReward(masterData)
		fmt.Println("\n[SUKSES] Data aset diurutkan berdasarkan reward dengan Insertion Sort.")
	default:
		fmt.Println("\n[ERROR] Pilihan pengurutan tidak valid!")
		return
	}

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

func insertionSortByReward(data []CryptoAsset) {
	// Insertion sort menyisipkan data ke posisi yang tepat pada bagian kiri
	// slice yang sudah dianggap terurut. Di sini reward terbesar ditempatkan dulu.
	// Kompleksitas waktunya O(n^2).
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

func bacaString(prompt string) string {
	fmt.Print(prompt)
	if !inputScanner.Scan() {
		return ""
	}
	return trimSpasi(inputScanner.Text())
}

func bacaStringWajib(prompt string) string {
	for {
		input := bacaString(prompt)
		if !teksKosong(input) {
			return input
		}
		fmt.Println("[ERROR] Input tidak boleh kosong.")
	}
}

func bacaInt(prompt string) int {
	for {
		input := bacaStringWajib(prompt)
		nilai, err := strconv.Atoi(input)
		if err == nil {
			return nilai
		}
		fmt.Println("[ERROR] Masukkan angka bulat yang valid.")
	}
}

func bacaFloatPositif(prompt string) float64 {
	for {
		input := bacaStringWajib(prompt)
		nilai, err := strconv.ParseFloat(input, 64)
		if err == nil && nilai > 0 {
			return nilai
		}
		fmt.Println("[ERROR] Masukkan angka lebih dari 0.")
	}
}

func trimSpasi(teks string) string {
	awal := 0
	akhir := len(teks) - 1

	for awal <= akhir && karakterSpasi(teks[awal]) {
		awal++
	}
	for akhir >= awal && karakterSpasi(teks[akhir]) {
		akhir--
	}
	if awal > akhir {
		return ""
	}
	return teks[awal : akhir+1]
}

func teksKosong(teks string) bool {
	return trimSpasi(teks) == ""
}

func karakterSpasi(karakter byte) bool {
	return karakter == ' ' || karakter == '\t' || karakter == '\n' || karakter == '\r'
}

func samaTanpaHurufBesarKecil(teks1 string, teks2 string) bool {
	return ubahKeHurufKecil(trimSpasi(teks1)) == ubahKeHurufKecil(trimSpasi(teks2))
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
