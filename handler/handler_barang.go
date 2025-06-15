package handler

import (
	"fmt"
	"strconv"
	"time"

	"project-app-inventaris-cli-rahmadhany/models"
	"project-app-inventaris-cli-rahmadhany/service"
	"project-app-inventaris-cli-rahmadhany/utils"
)

// BarangHandler manage CLI input for inventory barang features
type BarangHandler struct {
	Service service.BarangService
}

// NewBarangHandler return new handler for barang
func NewBarangHandler(s service.BarangService) *BarangHandler {
	return &BarangHandler{Service: s}
}

// List display all barang in table form
func (h *BarangHandler) List() {
	data, err := h.Service.GetAll()
	if err != nil {
		fmt.Println("Gagal ambil barang:", err)
		return
	}

	var rows [][]string
	for _, b := range data {
		days := int(time.Since(b.TanggalBeli).Hours() / 24)
		rows = append(rows, []string{
			strconv.Itoa(b.ID), b.Nama, b.TanggalBeli.Format("2006-01-02"),
			fmt.Sprintf("Rp%.0f", b.Harga), fmt.Sprintf("%d hari", days),
		})
	}
	utils.PrintTable([]string{"ID", "Nama", "Tgl Beli", "Harga", "Umur"}, rows)
}

// Add add barang based on CLI input
func (h *BarangHandler) Add() {
	nama := utils.ReadInput("Nama Barang: ")
	kategoriID, _ := strconv.Atoi(utils.ReadInput("ID Kategori: "))
	harga, _ := strconv.ParseFloat(utils.ReadInput("Harga: "), 64)
	tanggal := utils.ReadInput("Tanggal Beli (YYYY-MM-DD): ")
	tgl, _ := time.Parse("2006-01-02", tanggal)

	err := h.Service.Add(models.Barang{
		Nama: nama, KategoriID: kategoriID, Harga: harga, TanggalBeli: tgl,
	})
	if err != nil {
		fmt.Println("Gagal tambah barang:", err)
	} else {
		fmt.Println("✅ Barang ditambahkan.")
	}
}

// Detail display 1 barang detail based on ID
func (h *BarangHandler) Detail() {
	id, _ := strconv.Atoi(utils.ReadInput("ID Barang: "))
	b, err := h.Service.GetByID(id)
	if err != nil {
		fmt.Println("Gagal ambil detail:", err)
		return
	}
	fmt.Printf("ID: %d\nNama: %s\nKategori ID: %d\nHarga: Rp%.0f\nTanggal Beli: %s\n",
		b.ID, b.Nama, b.KategoriID, b.Harga, b.TanggalBeli.Format("2006-01-02"))
}

// Update update barang based on ID
func (h *BarangHandler) Update() {
	id, _ := strconv.Atoi(utils.ReadInput("ID Barang: "))
	nama := utils.ReadInput("Nama baru: ")
	kategoriID, _ := strconv.Atoi(utils.ReadInput("ID Kategori baru: "))
	harga, _ := strconv.ParseFloat(utils.ReadInput("Harga baru: "), 64)
	tanggal := utils.ReadInput("Tanggal Beli baru (YYYY-MM-DD): ")
	tgl, _ := time.Parse("2006-01-02", tanggal)

	err := h.Service.Update(models.Barang{
		ID: id, Nama: nama, KategoriID: kategoriID, Harga: harga, TanggalBeli: tgl,
	})
	if err != nil {
		fmt.Println("Gagal update:", err)
	} else {
		fmt.Println("✅ Barang berhasil diupdate.")
	}
}

// Delete delete barang based on ID
func (h *BarangHandler) Delete() {
	id, _ := strconv.Atoi(utils.ReadInput("ID Barang: "))
	err := h.Service.Delete(id)
	if err != nil {
		fmt.Println("Gagal hapus:", err)
	} else {
		fmt.Println("✅ Barang dihapus.")
	}
}

// BarangPerluDiganti display barang with the age of use > 100 hari
func (h *BarangHandler) BarangPerluDiganti() {
	data, err := h.Service.BarangPerluDiganti()
	if err != nil {
		fmt.Println("Gagal ambil data:", err)
		return
	}

	var rows [][]string
	for _, b := range data {
		age := int(time.Since(b.TanggalBeli).Hours() / 24)
		rows = append(rows, []string{
			strconv.Itoa(b.ID), b.Nama, b.TanggalBeli.Format("2006-01-02"),
			fmt.Sprintf("Rp%.0f", b.Harga), fmt.Sprintf("%d hari", age),
		})
	}
	utils.PrintTable([]string{"ID", "Nama", "Tgl Beli", "Harga", "Umur"}, rows)
}

// LaporanTotal calculate total investation value after depreciation
func (h *BarangHandler) LaporanTotal() {
	total, err := h.Service.TotalInvestasi()
	if err != nil {
		fmt.Println("Gagal hitung:", err)
		return
	}
	fmt.Printf("Total nilai investasi: Rp%.2f\n", total)
}

// LaporanPerBarang display 1 barang depreciation based on ID
func (h *BarangHandler) LaporanPerBarang() {
	id, _ := strconv.Atoi(utils.ReadInput("ID Barang: "))
	nilai, depresiasi, err := h.Service.DepresiasiBarang(id)
	if err != nil {
		fmt.Println("Gagal ambil laporan:", err)
		return
	}
	fmt.Printf("Nilai sekarang: Rp%.2f\nDepresiasi: Rp%.2f\n", nilai, depresiasi)
}

// SearchBarang search barang based on keyword
func (h *BarangHandler) SearchBarang(keyword string) {
	data, err := h.Service.Search(keyword)
	if err != nil {
		fmt.Println("Gagal cari barang:", err)
		return
	}

	var rows [][]string
	for _, b := range data {
		rows = append(rows, []string{
			strconv.Itoa(b.ID), b.Nama, b.TanggalBeli.Format("2006-01-02"),
			fmt.Sprintf("Rp%.0f", b.Harga),
		})
	}
	utils.PrintTable([]string{"ID", "Nama", "Tgl Beli", "Harga"}, rows)
}
