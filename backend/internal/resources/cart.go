package resources

import (
	service "online-store"
	"online-store/gen/models"

	"gorm.io/gorm"
)

func CreateCarts(rt *service.Runtime, data []*models.Cart, userID int64) error {
	err := rt.DB().Model(data).Save(data).Error
	return err
}

func GetCarts(rt *service.Runtime, params *GetListParams, userID int64) ([]*models.Cart, *models.Pagination, error) {
	carts, paginate := make([]*models.Cart, 0), new(models.Pagination)

	err := getList(rt, &carts, params.Page, params.Limit, paginate, func(model, countModel *gorm.DB) {
		model.Preload("Product")
		if params.Keyword != nil {
			model.Scopes(queryLikeScope("name", *params.Keyword))
			countModel.Scopes(queryLikeScope("name", *params.Keyword))
		}

		model.Where("user_id = ?", userID)
	})

	return carts, paginate, err
}
