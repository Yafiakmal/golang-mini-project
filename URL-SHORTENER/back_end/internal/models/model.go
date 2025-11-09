package models

import (
	"log"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Url    string `gorm:"type:varchar(2048);not null"`
	Name   string `gorm:"type:varchar(512);not null;unique"`
	UserID uint   `gorm:"not null"`
}

// func (u Url) String() string {
// 	str := "Url: " + u.Url + ", ShortUrl: " + u.Name + ", UserID: " + string(rune(u.UserID))
// 	return str
// }

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
