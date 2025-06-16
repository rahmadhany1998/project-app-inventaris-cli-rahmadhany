package service

import (
	"errors"
	"project-app-inventaris-cli-rahmadhany/models"
	"testing"
	"time"
)

type fakeBarangRepo struct{}

func (f *fakeBarangRepo) GetAll() ([]models.Barang, error) {
	return []models.Barang{{ID: 1, Nama: "Laptop", Harga: 1000000, TanggalBeli: time.Now().AddDate(0, 0, -200)}}, nil
}
func (f *fakeBarangRepo) GetByID(id int) (*models.Barang, error) {
	return &models.Barang{ID: id, Nama: "Laptop", Harga: 1000000, TanggalBeli: time.Now().AddDate(0, 0, -300)}, nil
}
func (f *fakeBarangRepo) Add(b models.Barang) error {
	if b.Nama == "" {
		return errors.New("nama kosong")
	}
	return nil
}
func (f *fakeBarangRepo) Update(b models.Barang) error { return nil }
func (f *fakeBarangRepo) Delete(id int) error          { return nil }
func (f *fakeBarangRepo) Search(k string) ([]models.Barang, error) {
	return []models.Barang{{ID: 1, Nama: "Cari", Harga: 123, TanggalBeli: time.Now()}}, nil
}

func TestBarangService(t *testing.T) {
	svc := NewBarangService(&fakeBarangRepo{})

	list, _ := svc.GetAll()
	if len(list) == 0 {
		t.Error("Barang kosong")
	}

	err := svc.Add(models.Barang{Nama: ""})
	if err == nil {
		t.Error("Seharusnya error karena nama kosong")
	}

	_, _, err = svc.DepresiasiBarang(1)
	if err != nil {
		t.Error("Gagal hitung depresiasi")
	}

	ganti, _ := svc.BarangPerluDiganti()
	if len(ganti) == 0 {
		t.Error("Seharusnya ada barang yang perlu diganti")
	}
}
