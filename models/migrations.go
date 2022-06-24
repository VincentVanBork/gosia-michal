package models

import (
	"gorm.io/gorm"
	"log"
)

func MigrateAll(db *gorm.DB) {
	err := db.AutoMigrate(
		&Invitation{},
		&Guest{},
		&TableInfo{},
	)
	if err != nil {
		log.Fatalln(err)
	}

}
