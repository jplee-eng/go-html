package db

import (
	"github.com/jasonpatricklee/gowebserver/inits"
	"github.com/jasonpatricklee/gowebserver/models"
	"github.com/jasonpatricklee/gowebserver/utils"
)

func AddUser(user models.User) error {
	db := inits.GetDb()

	result := db.Create(&user)
	utils.HandleError(result.Error)

	return nil
}
func RemoveUser(userID uint) error {
	db := inits.GetDb()

	var user models.User
	result := db.First(&user, userID)
	utils.HandleError(result.Error)

	result = db.Delete(&user)
	utils.HandleError(result.Error)

	return nil
}
