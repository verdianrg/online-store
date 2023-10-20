package history

import (
	"net/http"
	service "online-store"
	"online-store/gen/models"
	"online-store/gen/restapi/operations/history"
	"online-store/internal/resources"

	"github.com/go-openapi/runtime/middleware"
)

func CreateHistoryHandler(rt *service.Runtime, r *history.PostV1HistoriesParams, p *models.Principal) middleware.Responder {
	err := resources.CreateHistories(rt, &models.HistoryData{
		MetaData: *r.Data.Metadata,
	}, p.UserID)

	if err != nil {
		rt.Error().Println(err)
		return history.NewPostV1HistoriesDefault(http.StatusInternalServerError).WithPayload(&models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return history.NewPostV1HistoriesCreated().WithPayload(&models.BaseResponse{
		Code:    http.StatusCreated,
		Message: "success create new history",
	})
}
