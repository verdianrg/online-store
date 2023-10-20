package cart

import (
	"net/http"
	service "online-store"
	"online-store/gen/models"
	"online-store/gen/restapi/operations/cart"
	"online-store/internal/resources"

	"github.com/go-openapi/runtime/middleware"
)

func GetCartsHandler(rt *service.Runtime, r cart.GetV1CartsParams, p *models.Principal) middleware.Responder {
	params := resources.GetListParams{
		Page:    r.Page,
		Limit:   r.Limit,
		Keyword: r.Keyword,
	}

	data, pagination, err := resources.GetCarts(rt, &params, p.UserID)
	if err != nil {
		rt.Error().Println(err)
		return cart.NewGetV1CartsDefault(http.StatusInternalServerError).WithPayload(&models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return cart.NewGetV1CartsOK().WithPayload(&cart.GetV1CartsOKBody{
		Data:     data,
		Metadata: pagination,
	})
}
