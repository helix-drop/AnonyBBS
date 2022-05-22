package handlers

import (
	"github.com/helix-drop/AnonyBBS/models"
	"gorm.io/gorm"
)

func IsExistName(name string) bool {
	var user models.User
	err := models.DB.Where("name=?", name).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}
	return true
}
func IsExistEmail(Email string) bool {
	var user models.User
	err := models.DB.Where("name=?", Email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}
	return true
}
