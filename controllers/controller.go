package controllers

import (
	"am_tryout/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var tryout []models.Tryout

	models.DB.Db.Find(&tryout)

	return c.Status(fiber.StatusOK).JSON(tryout)

}

func Create(c *fiber.Ctx) error {

	user := new(models.Tryout)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var tryout models.Tryout

	result := models.DB.Db.Where("id = ?", id).First(&tryout)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": result.Error.Error(),
		})
	}

	return c.JSON(tryout)
}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")
	var updatedData models.Tryout

	// Parse the body to get the updated user data
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if models.DB.Db.Where("id = ?", id).Updates(&updatedData).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id tidak ditemukan.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate.",
	})

}

/* func Delete(c *fiber.Ctx) error {

	username := c.Params("username")

	// Delete the user record
	result := models.DB.Db.Where("username = ?", username).Delete(&models.Tryout{})
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).SendString("User deleted successfully")
} */
