package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"github.com/kulkasdingin/goooo/app/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

var err error

func Migration() {

	godotenv.Load("../.env")

	var dbconf = "host=" + os.Getenv("HOST") + " port=" + os.Getenv("PORT") + " user=" + os.Getenv("USER") + " dbname=" + os.Getenv("DBNAME") + " sslmode=" + os.Getenv("SSLMODE") + " password=" + os.Getenv("DBPASS") + ""

	db, err = gorm.Open("postgres", dbconf)
	if err != nil {

		panic(err)

	}

	defer db.Close()

	up(db)

	fmt.Println("Migration success")
}

func up(db *gorm.DB) {
	upBlog(db)
	upUser(db)
}

func upBlog(db *gorm.DB) {
	db.AutoMigrate(&models.Blog{})
}

func upUser(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
