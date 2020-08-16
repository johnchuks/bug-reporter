package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/johnchuks/bug-reporter/responses"
	"io/ioutil"
	"github.com/johnchuks/bug-reporter/models"
)

// CreateReport handles creating a new bug report
func (a *App) CreateReport(w http.ResponseWriter, r *http.Request) {
	
	report := &models.Report{}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(requestBody, &report)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	report.Strip() // strip white spaces from the text

	createdReport, err := report.Create(a.DB)

	resp := map[string]interface{}{"status": "success", "message": "Bug report created successfully"}

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["report"] = createdReport
	responses.JSON(w, http.StatusOK, resp)
	return

}