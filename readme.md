# Parking Lot Management System

Sebuah sistem manajemen tempat parkir yang dibangun menggunakan Go (Golang) dengan penerapan Domain-Driven Design (DDD) dan Clean Architecture. Aplikasi ini membaca perintah dari sebuah file input, mengeksekusi perintah tersebut, dan menampilkan hasilnya ke STDOUT.

---

## Struktur Projek

```
cmd/
  parking/
    main.go

internal/
  cli/
    handler.go

  domain/
    car.go
    parking_lot.go
    parking_manager.go

  repository/
    parking_repository.go

    memory/
      memory_parking_repository.go

  service/
    parking_service.go

go.mod
input.txt
readme.md
```

### Fungsi Setiap Folder

| Folder                         | Description                                                        |
| ------------------------------ | ------------------------------------------------------------------ |
| **cmd/parking**                | Application entry point. Initializes CLI, repository, and service. |
| **internal/cli**               | Handles command parsing and executes service methods.              |
| **internal/domain**            | Contains core business logic and domain entities.                  |
| **internal/repository**        | Repository abstraction (interfaces).                               |
| **internal/repository/memory** | In-memory repository implementation.                               |
| **internal/service**           | Application service orchestrating domain + repository.             |

---

## Cara menjalankan projek

Siapkan file input, cth: input.txt

```bash
go run cmd/parking/main.go input.txt
```

Example `input.txt`:

```
create_parking_lot 6
park KA-01-HH-1234
park KA-01-HH-9999
park KA-01-BB-0001
park KA-01-HH-7777
park KA-01-HH-2701
park KA-01-HH-3141
leave KA-01-HH-3141 4
status
```

---

## Business Rules

### Parking

* Sarankan urutan parkir lot terdekat

### Leaving

* Hitung biaya parkir
* Aturan Biaya:

  * 2 Jam pertama: $10
  * +$10 untuk setiap penambahan jam

### Status

* Data setiap slot parkir.

---
## Contoh Output

```
Allocated slot number: 1
Allocated slot number: 2
Slot No.   Registration No.
1          KA-01-HH-1111
2          KA-01-HH-2222
```

---

## Alur Pengerjaan Proyek

### 1. Membuat Domain Model

* Membuat entity **Car**, **ParkingLot**, dan **ParkingLotManager**.
* Membuat fungsi untuk masing" domain: `Park`, `Leave`, `IsEmpty`.
* Menambahkan logic seperti nearest-empty-slot dan perhitungan biaya parkir.

### 2. Membuat Repository Abstraction

* Membuat interface `ParkingRepository`.
* Menentukan method `Save()` dan `Load()` untuk menyimpan state manager.

### 3. Implementasi Repository (In-Memory)

* Membuat folder `internal/repository/memory`.
* Implementasi penyimpanan sederhana menggunakan pointer ke `ParkingLotManager`.

### 5. Membuat Service Layer

* Membuat `ParkingService` sebagai jembatan antara CLI dan domain.
* Service menangani aksi seperti `CreateParkingLot`, `Park`, `Leave`, dan `Status`.

### 6. Membuat CLI Handler

* Parsing command dari input file.
* Mengarahkan command ke service.
* Menampilkan output sesuai format test case.

### 7. Menyusun `main.go`

* Melakukan dependency injection.
* Membuat service, repository, CLI.
* Menjalankan `CLI.RunFromFile()`.

### 8. Menjalankan dan Menguji Menggunakan Test Case PDF

