package controllers

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/johnchuks/feature-reporter/responses"
	"io/ioutil"
	"github.com/johnchuks/feature-reporter/models"
	"github.com/slack-go/slack"
)

// CreateReport handles creating a new bug report
func (a *App) CreateReport(w http.ResponseWriter, r *http.Request) {
	
	report := &models.Report{}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var message slack.AttachmentActionCallback
	if err := json.Unmarshal(requestBody, &message); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if  message.Token != a.SlackVerificationToken {
		log.Printf("[ERROR] Invalid token: %s", message.Token)
		responses.ERROR(w, http.StatusUnauthorized, nil)
	}

	report.Strip() // strip white spaces from the text

	createdReport, err := report.Create(a.DB)

	resp := map[string]interface{}{"status": "success", "message": "Feature created successfully"}

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["report"] = createdReport
	responses.JSON(w, http.StatusOK, resp)
	return

}