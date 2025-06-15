package service

import (
	"fmt"
	"project-app-inventaris-cli-rahmadhany/models"
	"project-app-inventaris-cli-rahmadhany/repository"
)

// KategoriService interface
type KategoriService interface {
	GetAll() ([]models.Kategori, error)
	GetByID(id int) (*models.Kategori, error)
	Add(k models.Kategori) error
	Update(k models.Kategori) error
	Delete(id int) error
}

// kategoriService implementation
type kategoriService struct {
	repo repository.KategoriRepository
}

// NewKategoriService create instance for kategori service
func NewKategoriService(r repository.KategoriRepository) KategoriService {
	return &kategoriService{repo: r}
}

// GetAll retrieve all kategori data
func (s *kategoriService) GetAll() ([]models.Kategori, error) {
	return s.repo.GetAll()
}

// GetByID retrieve 1 kategori data based on ID
func (s *kategoriService) GetByID(id int) (*models.Kategori, error) {
	return s.repo.GetByID(id)
}

// Add save new kategori data with validation, name cannot be empty
func (s *kategoriService) Add(k models.Kategori) error {
	if k.Nama == "" {
		return fmt.Errorf("nama tidak boleh kosong")
	}
	return s.repo.Add(k)
}

// Update update kategori data
func (s *kategoriService) Update(k models.Kategori) error {
	return s.repo.Update(k)
}

// Delete delete kategori based on ID
func (s *kategoriService) Delete(id int) error {
	return s.repo.Delete(id)
}
