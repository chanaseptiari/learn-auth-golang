package repositories

import (
	"github.com/chanaseptiari/learn-auth-golang/models"
	"gorm.io/gorm"
)

// Insert to Register
func ProsesRegister(db *gorm.DB, Input *models.MsAuth) (err error) {
	if err = db.Create(&Input).Error; err != nil {
		return err
	}
	return nil
}
