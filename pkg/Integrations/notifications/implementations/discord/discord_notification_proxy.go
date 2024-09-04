package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codespace-id/codespace-x/pkg/Integrations/notifications"
	"net/http"
)

type discord struct {
}

func NewDiscord() notifications.NotificationProxy {
	return &discord{}
}

func (d *discord) Send(webhookURL, message string) (err error) {
	// Create the payload as a map
	payload := map[string]string{
		"content": message,
	}

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Send the HTTP POST request
	_, err = http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	return nil
}
