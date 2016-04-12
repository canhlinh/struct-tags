package main

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/goji/param"
	"github.com/zenazn/goji/web"
)

func GetIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Techtalk 03"))
}

func GetUser(c web.C, w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		GenerateJsonResponse(w, ErrorResponse{
			Code:    1000,
			Message: err.Error(),
		})
		return
	}
	var user User
	if err := databaseConnection.Find(&user, userID).Error; err != nil {

	}
	GenerateJsonResponse(w, user)
}

func CreateUser(c web.C, w http.ResponseWriter, r *http.Request) {
	firstName := r.PostFormValue("first_name")
	lastName := r.PostFormValue("last_name")
	user := User{
		FirstName: firstName,
		LastName:  lastName,
	}
	if _, err := govalidator.ValidateStruct(user); err != nil {
		GenerateJsonResponse(w, ErrorResponse{
			Code:    1000,
			Message: err.Error(),
		})
		return
	}
	GenerateJsonResponse(w, user)
}

func UpdateUser(c web.C, w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "No good!", 400)
		return
	}
	var user User
	if err := param.Parse(r.PostForm, &user); err != nil {
		GenerateJsonResponse(w, ErrorResponse{
			Code:    1000,
			Message: err.Error(),
		})
		return
	}
	if _, err := govalidator.ValidateStruct(user); err != nil {
		GenerateJsonResponse(w, ErrorResponse{
			Code:    1000,
			Message: err.Error(),
		})
		return
	}
	GenerateJsonResponse(w, user)
}

func GetConfig(c web.C, w http.ResponseWriter, r *http.Request) {
	data := ReadFileConfig()
	GenerateJsonResponse(w, data)
}
