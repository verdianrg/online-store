package handlers

import (
	"context"
	"net/http"
	service "online-store"
	"online-store/gen/models"
	"online-store/gen/restapi/operations/user"
	"online-store/internal/resources"

	"github.com/go-openapi/runtime/middleware"
)

func CreateUserHandler(rt *service.Runtime) func(r user.PostV1UserParams) middleware.Responder {
	return func(r user.PostV1UserParams) middleware.Responder {
		ctx := context.Background()
		data := &models.User{
			UserData: r.Data.UserData,
		}

		err := CreateUser(rt, ctx, data)
		if err != nil {
			rt.Error().Println(err)
			errData := rt.GetError(err)
			return user.NewPostV1UserDefault(int(errData.Code())).WithPayload(&models.BaseResponse{
				Code:    errData.Code(),
				Message: err.Error(),
			})
		}

		return user.NewPostV1UserCreated().WithPayload(&models.User{
			UserData: data.UserData,
		})
	}
}

func CreateUser(rt *service.Runtime, ctx context.Context, data *models.User) error {
	err := ValidateUser(rt, ctx, data.Username)
	if err != nil {
		return err
	}

	err = resources.CreateUser(rt, ctx, data)
	return err
}

func ValidateUser(rt *service.Runtime, ctx context.Context, email string) error {
	_, err := resources.GetUserByUserName(rt, ctx, email)
	if err == nil {
		return rt.SetError(http.StatusBadRequest, "email already exist!")
	}

	return nil
}
