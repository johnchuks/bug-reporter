package controllers

import (
	"encoding/json"
	"github.com/johnchuks/feature-reporter/responses"
	"github.com/johnchuks/feature-reporter/models"
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