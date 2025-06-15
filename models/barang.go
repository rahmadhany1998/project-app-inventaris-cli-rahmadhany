package models

import "time"

type Barang struct {
	ID          int
	Nama        string
	KategoriID  int
	Harga       float64
	TanggalBeli time.Time
}
