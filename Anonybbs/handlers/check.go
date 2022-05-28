package handlers

import (
	"github.com/helix-drop/AnonyBBS/models"
)

func IsExistName(Name string) bool {
	var user models.User
	models.DB.Where("Name=?", Name).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}
func IsExistEmail(Email string) bool {
	var user models.User
	models.DB.Where("Email=?", Email).Take(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}
