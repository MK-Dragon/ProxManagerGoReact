package main

import (
	//"proxmanager/backend/controllers"
	"proxmanager/backend/models"

	"fmt"
	//"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "proxmanager/backend/docs" // for Swagger
)

func main() {
	fmt.Println("Proxmox Manager - GO Backend")

	r := gin.Default()

	r.GET("/vms", GetVMs)
	//r.GET("/lxcs", GetLXCs)
	//r.POST("/turn-off", TurnOff)
	//r.POST("/turn-on", TurnOn)

	// Swagger Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// GetVMs godoc
// @Summary Simple Get List of VMs
// @Tags VMs
// @Produce json
// @Success 200 {array} models.VM "List of Virtual Machines"
// @Router /vms [get]
func GetVMs(c *gin.Context) {
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

	// Return the data
	c.JSON(200, data)
}
