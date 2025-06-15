package repository

import (
	"database/sql"
	"project-app-inventaris-cli-rahmadhany/models"
	"strings"
)

// BarangRepository interface
type BarangRepository interface {
	GetAll() ([]models.Barang, error)
	GetByID(id int) (*models.Barang, error)
	Add(b models.Barang) error
	Update(b models.Barang) error
	Delete(id int) error
	Search(keyword string) ([]models.Barang, error)
}

type barangRepo struct {
	db *sql.DB
}

// NewBarangRepository create instance for barang repository
func NewBarangRepository(db *sql.DB) BarangRepository {
	return &barangRepo{db}
}

// GetAll retrieve all barang data
func (r *barangRepo) GetAll() ([]models.Barang, error) {
	rows, err := r.db.Query("SELECT id, nama, kategori_id, harga, tanggal_beli FROM barang")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Barang
	for rows.Next() {
		var b models.Barang
		if err := rows.Scan(&b.ID, &b.Nama, &b.KategoriID, &b.Harga, &b.TanggalBeli); err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}

// GetByID retrieve 1 barang data based on ID
func (r *barangRepo) GetByID(id int) (*models.Barang, error) {
	row := r.db.QueryRow("SELECT id, nama, kategori_id, harga, tanggal_beli FROM barang WHERE id=$1", id)
	var b models.Barang
	if err := row.Scan(&b.ID, &b.Nama, &b.KategoriID, &b.Harga, &b.TanggalBeli); err != nil {
		return nil, err
	}
	return &b, nil
}

// Add save new barang data to database
func (r *barangRepo) Add(b models.Barang) error {
	_, err := r.db.Exec("INSERT INTO barang (nama, kategori_id, harga, tanggal_beli) VALUES ($1, $2, $3, $4)",
		b.Nama, b.KategoriID, b.Harga, b.TanggalBeli)
	return err
}

// Update update barang data
func (r *barangRepo) Update(b models.Barang) error {
	_, err := r.db.Exec("UPDATE barang SET nama=$1, kategori_id=$2, harga=$3, tanggal_beli=$4 WHERE id=$5",
		b.Nama, b.KategoriID, b.Harga, b.TanggalBeli, b.ID)
	return err
}

// Delete delete barang based on ID
func (r *barangRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM barang WHERE id=$1", id)
	return err
}

// Search search barang based on keyword (name)
func (r *barangRepo) Search(keyword string) ([]models.Barang, error) {
	keyword = "%" + strings.ToLower(keyword) + "%"
	rows, err := r.db.Query("SELECT id, nama, kategori_id, harga, tanggal_beli FROM barang WHERE LOWER(nama) LIKE $1", keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Barang
	for rows.Next() {
		var b models.Barang
		if err := rows.Scan(&b.ID, &b.Nama, &b.KategoriID, &b.Harga, &b.TanggalBeli); err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}
