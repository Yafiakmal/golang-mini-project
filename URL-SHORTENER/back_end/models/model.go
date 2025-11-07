package models

import (
	"log"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Url      string `gorm:"type:varchar(2048);not null"`
	ShortUrl string `gorm:"type:varchar(512);not null;unique"`
	UserID   string `gorm:"type:varchar(256);not null"`
}

func AutoMigrate(db *gorm.DB) error {
	log.Printf("Dropping All Table")
	if err := db.Migrator().DropTable(&Url{}); err != nil {
		return err
	}
	log.Printf("Auto Migrate Table")
	if err := db.AutoMigrate(&Url{}); err != nil {
		return err
	}
	return nil

}
