package infrastructures

import (
	"database/sql"
	"fmt"
	"log"
	"service-pattern-go/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbHandler struct {
	DB *gorm.DB
}

type OrmRow struct {
	Rows *sql.Rows
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

func (handler *DbHandler) Execute(statement string) {
	handler.DB.Exec(statement)
}

func (handler *DbHandler) Query(name string) ([]models.PlayerModel, error) {
	// Execute raw SQL query using GORM's Raw() method
	sql := "SELECT * FROM player_models WHERE name =";
	fmt.Println("%s '%s'", sql, name);
	result := handler.DB.Raw(sql + " ?", name)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	var players []models.PlayerModel
	result.Scan(&players) // GORM will map the result into the Player struct

	return players, nil
}
