package controllers

import (
	"encoding/json"
	"github.com/johnchuks/feature-reporter/responses"
	"io/ioutil"
	"github.com/johnchuks/feature-reporter/models"
	"net/http"
)

// SignUp a new user 
func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Strip()
	err = user.CheckInput()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := user.Create(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusCreated, newUser.StripPassword())
	return
}


func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	login := &models.LoginPayload{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(reqBody, login)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
}