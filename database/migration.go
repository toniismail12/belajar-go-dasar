package database

import (
	"belajar-go-dasar/table"
)

func Migration() {

	DB.AutoMigrate(

		&table.DataMahasiswa{},

	)

}
