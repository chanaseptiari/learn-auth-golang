package repositories

import (
	"errors"

	"github.com/chanaseptiari/learn-auth-golang/models"
	"gorm.io/gorm"
)

func GetUsername(db *gorm.DB, Input *models.LoginSearch) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	report := db.Table("ms_auth").Where("Username = ? ", Input.Username).Take(&res)

	err := errors.Is(report.Error, gorm.ErrRecordNotFound)
	if err == true {
		return nil, report.Error
	}

	return res, nil
}
