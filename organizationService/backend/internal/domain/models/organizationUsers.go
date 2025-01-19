package models

import "time"

type OrganizationUsers struct {
	ID             int64
	UserID         int64
	OrganizationID int64
	Name           string
	Surname        string
	Email          string
	Sex            string
	dateOfBirth    time.Time
}
