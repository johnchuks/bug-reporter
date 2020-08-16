package slackapi

import (
	"fmt"
	"log"
	"github.com/slack-go/slack"
)


type SlackListener struct {
	Client *slack.Client
	BotID string
}

//ListenAndResponse listens for specific slack events
func (s *SlackListener) ListenAndResponse() {
	rtm := s.Client.NewRTM()

	// start listening for slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(event); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

// handleMessageEvent handles message events
func (s *SlackListener) handleMessageEvent(event *slack.MessageEvent) error {
	// Only response in specific channel
	fmt.Println(event, "======>>>>>>>>>>>>")
	// if event.Channel != s.channelID {
	// 	log.Printf("%s %s", event.Channel, event.Msg.Text)
	// 	return nil
	// }
	return fmt.Errorf("failed to post message")
	
	// Only response mention to bot 
	// if !strings.HasPrefix(ev.Msg.Text)
}