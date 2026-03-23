package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Driver postgres
)

func main() {
	// 1. Connection String (Mirip JDBC URL)
	connStr := "postgresql://postgres:sinarmulia46@localhost:5432/postgres?sslmode=disable"

	// 2. Buka Koneksi
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err) // Di Go, kita cek error secara eksplisit, bukan try-catch
	}
	defer db.Close() // Pastikan koneksi tutup saat aplikasi selesai (mirip finally)

	// 3. Tes Koneksi (Ping)
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	fmt.Println("Mantap! Berhasil konek ke Postgres via Go + Docker!")

	// 4. Buat Tabel Sederhana (DDL)
	/* query := `CREATE TABLE IF NOT EXISTS projects (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		tech TEXT NOT NULL
	)`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Gagal buat tabel:", err)
	}
	fmt.Println("Tabel 'projects' siap digunakan!") */

	// 5. INSERT DATA (Create)
	// Kita pakai placeholder $1, $2 (di MySQL pakai ?, di Postgres pakai $n)
	/* insertQuery := `INSERT INTO projects (name, tech) VALUES ($1, $2) RETURNING id`

	projectName := "Belajar Go Dasar"
	projectTech := "Golang & Postgres"
	var lastInsertID int */

	// QueryRow digunakan jika kita mengharapkan ada baris yang dikembalikan (RETURNING id)
	/* err = db.QueryRow(insertQuery, projectName, projectTech).Scan(&lastInsertID)
	if err != nil {
		log.Fatal("Gagal insert data:", err)
	}
	fmt.Printf("Data berhasil masuk! ID: %d\n", lastInsertID) */

	deleteQuery := `DELETE FROM projects`
	fmt.Println("test")

	// 6. SELECT DATA (Read)
	fmt.Println("\n--- Daftar Project di Database ---")
	rows, err := db.Query("SELECT id, name, tech FROM projects")
	if err != nil {
		log.Fatal("Gagal ambil data:", err)
	}
	defer rows.Close() // Jangan lupa ditutup agar tidak memory leak

	for rows.Next() {
		var id int
		var name, tech string

		// Scan memasukkan kolom dari DB ke variabel Go
		err := rows.Scan(&id, &name, &tech)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[%d] Nama: %-20s | Tech: %s\n", id, name, tech)
	}
}
