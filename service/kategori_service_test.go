package service

import (
	"errors"
	"project-app-inventaris-cli-rahmadhany/models"
	"testing"
)

type fakeKategoriRepo struct{}

func (f *fakeKategoriRepo) GetAll() ([]models.Kategori, error) {
	return []models.Kategori{{ID: 1, Nama: "Elektronik", Deskripsi: "Alat elektronik"}}, nil
}
func (f *fakeKategoriRepo) GetByID(id int) (*models.Kategori, error) {
	return &models.Kategori{ID: id, Nama: "Dummy", Deskripsi: "Desc"}, nil
}
func (f *fakeKategoriRepo) Add(k models.Kategori) error {
	if k.Nama == "" {
		return errors.New("nama kosong")
	}
	return nil
}
func (f *fakeKategoriRepo) Update(k models.Kategori) error { return nil }
func (f *fakeKategoriRepo) Delete(id int) error            { return nil }

func TestKategoriService(t *testing.T) {
	svc := NewKategoriService(&fakeKategoriRepo{})

	data, err := svc.GetAll()
	if err != nil || len(data) == 0 {
		t.Error("Gagal ambil data kategori")
	}

	if err := svc.Add(models.Kategori{Nama: ""}); err == nil {
		t.Error("Seharusnya gagal karena nama kosong")
	}

	detail, err := svc.GetByID(1)
	if err != nil || detail.ID != 1 {
		t.Error("Detail gagal")
	}
}
