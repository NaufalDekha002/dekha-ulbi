package controller

import (
	"net/http"

	"github.com/NaufalDekha002/dekha-ulbi/config"
	inimodel "github.com/NaufalDekha002/penelitian-tugas/model"
	inimodul "github.com/NaufalDekha002/penelitian-tugas/module"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/NaufalDekha002/dekha-ulbi",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

// func Homepage(c *fiber.Ctx) error {
// 	ipaddr := musik.GetIPaddress()
// 	return c.JSON(ipaddr)
// }

func GetPresensilama(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

func GetPresensi(c *fiber.Ctx) error {
	nl := inimodul.GetAllPresensiFromPeneliti("Dekha Widana", config.Ulbimongoconn, "peneliti")
	return c.JSON(nl)
}

func GetPeneliti(c *fiber.Ctx) error {
	nl := inimodul.GetPenelitiFromPhoneNumber("083543242546", config.Ulbimongoconn, "phone_number")
	return c.JSON(nl)
}

func GetHasilTeliti(c *fiber.Ctx) error {
	nl := inimodul.GetHasilTelitiFromNama("Big Data Project", config.Ulbimongoconn, "hasil teliti")
	return c.JSON(nl)
}

func GetAll(c *fiber.Ctx) error {
	nl := inimodul.GetAllPresensiFromPeneliti("Dekha Widana", config.Ulbimongoconn, "peneliti")
	return c.JSON(nl)
}

func InsertStaff(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var peneliti inimodel.Peneliti
	if err := c.BodyParser(&peneliti); err != nil {
		return err
	}
	insertedID := inimodul.InsertPeneliti(db, "peneliti",
		peneliti.Nama,
		peneliti.Phone_number,
		peneliti.Jabatan,
		peneliti.JadwalPenelitian,
	)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Biodata berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertAbsen(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi inimodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return err
	}
	insertedID := inimodul.InsertPresensi(db, "absen",
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata,
	)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Biodata berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
