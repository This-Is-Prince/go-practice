package lead

import (
	"github.com/This-Is-Prince/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DB
	var leads []Lead

	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var lead Lead

	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}

	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DB
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var lead Lead

	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}

	db.Delete(&lead, id)
	c.Send("Lead successfully deleted")
}
