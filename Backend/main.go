package main

import (
	//"proxmanager/backend/controllers"
	"proxmanager/backend/models"

	"fmt"
	//"strings"

	//"github.com/gin-gonic/gin"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // Fiber-specific swagger
	_ "proxmanager/backend/docs" // for Swagger
)

func main() {
	fmt.Println("Proxmox Manager - GO Backend (Fiber)")

	app := fiber.New()

	// Routes
	app.Get("/vms", GetVMs)
	//r.GET("/lxcs", GetLXCs)
	//r.POST("/turn-off", TurnOff)
	//r.POST("/turn-on", TurnOn)

	// Swagger Route - Access at /swagger/index.html
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":8080")
}

// GetVMs godoc
// @Summary Simple Get List of VMs
// @Tags VMs
// @Produce json
// @Success 200 {array} models.VM "List of Virtual Machines"
// @Router /vms [get]
func GetVMs(c *fiber.Ctx) error {
	// test data
	data := []models.VM{
		{
			Machine: models.Machine{
				Id:     101,
				Name:   "Ubuntu-Server-24",
				Status: "running",
				Cores:  4,
				Ram:    8192,
			},
			Usbs: []string{"USB-Controller-0"},
			Pcie: []string{"RTX-4090-Passthrough"},
		},
		{
			Machine: models.Machine{
				Id:     102,
				Name:   "Windows-11-Gaming",
				Status: "stopped",
				Cores:  8,
				Ram:    16384,
			},
			Usbs: []string{},
			Pcie: []string{},
		},
	}

	// In Fiber, we return the result of c.JSON
	return c.Status(fiber.StatusOK).JSON(data)
}
