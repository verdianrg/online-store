package resources

import (
	"context"
	service "online-store"
	"online-store/gen/models"
)

func CreateUser(rt *service.Runtime, ctx context.Context, user *models.User) (err error) {
	model := rt.DB().Model(&user)
	err = model.Create(&user).Error
	return
}

func GetUserLogin(rt *service.Runtime, ctx context.Context, username, password string) (user *models.User, err error) {
	model := rt.DB().Model(&user)
	model = model.First(&user, "username = ? AND password = ?", username, password)
	if model.Error != nil {
		err = model.Error
		return
	}

	return
}

func GetUserByID(rt *service.Runtime, ctx context.Context, id int64) (user *models.User, err error) {
	model := rt.DB().Model(&user)
	model = model.First(&user, id)

	if err = model.Error; err != nil {
		return
	}

	return
}

func GetUserByUserName(rt *service.Runtime, ctx context.Context, email string) (user *models.User, err error) {
	err = rt.DB().Where("username = ?", email).First(&user).Error
	if err != nil {
		return
	}

	return
}
