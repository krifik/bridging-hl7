package model

type Rujukan struct {
	IdDokter    string `json:"id_dokter"`
	NamaDokter  string `json:"nama_dokter"`
	IdWard      string `json:"id_ward"`
	Ward        string `json:"ward"`
	IdFasilitas string `json:"id_fasilitas"`
	Fasilitas   string `json:"fasilitas"`
}
