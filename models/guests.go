package models

import "gorm.io/gorm"

type Guest struct {
	gorm.Model
	FirstName    *string
	LastName     *string
	InvitationID uint
}

type Invitation struct {
	gorm.Model
	email              *string
	Token              string
	hasKids            bool
	isWedding          bool
	isWeddingReception bool
	Guests             []Guest
}
