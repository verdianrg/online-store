package rest

import (
	"net/http"
	service "online-store"
	"online-store/gen/models"
	"online-store/gen/restapi/operations"
	"online-store/internal/utils"
)

func Authorization(rt *service.Runtime, api *operations.OnlineStoreServerAPI) {
	api.KeyAuth = func(token string) (p *models.Principal, err error) {
		claims, err := utils.ParseAndCheckToken(token)
		if err != nil {
			return nil, rt.SetError(http.StatusUnauthorized, "Unauthorized: insufficient API key privileges")
		}

		principal := new(models.Principal)
		if claims.Role != "" {
			principal.Role = claims.Role
		}

		if claims.UserID != 0 {
			principal.UserID = claims.UserID
		}

		p = principal
		return
	}
}
