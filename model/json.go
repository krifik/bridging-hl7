package model

type Json struct {
	NoPendaftaran string    `json:"no_pendaftaran"`
	NoRm          string    `json:"no_rm"`
	NoOrder       string    `json:"no_order"`
	NamaPasien    string    `json:"nama_pasien"`
	TempatLahir   string    `json:"templat_lahir"`
	TglLahir      string    `json:"tgl_lahir"`
	Jk            string    `json:"jenis_kelamin"`
	Alamat        string    `json:"alamat"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Nik           string    `json:"nik"`
	IdJenisPasien string    `json:"id_jenis_pasien"`
	JenisPasien   string    `json:"jenis_pasien"`
	IdPenjamin    string    `json:"id_penjamin"`
	Penjamin      string    `json:"penjamin"`
	RujukanAsal   string    `json:"rujukan_asal"`
	DetailRujukan []Rujukan `json:"detail_rujukan"`
	Cito          string    `json:"cito"`
	Diagnosa      string    `json:"diagnose"`
	Icd10         []Icd10   `json:"icd10"`
	Order         []Order   `json:"order"`
}
