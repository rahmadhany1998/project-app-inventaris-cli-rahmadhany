package repository

import (
	"database/sql"
	"project-app-inventaris-cli-rahmadhany/models"
)

// KategoriRepository interface
type KategoriRepository interface {
	GetAll() ([]models.Kategori, error)
	GetByID(id int) (*models.Kategori, error)
	Add(k models.Kategori) error
	Update(k models.Kategori) error
	Delete(id int) error
}

type kategoriRepo struct {
	db *sql.DB
}

// NewKategoriRepository create instance for kategori repository
func NewKategoriRepository(db *sql.DB) KategoriRepository {
	return &kategoriRepo{db}
}

// GetAll retrieve all kategori data
func (r *kategoriRepo) GetAll() ([]models.Kategori, error) {
	rows, err := r.db.Query("SELECT id, nama, deskripsi FROM kategori")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Kategori
	for rows.Next() {
		var k models.Kategori
		if err := rows.Scan(&k.ID, &k.Nama, &k.Deskripsi); err != nil {
			return nil, err
		}
		result = append(result, k)
	}
	return result, nil
}

// GetByID retrieve 1 kategori data based on ID
func (r *kategoriRepo) GetByID(id int) (*models.Kategori, error) {
	row := r.db.QueryRow("SELECT id, nama, deskripsi FROM kategori WHERE id = $1", id)
	var k models.Kategori
	if err := row.Scan(&k.ID, &k.Nama, &k.Deskripsi); err != nil {
		return nil, err
	}
	return &k, nil
}

// Add save new kategori data to database
func (r *kategoriRepo) Add(k models.Kategori) error {
	_, err := r.db.Exec("INSERT INTO kategori (nama, deskripsi) VALUES ($1, $2)", k.Nama, k.Deskripsi)
	return err
}

// Update update kategori data
func (r *kategoriRepo) Update(k models.Kategori) error {
	_, err := r.db.Exec("UPDATE kategori SET nama=$1, deskripsi=$2 WHERE id=$3", k.Nama, k.Deskripsi, k.ID)
	return err
}

// Delete delete kategori based on ID
func (r *kategoriRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM kategori WHERE id = $1", id)
	return err
}
