package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/johnchuks/feature-reporter/responses"
	"io/ioutil"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack"
)


// SlackHandler handles incoming requests from slack
func (a *App) SlackHandler(w http.ResponseWriter, r *http.Request) {
	
	// report := &models.Report{}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(requestBody), 
		slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: a.SlackVerificationToken}))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(requestBody), &r)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		resp := map[string]interface{}{"Challenge": r.Challenge}
		responses.JSON(w, http.StatusOK, resp)
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			// Header Section
			headerText := slack.NewTextBlockObject("mrkdwn", "**Create a New Feature Request**", false, false)
			headerSection := slack.NewSectionBlock(headerText, nil, nil)

			// Button Section
			selectButtonText := slack.NewTextBlockObject("plain_text", "Select", true, false)
			selectCoreButton := slack.NewButtonBlockElement("", "C", selectButtonText)
			selectMobileButton := slack.NewButtonBlockElement("", "M", selectButtonText)
			selectARButton := slack.NewButtonBlockElement("", "AR", selectButtonText)

			//Fields Section
			titleField := slack.NewTextBlockObject("mrkdwn", "*Title", false, false)
			descriptionField := slack.NewTextBlockObject("mrkdown", "*Description*", false, false)
			projectField := slack.NewTextBlockObject("mrkdown", "*Select a project*", false, false)

			// Option 1
			optionOneText := slack.NewTextBlockObject("mrkdwn", "*Findspace-Core*", false, false)
			optionOneSection := slack.NewSectionBlock(optionOneText, nil, slack.NewAccessory(selectCoreButton))

			// Option 2
			optionTwoText := slack.NewTextBlockObject("mrkdwn", "*Findspace Mobile*", false, false)
			optionTwoSection := slack.NewSectionBlock(optionTwoText, nil, slack.NewAccessory(selectMobileButton))

			// Option 3
			optionThreeText := slack.NewTextBlockObject("mrkdwn", "*Findspace Data Aggregation*", false, false)
			optionThreeSection := slack.NewSectionBlock(optionThreeText, nil, slack.NewAccessory(selectARButton))

			message := slack.NewBlockMessage(
				headerSection,
				titleField,
				descriptionField,
				projectField,
				optionOneSection,
				optionTwoSection,
				optionThreeSection,
			)

			b, err := json.MarshalIndent(message, "", "    ")
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			a.SlackClient.PostMessage(ev.Channel, slack.MsgOptionText(string(b), false))
		}
	}

}
