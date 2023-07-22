package initializers

import (
	"fmt"
	"os"

	"github.com/Lobaton2020/go-mvc-template/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase()  {
	var err error
	dns := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Could nod connect to database")
	}
	fmt.Println("Connected to the database successfully")
	DB = db
}

func SyncDB(){
	DB.AutoMigrate(&models.Post{})
}
