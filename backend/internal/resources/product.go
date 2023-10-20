package resources

import (
	service "online-store"
	"online-store/gen/models"

	"gorm.io/gorm"
)

func CreateProduct(rt *service.Runtime, data *models.ProductData) error {
	product := &models.Product{
		ProductData: *data,
	}

	err := rt.DB().Model(product).Save(product).Error
	return err
}

func GetProducts(rt *service.Runtime, params *GetListParams) ([]*models.Product, *models.Pagination, error) {
	products, paginate := make([]*models.Product, 0), new(models.Pagination)

	err := getList(rt, &products, params.Page, params.Limit, paginate, func(model, countModel *gorm.DB) {
		if params.Keyword != nil {
			model.Scopes(queryLikeScope("name", *params.Keyword))
			countModel.Scopes(queryLikeScope("name", *params.Keyword))
		}
	})

	return products, paginate, err
}

func GetProductByIDs(rt *service.Runtime, id []uint64) ([]*models.Product, error) {
	product := make([]*models.Product, 0)
	err := rt.DB().Model(&product).Where("id IN (?)", id).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}
