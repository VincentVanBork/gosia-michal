package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	FirstName    *string
	LastName     *string
	InvitationID uint
}

type Invitation struct {
	gorm.Model
	Email              sql.NullString
	Token              string
	IsSingle           bool `gorm:"default:false"`
	HasKids            bool `gorm:"default:false"`
	IsWedding          bool `gorm:"default:false"`
	IsWeddingReception bool `gorm:"default:false"`
	Hotel              bool `gorm:"default:false"`
	Transport          bool `gorm:"default:false"`
	Guests             []Guest
}
