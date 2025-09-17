package worker

import (
	"fmt"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// GetUserEmailWorker equivalent to Python's @worker_task get_user_email function
func GetUserEmailWorker(task *model.Task) (interface{}, error) {
	userid, exists := task.InputData["userid"]
	if !exists {
		return nil, fmt.Errorf("userid not provided")
	}

	// Same logic as Python: return f'{userid}@example.com'
	userEmail := fmt.Sprintf("%v@example.com", userid)

	return map[string]interface{}{
		"result": userEmail,
	}, nil
}

// SendEmailWorker equivalent to Python's @worker_task send_email function
func SendEmailWorker(task *model.Task) (interface{}, error) {
	email, hasEmail := task.InputData["email"]
	subject, hasSubject := task.InputData["subject"]
	body, hasBody := task.InputData["body"]

	if !hasEmail {
		return nil, fmt.Errorf("email address not provided")
	}

	// Default values like Python function
	if !hasSubject {
		subject = "Notification from Conductor"
	}
	if !hasBody {
		body = "This is an automated message"
	}

	return map[string]interface{}{
		"success": true,
		"email":   email,
		"subject": subject,
		"body":    body,
	}, nil
}
