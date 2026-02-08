package model

type ReportRes struct {
	TotalRevenue   int            `json:"total_revenue"`
	TotalTransaksi int            `json:"total_transaksi"`
	ProdukTerlaris ProdukTerlaris `json:"produk_terlaris"`
}

type ProdukTerlaris struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

type ReportReq struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}
