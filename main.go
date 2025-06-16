package main

import (
	"log"

	"project-app-inventaris-cli-rahmadhany/cmd"
	db "project-app-inventaris-cli-rahmadhany/database"
	"project-app-inventaris-cli-rahmadhany/handler"
	"project-app-inventaris-cli-rahmadhany/repository"
	"project-app-inventaris-cli-rahmadhany/service"
)

func main() {
	// Initialize PostgreSQL database connection
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}

	// Dependency injection for kategori modul
	kategoriRepo := repository.NewKategoriRepository(conn)
	kategoriService := service.NewKategoriService(kategoriRepo)
	kategoriHandler := handler.NewKategoriHandler(kategoriService)

	// Dependency injection for barang ategori modul
	barangRepo := repository.NewBarangRepository(conn)
	barangService := service.NewBarangService(barangRepo)
	barangHandler := handler.NewBarangHandler(barangService)

	// Register Handler
	cmd.RegisterHandlers(kategoriHandler, barangHandler)

	// Execute CLI
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Command execution error: %v", err)
	}
}
