package handler

import (
	"belajar-go-dasar/database"
	"belajar-go-dasar/table"
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx) error {

	db := database.DB

	var (
		request = new(table.RequestDataMahasiswa)
		dm      table.DataMahasiswa
	)

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	dm.Nama = request.Nama
	dm.Nim = request.Nim
	dm.Wa = request.Wa
	dm.Alamat = request.Alamat

	if err := db.Create(&dm).Error; err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Succes Create Data",
		"data":    request,
	})

}

func Read(c *fiber.Ctx) error {

	db := database.DB

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	var (
		total_rows int64
		data       []table.DataMahasiswa
	)

	nama := c.Query("nama")

	query := db.Table("data_mahasiswas").
		Where("nama LIKE ?", "%"+nama+"%").
		Order("id desc")

	query.Count(&total_rows)
	query.Offset(offset).Limit(limit)
	query.Find(&data)

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Data Mahasiswa",
		"data":    data,
	})

}

func Update(c *fiber.Ctx) error {

	db := database.DB

	id := c.Query("id")

	var (
		request = new(table.RequestDataMahasiswa)
	)

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	update := table.DataMahasiswa{
		Nama:   request.Nama,
		Nim:    request.Nim,
		Wa:     request.Wa,
		Alamat: request.Alamat,
	}

	if err2 := db.Where("id = ?", id).Updates(&update); err2 != nil {
		log.Println(err2)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err2,
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Succes Update Data",
		"data":    request,
	})
}

func Delete(c *fiber.Ctx) error {

	db := database.DB

	id := c.Query("id")

	var data table.DataMahasiswa

	// delete
	deleteQuery := db.Model(&data).Where("id = ?", id).Delete(&data)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "opps!, record not found",
		})
	}

	db.Model(&data).Where("id = ?", id).Delete(&data)

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Succes Delete Data",
	})

}
