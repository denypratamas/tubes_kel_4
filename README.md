# Crypto Mining Simulation

Crypto Mining Simulation adalah aplikasi CLI sederhana yang dibuat dengan Go. Program ini membantu pengguna memahami konsep dasar mining cryptocurrency melalui data aset, simulasi mining, pencarian, pengurutan, dan laporan hasil mining.

Program ini cocok untuk latihan CRUD pertama karena semua data disimpan di dalam slice dan dijalankan melalui menu terminal.

## Isi Program

Program memiliki fitur utama:

- Mengelola data aset crypto
- Menandai aset yang ingin ditambang
- Menjalankan simulasi mining
- Mencari aset dengan sequential search dan binary search
- Mengurutkan aset dengan selection sort dan insertion sort
- Menampilkan laporan hasil mining

## Cara Menjalankan Program

1. Buka terminal atau PowerShell.
2. Masuk ke folder project:
3. Jalankan program:

```powershell
go run crypto_mining.go
```

4. Setelah program berjalan, menu utama akan muncul:

```text
=== CRYPTO MINING SIMULATION ===
1. Kelola aset kripto
2. Simulasi mining
3. Cari aset
4. Urutkan aset
5. Laporan mining
0. Keluar
```

Ketik angka menu yang ingin dipilih, lalu tekan Enter.

## Data Awal

Saat program pertama kali dijalankan, data aset crypto sudah diisi otomatis:

| Aset | Difficulty | Reward | Algoritma | Energi |
| --- | ---: | ---: | --- | ---: |
| Bitcoin | 100.00 | 0.005000 | SHA-256 | 1500.00 kWh |
| Zcash | 10.00 | 0.080000 | Equihash | 450.00 kWh |
| Ethereum | 50.00 | 0.020000 | Ethash | 900.00 kWh |
| Litecoin | 25.00 | 0.120000 | Scrypt | 650.00 kWh |

Status awal semua aset adalah `belum dipilih`, artinya aset belum dipilih untuk simulasi mining.

## Menu 1: Kelola Aset Kripto

Menu ini digunakan untuk melakukan CRUD dan memilih aset yang akan ditambang.

Pilihan di dalam menu:

```text
1. Tambah aset baru
2. Edit aset
3. Hapus aset
4. Tandai aset untuk ditambang
5. Batalkan aset dari daftar tambang
0. Kembali ke menu utama
```

### 1. Tambah Aset Baru

Gunakan menu ini untuk menambahkan aset crypto baru.

Data yang harus diisi:

- Nama aset
- Mining difficulty
- Estimasi reward
- Tipe algoritma
- Estimasi konsumsi energi dalam kWh

Contoh input:

```text
Nama aset: Dogecoin
Mining difficulty: 15
Estimasi reward: 0.05
Tipe algoritma: Scrypt
Estimasi konsumsi energi (kWh): 500
```

Catatan:

- Difficulty, reward, dan energi harus lebih dari 0.
- Nama aset dan algoritma tidak boleh kosong.

### 2. Edit Aset

Gunakan menu ini untuk mengubah data aset yang sudah ada.

Langkah:

1. Pilih menu `2. Edit aset`.
2. Masukkan nomor aset yang ingin diedit.
3. Isi ulang data aset dengan nilai baru.

Program akan mengganti data lama dengan data baru.

### 3. Hapus Aset

Gunakan menu ini untuk menghapus aset dari daftar.

Langkah:

1. Pilih menu `3. Hapus aset`.
2. Masukkan nomor aset yang ingin dihapus.
3. Program akan menghapus aset tersebut dari slice `masterData`.

### 4. Tandai Aset Untuk Ditambang

Gunakan menu ini sebelum menjalankan simulasi mining.

Langkah:

1. Pilih menu `4. Tandai aset untuk ditambang`.
2. Masukkan nomor aset.
3. Status aset berubah menjadi `dipilih untuk mining`.

Hanya aset dengan status `dipilih untuk mining` yang akan diproses oleh simulasi mining.

### 5. Batalkan Aset Dari Daftar Tambang

Gunakan menu ini jika aset tidak ingin ikut disimulasikan.

Langkah:

1. Pilih menu `5. Batalkan aset dari daftar tambang`.
2. Masukkan nomor aset.
3. Status aset kembali menjadi `belum dipilih`.

## Menu 2: Simulasi Mining

Menu ini digunakan untuk menghitung estimasi hasil mining berdasarkan aset yang sudah dipilih.

Langkah penggunaan:

1. Masuk ke `Kelola aset kripto`.
2. Tandai satu atau beberapa aset untuk ditambang.
3. Kembali ke menu utama.
4. Pilih `2. Simulasi mining`.
5. Masukkan computational power pengguna.

Contoh:

```text
Masukkan computational power pengguna (hashrate): 50
```

Program akan menampilkan:

- Nama aset
- Algoritma
- Estimasi durasi mining
- Computational power yang digunakan
- Estimasi energi yang digunakan
- Estimasi reward yang diterima

### Rumus Simulasi

Program menggunakan rumus sederhana:

```text
duration = difficulty / computingPower
reward = estimatedReward * (computingPower / (difficulty + 100))
energyUsed = energyConsumption * duration
```

Arti rumus:

- Semakin besar difficulty, durasi mining semakin lama.
- Semakin besar computational power, durasi mining semakin cepat.
- Reward dipengaruhi oleh reward awal, difficulty, dan computational power.
- Energi dihitung dari konsumsi energi aset dikali durasi mining.

Hasil simulasi juga disimpan ke `miningSessions` agar bisa muncul di laporan.

## Menu 3: Cari Aset

Menu ini digunakan untuk mencari aset berdasarkan nama.

Pilihan pencarian:

```text
1. Sequential search
2. Binary search
```

### Sequential Search

Sequential search mengecek data satu per satu dari awal sampai akhir.

Contoh:

```text
Pilih metode pencarian: 1
Masukkan nama aset: Bitcoin
```

Jika aset ditemukan, program menampilkan detail aset tersebut.

Kompleksitas waktu:

```text
O(n)
```

Artinya, semakin banyak data, pencarian bisa semakin lama karena data diperiksa satu per satu.

### Binary Search

Binary search mencari data dengan cara membagi area pencarian menjadi dua bagian.

Syarat penting:

- Binary search hanya benar jika data sudah terurut.
- Di program ini, data disalin lalu diurutkan berdasarkan nama sebelum binary search dijalankan.
- Data asli di `masterData` tidak berubah saat binary search.

Kompleksitas waktu:

```text
O(log n)
```

Artinya, binary search lebih cepat untuk data besar, tetapi membutuhkan data yang sudah terurut.

## Menu 4: Urutkan Aset

Menu ini digunakan untuk mengurutkan data aset.

Pilihan pengurutan:

```text
1. Selection sort berdasarkan difficulty
2. Selection sort berdasarkan reward
3. Insertion sort berdasarkan difficulty
4. Insertion sort berdasarkan reward
```

### Selection Sort Berdasarkan Difficulty

Mengurutkan aset dari difficulty terkecil ke terbesar.

Cara kerja:

1. Cari difficulty paling kecil.
2. Tukar ke posisi paling depan.
3. Ulangi untuk posisi berikutnya.

Kompleksitas waktu:

```text
O(n^2)
```

### Selection Sort Berdasarkan Reward

Mengurutkan aset dari reward terbesar ke terkecil.

Cara kerja sama seperti selection sort biasa, tetapi pembandingnya adalah `EstimatedReward`.

### Insertion Sort Berdasarkan Difficulty

Mengurutkan aset dari difficulty terkecil ke terbesar.

Cara kerja:

1. Ambil satu data sebagai `key`.
2. Geser data yang lebih besar ke kanan.
3. Sisipkan `key` ke posisi yang tepat.

Kompleksitas waktu:

```text
O(n^2)
```

### Insertion Sort Berdasarkan Reward

Mengurutkan aset dari reward terbesar ke terkecil.

Cara kerja sama seperti insertion sort biasa, tetapi pembandingnya adalah `EstimatedReward`.

## Menu 5: Laporan Mining

Menu ini menampilkan ringkasan semua simulasi mining yang sudah dijalankan.

Isi laporan:

- Total sesi mining
- Total estimasi reward
- Rata-rata durasi mining
- Total konsumsi energi
- Aset paling menguntungkan

Jika belum pernah menjalankan simulasi mining, program akan menampilkan pesan:

```text
[INFO] Belum ada sesi mining. Jalankan simulasi mining terlebih dahulu.
```

## Alur Penggunaan Yang Disarankan

Untuk mencoba program dari awal, gunakan urutan ini:

1. Jalankan program dengan `go run crypto_mining.go`.
2. Pilih menu `1. Kelola aset kripto`.
3. Tandai aset yang ingin ditambang.
4. Kembali ke menu utama.
5. Pilih menu `2. Simulasi mining`.
6. Masukkan computational power.
7. Pilih menu `5. Laporan mining`.
8. Coba menu `3. Cari aset`.
9. Coba menu `4. Urutkan aset`.
10. Pilih `0. Keluar` jika sudah selesai.

## Contoh Skenario Singkat

Input:

```text
1
4
1
0
2
50
5
0
```

Arti input:

1. Masuk ke menu kelola aset.
2. Tandai aset untuk ditambang.
3. Pilih aset nomor 1, yaitu Bitcoin.
4. Kembali ke menu utama.
5. Jalankan simulasi mining.
6. Masukkan computational power 50.
7. Lihat laporan mining.
8. Keluar dari program.

## Catatan Penting

- Data hanya disimpan di memori selama program berjalan.
- Jika program ditutup, data tambahan dan laporan mining akan hilang.
- Program ini tidak menggunakan database atau file penyimpanan.
- Input nama aset dan algoritma sebaiknya tidak memakai spasi karena program memakai `fmt.Scanln`.
- Program dibuat sederhana agar mudah dipahami oleh pemula.

## Struktur File

Project ini hanya memakai satu file Go utama:

```text
crypto_mining.go
```

File tersebut berisi:

- Struct `CryptoAsset`
- Struct `MiningSession`
- Data awal aset crypto
- Menu utama
- Fungsi CRUD aset
- Fungsi simulasi mining
- Fungsi searching
- Fungsi sorting
- Fungsi laporan mining

