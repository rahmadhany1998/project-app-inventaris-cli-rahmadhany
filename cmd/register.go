package cmd

import "project-app-inventaris-cli-rahmadhany/handler"

var kategoriHandler *handler.KategoriHandler
var barangHandler *handler.BarangHandler

func RegisterHandlers(kh *handler.KategoriHandler, bh *handler.BarangHandler) {
	kategoriHandler = kh
	barangHandler = bh
}
