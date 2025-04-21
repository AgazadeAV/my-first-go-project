package job

import (
	"github.com/google/uuid"
	"log"
)

type WelcomeEmailJob struct {
	UserID uuid.UUID
	Email  string
}

func (job WelcomeEmailJob) CustomExecute() {
	log.Printf("Sending welcome email to %s (UserID: %s)", job.Email, job.UserID)
}
