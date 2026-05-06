# Crypto Mining Simulation

Crypto Mining Simulation adalah aplikasi CLI sederhana berbasis Go untuk latihan CRUD, simulasi mining, searching, sorting, dan laporan hasil mining.

Program ini memakai data di memori dengan slice, sehingga cocok untuk belajar dasar pemrograman Go tanpa database.

## Cara Menjalankan

1. Buka terminal atau PowerShell.
2. Masuk ke folder project
3. Jalankan program:

```powershell
go run crypto_mining.go
```

## Menu Utama

Saat program dijalankan, menu berikut akan muncul:

```text
=== CRYPTO MINING SIMULATION ===
1. Kelola aset kripto
2. Simulasi mining
3. Cari aset
4. Urutkan aset
5. Laporan mining
0. Keluar
```

Ketik angka menu yang diinginkan, lalu tekan Enter.

## Data Awal

Program langsung menyediakan beberapa aset:

| Aset | Difficulty | Reward | Algoritma | Energi |
| --- | ---: | ---: | --- | ---: |
| Bitcoin | 100.00 | 0.005000 | SHA-256 | 1500.00 kWh |
| Zcash | 10.00 | 0.080000 | Equihash | 450.00 kWh |
| Ethereum | 50.00 | 0.020000 | Ethash | 900.00 kWh |
| Litecoin | 25.00 | 0.120000 | Scrypt | 650.00 kWh |

Status `belum pernah ditambang` akan berubah menjadi `pernah ditambang` setelah aset disimulasikan.

## Menu 1: Kelola Aset Kripto

Menu ini khusus untuk CRUD aset.

```text
1. Tambah aset baru
2. Edit aset
3. Hapus aset
0. Kembali ke menu utama
```

### Tambah Aset

Data yang diminta:

- Nama aset
- Mining difficulty
- Estimasi reward
- Tipe algoritma
- Estimasi konsumsi energi

Program sudah memakai `bufio.Scanner`, sehingga input string boleh memakai spasi.

Contoh:

```text
Nama aset: Bitcoin Cash
Mining difficulty: 20
Estimasi reward: 0.03
Tipe algoritma: SHA-256 Variant
Estimasi konsumsi energi (kWh): 700
```

### Edit Aset

Pilih nomor aset yang ingin diedit, lalu isi ulang data aset dengan nilai baru.

### Hapus Aset

Pilih nomor aset yang ingin dihapus. Data tersebut akan dihapus dari slice `masterData`.

## Menu 2: Simulasi Mining

Menu ini langsung menjalankan proses mining untuk satu aset.

Alur:

1. Program menampilkan katalog aset.
2. Masukkan nomor aset atau nama aset.
3. Masukkan computational power.
4. Program menampilkan teks `Sedang menambang [Nama Aset]...`.
5. Program memberi delay dinamis berdasarkan difficulty, algoritma, dan compute power.
6. Hasil simulasi ditampilkan dan disimpan ke laporan.

Contoh input:

```text
Masukkan nomor atau nama aset yang ingin ditambang: Bitcoin Cash
Masukkan computational power pengguna (hashrate): 1000
```

Rumus utama:

```text
duration = (difficulty * algoMultiplier) / computingPower
reward = estimatedReward * (computingPower / (difficulty + 100))
energyUsed = energyConsumption * duration
delay = (difficulty * algoMultiplier) / computingPower
```

Delay diskalakan dan dibatasi agar program tetap nyaman dipakai.

## Menu 3: Cari Aset

Pengguna cukup memasukkan nama aset.

```text
Masukkan nama aset: Bitcoin Cash
```

Di balik layar:

1. Program menyalin data aset.
2. Salinan data diurutkan berdasarkan nama.
3. Program mencari aset dengan Binary Search.

Binary Search membutuhkan data terurut dan memiliki kompleksitas `O(log n)`.

## Menu 4: Urutkan Aset

Menu ini hanya menampilkan pilihan berdasarkan kebutuhan pengguna, bukan nama algoritma.

```text
1. Urutkan berdasarkan Difficulty
2. Urutkan berdasarkan Reward
```

Di balik layar:

- Difficulty memakai Selection Sort.
- Reward memakai Insertion Sort.

## Menu 5: Laporan Mining

Laporan menampilkan:

- Total sesi mining
- Total estimasi reward
- Rata-rata durasi mining
- Total konsumsi energi
- Aset paling menguntungkan

Jika belum ada simulasi, program akan meminta pengguna menjalankan simulasi mining terlebih dahulu.

## Catatan Penting

- Data hanya tersimpan selama program berjalan.
- Program tidak memakai database atau file penyimpanan.
- Input kosong akan ditolak.
- Nama aset dan algoritma boleh memakai spasi (sudah pake bufio.Scanner)