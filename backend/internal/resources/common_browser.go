package resources

import (
	"errors"
	service "online-store"
	"online-store/gen/models"
	"reflect"

	"gorm.io/gorm"
)

// DEFINE COMMON STRUC PARAMETERS
type GetListParams struct {
	Page    *int32
	Limit   *int32
	Keyword *string
}

// DEFINE COMMON FUNCTION
func getList(rt *service.Runtime, targetData interface{}, page *int32, limit *int32, paginate *models.Pagination, callbackFn func(model, countModel *gorm.DB)) error {
	if reflect.ValueOf(targetData).Kind() != reflect.Ptr {
		return errors.New("target data must pointer")
	}

	model := rt.DB().Model(targetData)
	countModel := rt.DB().Model(targetData)

	if page != nil && limit != nil {
		model.Scopes(paginationScope(int(*page), int(*limit)))
	}

	if callbackFn != nil {
		callbackFn(model, countModel)
	}

	if page != nil && limit != nil {
		var count int64
		if err := countModel.Count(&count).Error; err != nil {
			return err
		}
		*paginate = *getPaginationData(int(*page), float64(*limit), float64(count))
	}

	if err := model.Find(targetData).Error; err != nil {
		return err
	}

	return nil
}
