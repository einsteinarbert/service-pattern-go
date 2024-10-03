package infrastructures

import (
	"log"
	"service-pattern-go/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbHandler struct {
	DB *gorm.DB
}

func NewDbHandler(dbPath string) *DbHandler {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Automatically migrate your schema
	db.AutoMigrate(&models.PlayerModel{})

	return &DbHandler{DB: db}
}
