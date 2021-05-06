package tasks

import (
	"github.com/hibiken/asynq"
)

const (
	SendEmail = "send:email"
)

func NewSendEmailTask(email string, body string) *asynq.Task {
	// Specify task payload.
	payload := map[string]interface{}{
		"email": email, // set email
		"body":  body,  // set body
	}

	// Return a new task with given type and payload.
	return asynq.NewTask(SendEmail, payload)
}
