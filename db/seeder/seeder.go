package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kulkasdingin/goooo/app/models"
	"golang.org/x/crypto/bcrypt"
)

func Seeder() {
	var pwd = []byte("farizfariz")
	hash, errr := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if errr != nil {
		fmt.Println(errr)
	}
	// pwdhashed := string(hash[:])

	var (
		users = []models.User{
			{
				Username: "fariz",
				Email:    "fariz@frz.codes",
				Password: hash,
			},
		}

		blogs = []models.Blog{
			{
				Title:       "How to be Handsome 101",
				HeaderImage: "https://media-exp1.licdn.com/dms/image/C5103AQEUwrQpsiZ4ag/profile-displayphoto-shrink_200_200/0/1551580769854?e=1630540800&v=beta&t=taFYtsLntScNhnFCKMw7qQ1t82lw0HYRttm3uTwmgAs",
				Content:     "Fariz sangat ganteng",
				UserID:      1,
			},
			{
				Title:       "How to Mencintai dan Dicintai",
				HeaderImage: "https://media-exp1.licdn.com/dms/image/C5103AQEUwrQpsiZ4ag/profile-displayphoto-shrink_200_200/0/1551580769854?e=1630540800&v=beta&t=taFYtsLntScNhnFCKMw7qQ1t82lw0HYRttm3uTwmgAs",
				Content:     "Fariz tetap sangat ganteng",
				UserID:      1,
			},
		}
	)

	var db *gorm.DB

	var err error

	godotenv.Load("../.env")

	var dbconf = "host=" + os.Getenv("HOST") + " port=" + os.Getenv("PORT") + " user=" + os.Getenv("USER") + " dbname=" + os.Getenv("DBNAME") + " sslmode=" + os.Getenv("SSLMODE") + " password=" + os.Getenv("DBPASS") + ""

	db, err = gorm.Open("postgres", dbconf)

	if err != nil {

		panic(err)

	}

	defer db.Close()

	for index := range users {

		db.Create(&users[index])

	}

	for index := range blogs {

		db.Create(&blogs[index])

	}
}
