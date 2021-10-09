package models

import "gorm.io/gorm"

type Guest struct {
	gorm.Model
	FirstName    *string
	LastName     *string
	phoneNumber  string
	InvitationID int
	Invitation   Invitation
}

type Invitation struct {
	gorm.Model
	email              *string
	token              *string
	hasKids            bool
	isWedding          bool
	isWeddingReception bool
}
