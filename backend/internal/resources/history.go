package resources

import (
	service "online-store"
	"online-store/gen/models"

	"gorm.io/gorm"
)

func CreateHistories(rt *service.Runtime, data *models.HistoryData, userID int64) error {
	history := &models.History{
		HistoryData: *data,
		UserID:      userID,
	}

	err := rt.DB().Model(history).Create(history).Error
	return err
}

func GetHistories(rt *service.Runtime, params GetListParams) ([]*models.History, *models.Pagination, error) {
	histories, paginate := make([]*models.History, 0), new(models.Pagination)

	err := getList(rt, &histories, params.Page, params.Limit, paginate, func(model, countModel *gorm.DB) {})
	return histories, paginate, err
}
