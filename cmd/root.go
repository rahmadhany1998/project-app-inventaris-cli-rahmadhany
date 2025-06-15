package cmd

import (
	"fmt"
	"project-app-inventaris-cli-rahmadhany/utils"

	"github.com/spf13/cobra"
)

// Execute run root command CLI
func Execute() error {
	return rootCmd.Execute()
}

// rootCmd is am main command CLI
var rootCmd = &cobra.Command{
	Use:   "inventaris",
	Short: "Aplikasi Sistem Inventaris Kantor CLI",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			utils.ClearScreen()
			fmt.Println("========== SISTEM INVENTARIS KANTOR ==========")
			fmt.Println("1. Kelola Kategori Barang")
			fmt.Println("2. Kelola Barang Inventaris")
			fmt.Println("3. Laporan & Pengecekan")
			fmt.Println("4. Cari Barang")
			fmt.Println("0. Keluar")
			fmt.Println("==============================================")
			choice := utils.ReadInput("Pilih menu: ")

			switch choice {
			case "1":
				handleKategoriMenu()
			case "2":
				handleBarangMenu()
			case "3":
				handleLaporanMenu()
			case "4":
				keyword := utils.ReadInput("Masukkan kata kunci pencarian: ")
				barangHandler.SearchBarang(keyword)
			case "0":
				fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
				return
			default:
				fmt.Println("Pilihan tidak valid.")
			}
			utils.PromptEnter()
		}
	},
}

// handleKategoriMenu handle sub-menu kategori
func handleKategoriMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("--- Menu Kategori Barang ---")
		fmt.Println("1. Lihat semua kategori")
		fmt.Println("2. Tambah kategori")
		fmt.Println("3. Lihat detail kategori")
		fmt.Println("4. Update kategori")
		fmt.Println("5. Hapus kategori")
		fmt.Println("0. Kembali")
		choice := utils.ReadInput("Pilih menu kategori: ")

		switch choice {
		case "1":
			kategoriHandler.List()
		case "2":
			kategoriHandler.Add()
		case "3":
			kategoriHandler.Detail()
		case "4":
			kategoriHandler.Update()
		case "5":
			kategoriHandler.Delete()
		case "0":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		utils.PromptEnter()
	}
}

// handleBarangMenu handle sub-menu barang
func handleBarangMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("--- Menu Barang Inventaris ---")
		fmt.Println("1. Lihat semua barang")
		fmt.Println("2. Tambah barang")
		fmt.Println("3. Lihat detail barang")
		fmt.Println("4. Update barang")
		fmt.Println("5. Hapus barang")
		fmt.Println("0. Kembali")
		choice := utils.ReadInput("Pilih menu barang: ")

		switch choice {
		case "1":
			barangHandler.List()
		case "2":
			barangHandler.Add()
		case "3":
			barangHandler.Detail()
		case "4":
			barangHandler.Update()
		case "5":
			barangHandler.Delete()
		case "0":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		utils.PromptEnter()
	}
}

// handleLaporanMenu handle sub-menu report
func handleLaporanMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("--- Menu Laporan & Pengecekan ---")
		fmt.Println("1. Barang yang perlu diganti (>100 hari)")
		fmt.Println("2. Total nilai investasi & depresiasi semua barang")
		fmt.Println("3. Laporan investasi & depresiasi berdasarkan ID")
		fmt.Println("0. Kembali")
		choice := utils.ReadInput("Pilih menu laporan: ")

		switch choice {
		case "1":
			barangHandler.BarangPerluDiganti()
		case "2":
			barangHandler.LaporanTotal()
		case "3":
			barangHandler.LaporanPerBarang()
		case "0":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		utils.PromptEnter()
	}
}
