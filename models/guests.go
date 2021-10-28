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
	IsWedding          bool `gorm:"default:true"`
	IsWeddingReception bool `gorm:"default:true"`
	Guests             []Guest
}
