package db

import (
	"demo/models"
	"demo/models/client"
	"github.com/dreamlu/gt"
)

func InitDB() {
	gt.NewDBTool().AutoMigrate(
		&client.Client{},
		&models.Service{},
		&models.Order{},
	)
}
