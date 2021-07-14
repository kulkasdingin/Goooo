package routes

import (
	"github.com/kulkasdingin/goooo/app/controllers"
)

func HandleApiRequests() {
	Router.HandleFunc(ra+"/blog/{id}", controllers.ShowBlog).Methods("GET")

	Router.HandleFunc(ra+"/blog", controllers.StoreBlog).Methods("POST")

	Router.HandleFunc(ra+"/blog/{id}", controllers.UpdateBlog).Methods("PUT")

	Router.HandleFunc(ra+"/blog/{id}", controllers.DestroyBlog).Methods("DELETE")
}
