package handler

import (
	"fmt"
	"strconv"

	"project-app-inventaris-cli-rahmadhany/models"
	"project-app-inventaris-cli-rahmadhany/service"
	"project-app-inventaris-cli-rahmadhany/utils"
)

// KategoriHandler manage CLI input for inventory kategori features
type KategoriHandler struct {
	Service service.KategoriService
}

// NewKategoriHandler return new handler instance for kategori
func NewKategoriHandler(s service.KategoriService) *KategoriHandler {
	return &KategoriHandler{Service: s}
}

// List display all kategori in table form
func (h *KategoriHandler) List() {
	data, err := h.Service.GetAll()
	if err != nil {
		fmt.Println("Gagal ambil data:", err)
		return
	}

	var rows [][]string
	for _, k := range data {
		rows = append(rows, []string{
			strconv.Itoa(k.ID), k.Nama, k.Deskripsi,
		})
	}
	utils.PrintTable([]string{"ID", "Nama", "Deskripsi"}, rows)
}

// Add add kategori based on CLI input
func (h *KategoriHandler) Add() {
	nama := utils.ReadInput("Nama Kategori: ")
	deskripsi := utils.ReadInput("Deskripsi: ")
	err := h.Service.Add(models.Kategori{Nama: nama, Deskripsi: deskripsi})
	if err != nil {
		fmt.Println("Gagal tambah:", err)
	} else {
		fmt.Println("✅ Kategori ditambahkan.")
	}
}

// Detail display kategori detail based on ID
func (h *KategoriHandler) Detail() {
	id, _ := strconv.Atoi(utils.ReadInput("ID Kategori: "))
	data, err := h.Service.GetByID(id)
	if err != nil {
		fmt.Println("Gagal ambil detail:", err)
		return
	}
	fmt.Printf("ID: %d\nNama: %s\nDeskripsi: %s\n", data.ID, data.Nama, data.Deskripsi)
}

// Update update kategori based on ID
func (h *KategoriHandler) Update() {
	id, _ := strconv.Atoi(utils.ReadInput("ID Kategori: "))
	nama := utils.ReadInput("Nama baru: ")
	deskripsi := utils.ReadInput("Deskripsi baru: ")
	err := h.Service.Update(models.Kategori{ID: id, Nama: nama, Deskripsi: deskripsi})
	if err != nil {
		fmt.Println("Gagal update:", err)
	} else {
		fmt.Println("✅ Kategori berhasil diupdate.")
	}
}

// Delete delete kategori based on ID
func (h *KategoriHandler) Delete() {
	id, _ := strconv.Atoi(utils.ReadInput("ID Kategori yang akan dihapus: "))
	err := h.Service.Delete(id)
	if err != nil {
		fmt.Println("Gagal hapus:", err)
	} else {
		fmt.Println("✅ Kategori berhasil dihapus.")
	}
}
