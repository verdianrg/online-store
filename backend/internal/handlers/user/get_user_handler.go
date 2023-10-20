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

func GetUserHandler(rt *service.Runtime) func(r user.GetV1UserIDParams, p *models.Principal) middleware.Responder {
	return func(r user.GetV1UserIDParams, p *models.Principal) middleware.Responder {
		ctx := context.Background()
		data, err := GetUser(rt, ctx, r.ID)
		if err != nil {
			rt.Error().Println(err)
			errData := rt.GetError(err)
			return user.NewGetV1UserIDDefault(int(errData.Code())).WithPayload(&models.BaseResponse{
				Code:    errData.Code(),
				Message: err.Error(),
			})
		}

		return user.NewGetV1UserIDOK().WithPayload(&user.GetV1UserIDOKBody{
			Data: data,
		})
	}
}

func GetUser(rt *service.Runtime, ctx context.Context, id int64) (data *models.User, err error) {
	data, err = resources.GetUserByID(rt, ctx, id)
	if err != nil {
		rt.Error().Println(err)
		err = rt.SetError(http.StatusBadRequest, "user not found!")
		return
	}

	return
}
