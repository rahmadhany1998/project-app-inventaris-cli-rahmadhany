package service

import (
	"fmt"
	"math"
	"time"

	"project-app-inventaris-cli-rahmadhany/models"
	"project-app-inventaris-cli-rahmadhany/repository"
)

// BarangService interface
type BarangService interface {
	GetAll() ([]models.Barang, error)
	GetByID(id int) (*models.Barang, error)
	Add(b models.Barang) error
	Update(b models.Barang) error
	Delete(id int) error
	Search(keyword string) ([]models.Barang, error)
	BarangPerluDiganti() ([]models.Barang, error)
	TotalInvestasi() (float64, error)
	DepresiasiBarang(id int) (float64, float64, error)
}

// barangService implementation
type barangService struct {
	repo repository.BarangRepository
}

// NewBarangService create instance for barang service
func NewBarangService(r repository.BarangRepository) BarangService {
	return &barangService{repo: r}
}

// GetAll retrieve all barang data
func (s *barangService) GetAll() ([]models.Barang, error) {
	return s.repo.GetAll()
}

// GetByID retrieve 1 barang data based on ID
func (s *barangService) GetByID(id int) (*models.Barang, error) {
	return s.repo.GetByID(id)
}

// Add save new barang data with validation, name cannot be empty
func (s *barangService) Add(b models.Barang) error {
	if b.Nama == "" {
		return fmt.Errorf("nama tidak boleh kosong")
	}
	return s.repo.Add(b)
}

// Update update barang data
func (s *barangService) Update(b models.Barang) error {
	return s.repo.Update(b)
}

// Delete delete barang based on ID
func (s *barangService) Delete(id int) error {
	return s.repo.Delete(id)
}

// Search search barang based on keyword (name)
func (s *barangService) Search(keyword string) ([]models.Barang, error) {
	return s.repo.Search(keyword)
}

// BarangPerluDiganti return barang with the age of use > 100 days
func (s *barangService) BarangPerluDiganti() ([]models.Barang, error) {
	barangs, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []models.Barang
	now := time.Now()
	for _, b := range barangs {
		if int(now.Sub(b.TanggalBeli).Hours()/24) > 100 {
			result = append(result, b)
		}
	}
	return result, nil
}

// TotalInvestasi calculate the total value of barang after depreciation
func (s *barangService) TotalInvestasi() (float64, error) {
	barangs, err := s.repo.GetAll()
	if err != nil {
		return 0, err
	}

	now := time.Now()
	total := 0.0
	for _, b := range barangs {
		age := now.Sub(b.TanggalBeli).Hours() / 24 / 365
		value := b.Harga * math.Pow(0.8, age) // depresiasi saldo menurun 20%
		total += value
	}
	return total, nil
}

// DepresiasiBarang calculate the present value and depreciation of 1 barang based on ID
func (s *barangService) DepresiasiBarang(id int) (float64, float64, error) {
	b, err := s.repo.GetByID(id)
	if err != nil {
		return 0, 0, err
	}

	now := time.Now()
	age := now.Sub(b.TanggalBeli).Hours() / 24 / 365
	nilai := b.Harga * math.Pow(0.8, age)
	depresiasi := b.Harga - nilai
	return nilai, depresiasi, nil
}
