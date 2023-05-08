package table

type DataMahasiswa struct {
	Id     uint   `json:"id"`
	Nama   string `json:"nama"`
	Nim    string `json:"nim"`
	Wa     string `json:"wa"`
	Alamat string `json:"alamat"`
	// Created_at string
	// Updated_at string
	// Created_by string
	// Updated_by string
	// Deleted_at string
	// Deleted_by string
}

type RequestDataMahasiswa struct {
	Nama   string
	Nim    string
	Wa     string
	Alamat string
}
