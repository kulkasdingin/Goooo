package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/kulkasdingin/goooo/app/repositories"
	"github.com/kulkasdingin/goooo/routes"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Init() {
	godotenv.Load("../.env") // keluar folder

	var dbconf = "host=" + os.Getenv("HOST") + " port=" + os.Getenv("PORT") + " user=" + os.Getenv("USER") + " dbname=" + os.Getenv("DBNAME") + " sslmode=" + os.Getenv("SSLMODE") + " password=" + os.Getenv("DBPASS") + ""

	var err error
	a.DB, err = gorm.Open("postgres", dbconf)

	if err != nil {

		panic(err)

	}

	repositories.DB = a.DB

	// defer a.DB.Close() // gabisa dilakuin soalnya sekarang initnya dipisah dari main, pas fungsi init kelar sqlnya ketutup // TODO: Taro di main() tp gatau bener atau engga

	a.Router = mux.NewRouter()

	routes.Router = a.Router

	routes.Main()
}

func (a *App) Run() {
	handler := cors.Default().Handler(a.Router)

	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
