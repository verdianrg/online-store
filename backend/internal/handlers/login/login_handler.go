package login

import (
	"context"
	"net/http"
	service "online-store"
	"online-store/gen/models"
	"online-store/gen/restapi/operations/login"
	"online-store/internal/resources"
	"online-store/internal/utils"

	"github.com/go-openapi/runtime/middleware"
)

func Login(rt *service.Runtime) func(r login.PostV1LoginParams) middleware.Responder {
	return func(r login.PostV1LoginParams) middleware.Responder {
		user, err := resources.GetUserLogin(rt, context.Background(), *r.Data.Username, *r.Data.Password)
		if err != nil {
			rt.Error().Println(err)
			return login.NewPostV1LoginDefault(http.StatusNotFound).WithPayload(&models.BaseResponse{
				Code:    http.StatusNotFound,
				Message: "user not found!",
			})
		}

		tokenString, err := utils.CreateToken(user.Username, user.Role, int64(user.ID))
		if err != nil {
			rt.Error().Println(err)
			return login.NewPostV1LoginDefault(http.StatusInternalServerError).WithPayload(&models.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return login.NewPostV1LoginOK().WithPayload(&models.LoginResponse{
			Username: user.Username,
			Role:     user.Role,
			Token:    tokenString,
		})
	}
}
