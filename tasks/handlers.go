package tasks

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

// HandleSendEmailTask
func HandleSendEmailTask(c context.Context, t *asynq.Task) error {
	// Get email from given task.
	email, err := t.Payload.GetString("email")
	if err != nil {
		return err
	}

	// Get email body from given task.
	body, err := t.Payload.GetString("body")
	if err != nil {
		return err
	}

	// mock sending email
	fmt.Println("Sending Email To :", email)
	fmt.Println("Email Body :")
	fmt.Println(body)

	return nil
}
