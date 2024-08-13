package models

import "time"

type EmailRequest struct {
	FirstName 	string 		`json:"firstName"`
	LastName 	string 		`json:"lastName"`
	Email      	string 		`json:"email"`
	Timestamp 	time.Time 	`json:"timestamp"`
}
