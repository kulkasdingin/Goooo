package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kulkasdingin/goooo/app/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var a App

func TestMain(m *testing.M) {
	a.Init()

	code := m.Run()

	os.Exit(code)
}

func TestGetNonExistentBlog(t *testing.T) {
	clearTable("blogs")

	req, _ := http.NewRequest("GET", "/api/blog/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Blog not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Blog not found'. Got '%s'", m["error"])
	}
}

func TestCreateBlog(t *testing.T) {
	clearTable("blogs")

	var jsonStr = []byte(`{"Title":"Test Title", "Content":"Test Content", "HeaderImage":"Test HeaderImage", "UserID":1}`)
	req, _ := http.NewRequest("POST", "/api/blog", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["Title"] != "Test Title" {
		t.Errorf("Expected Blog 'Title' to be 'Test Title'. Got '%v'", m["Title"])
	}
	if m["Content"] != "Test Content" {
		t.Errorf("Expected Blog 'Content' to be 'Test Content'. Got '%v'", m["Content"])
	}
	if m["HeaderImage"] != "Test HeaderImage" {
		t.Errorf("Expected Blog 'HeaderImage' to be 'Test HeaderImage'. Got '%v'", m["HeaderImage"])
	}
	if m["UserID"] != 1.0 {
		t.Errorf("Expected Blog 'UserID' to be '1'. Got '%v'", m["UserID"])
	}

	if m["ID"] != 1.0 {
		t.Errorf("Expected Blog 'ID' to be '1'. Got '%v'", m["ID"])
	}
}

func TestFetchProduct(t *testing.T) {
	clearTable("blogs")

	addATestBlog()

	req, _ := http.NewRequest("GET", "/api/blog/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateProduct(t *testing.T) {
	clearTable("blogs")

	addATestBlog()

	req, _ := http.NewRequest("GET", "/api/blog/1", nil)
	response := executeRequest(req)
	var originalBlog map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalBlog)

	var jsonStr = []byte(`"Title":"Test Updated Title", "HeaderImage":"Test Updated HeaderImage", "Content":"Test Updated Content", "UserID":1}`)

	req, _ = http.NewRequest("PUT", "/api/blog/100", bytes.NewBuffer(jsonStr))
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

	req, _ = http.NewRequest("PUT", "/api/blog/1", bytes.NewBuffer(jsonStr))
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}

	json.Unmarshal(response.Body.Bytes(), &m)

	if m["ID"] != originalBlog["ID"] {
		t.Errorf("Expected the ID to remain the same (%v). Got %v", originalBlog["ID"], m["ID"])
	}

	if m["Title"] == originalBlog["Title"] {
		t.Errorf("Expected the Title to change from '%v' to Test Updated Title. Got '%v'", originalBlog["Title"], m["Title"])
	}

	if m["HeaderImage"] == originalBlog["HeaderImage"] {
		t.Errorf("Expected the HeaderImage to change from '%v' to Test Updated HeaderImage. Got '%v'", originalBlog["HeaderImage"], m["HeaderImage"])
	}

	if m["Content"] == originalBlog["Content"] {
		t.Errorf("Expected the Content to change from '%v' to Test Updated Content. Got '%v'", originalBlog["Content"], m["Content"])
	}

	if m["UserID"] != originalBlog["UserID"] {
		t.Errorf("Expected the UserID to remain the same (%v). Got %v", originalBlog["UserID"], m["UserID"])
	}
}

func TestDeleteBlog(t *testing.T) {
	clearTable("blogs")

	addATestBlog()

	req, _ := http.NewRequest("GET", "/api/blog/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/api/blog/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/api/blog/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

}

func clearTable(table string) {
	a.DB.Exec("DELETE FROM blogs")
	a.DB.Exec("ALTER SEQUENCE blogs_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func addATestBlog() {
	blogs := []models.Blog{
		{
			Title:       "Test Title",
			HeaderImage: "Test HeaderImage",
			Content:     "Test Content",
			UserID:      1,
		},
	}

	for index := range blogs {

		a.DB.Create(&blogs[index])

	}
}
