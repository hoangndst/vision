package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&UserModel{}); err != nil {
		return err
	}
	return nil
}
