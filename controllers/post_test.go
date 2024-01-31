package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Rahul06x1/go_crud/models"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func savePost(t *testing.T, entry models.Post, method string, id int) (*http.Request, *httptest.ResponseRecorder) {
	requestBody, err := json.Marshal(entry)
	if err != nil {
		t.Fatal(err)
	}
	url :="/api/posts"
	if id != 0 {
		url =fmt.Sprintf("/api/posts/%d", id)
	}
	req := httptest.NewRequest(method, url, bytes.NewReader(requestBody))
	w := httptest.NewRecorder()
	return req, w
}

func TestPostsIndex(t *testing.T) {
	router := router()
	// fmt.Println("Registered routes:", router.Routes())
	req := httptest.NewRequest("GET", "/api/posts", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, req)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestPostCreate(t *testing.T) {
	router := router()
	newEntry := models.Post{Title: "Hi", Body: "Hello"}
	req, w := savePost(t, newEntry, "POST", 0)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostsShow(t *testing.T) {
	router := router()
	// create a post 
	newEntry := models.Post{Title: "Hi", Body: "Hello"}
	req, w := savePost(t, newEntry, "POST", 0)
	router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusCreated, w.Code)

	url := fmt.Sprintf("/api/posts/%d", 1)
	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostsUpdate(t *testing.T) {
	router := router()
	// create a post 
	newEntry := models.Post{Title: "Hi", Body: "Hello"}
	req, w := savePost(t, newEntry, "POST", 0)
	router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusCreated, w.Code)

	// update post 
	updatedEntry := models.Post{Title: "Go", Body: "Lang"}
	req, w = savePost(t, updatedEntry, "PATCH", 1)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostsDelete(t *testing.T) {
	router := router()
	// create a post 
	newEntry := models.Post{Title: "Hi", Body: "Hello"}
	req, w := savePost(t, newEntry, "POST", 0)
	router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusCreated, w.Code)

	url := fmt.Sprintf("/api/posts/%d", 1)
	req = httptest.NewRequest("DELETE", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}