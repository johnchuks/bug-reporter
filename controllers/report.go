package controllers

import (
	"strconv"
	"encoding/json"
	"github.com/johnchuks/feature-reporter/responses"
	"github.com/johnchuks/feature-reporter/models"
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"

)

// CreateReport creates a new report
func (a *App) CreateReport(w http.ResponseWriter, r *http.Request) {
	report := models.Report{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(reqBody, &report)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	report.Strip()
	newReport, err := report.Create(a.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, newReport)
	return
}

// GetReport retrieves a new report based on the ID passed
func (a *App) GetReport(w http.ResponseWriter, r *http.Request) {
	rep := &models.Report{}
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		responses.ERROR(w, http.StatusBadRequest, nil)
		return
	}
	
	newID, err := strconv.Atoi(id)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	existingReport, err := rep.Get(newID, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, existingReport)
	return
}