package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/kulkasdingin/goooo/app/complements"
	"github.com/kulkasdingin/goooo/app/models"
	"github.com/kulkasdingin/goooo/app/services"
)

func IndexBlog(w http.ResponseWriter, r *http.Request) {
	//
}

func StoreBlog(w http.ResponseWriter, r *http.Request) { // TODO: Bikin test kalo requestnya jelek atau salah
	reqBody, _ := ioutil.ReadAll(r.Body)

	var rawBlog models.Blog

	json.Unmarshal(reqBody, &rawBlog)

	blog := services.CreateBlogFromRequest(rawBlog)

	complements.RespondWithJSON(w, http.StatusCreated, blog)

}

func ShowBlog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		complements.RespondWithError(w, http.StatusBadRequest, "Invalid blog ID")
		return
	}

	blog, err := services.GetBlogById(id)

	if err != nil {
		complements.RespondWithError(w, http.StatusNotFound, "Blog not found")
		return
	}

	complements.RespondWithJSON(w, http.StatusOK, blog)
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	old_id, err := strconv.Atoi(params["id"])

	if err != nil {
		complements.RespondWithError(w, http.StatusBadRequest, "Invalid blog ID")
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newBlog models.Blog

	json.Unmarshal(reqBody, &newBlog)

	blog, err := services.UpdateBlogFromRequest(newBlog, old_id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		complements.RespondWithError(w, http.StatusNotFound, "Blog not found")
		return
	}

	complements.RespondWithJSON(w, http.StatusOK, blog)
}

func DestroyBlog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		complements.RespondWithError(w, http.StatusBadRequest, "Invalid blog ID")
		return
	}

	err = services.DeleteBlogFromRequest(id)

	if err != nil {
		complements.RespondWithError(w, http.StatusNotFound, "Blog not found")
		return
	}

	complements.RespondWithJSON(w, http.StatusOK, `{"Message": "Blog successfully deleted"}`)
}
