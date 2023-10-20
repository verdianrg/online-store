package resources

import (
	"fmt"
	"math"
	"online-store/gen/models"
	"strings"

	"gorm.io/gorm"
)

func paginationScope(page int, limit int) func(db *gorm.DB) *gorm.DB {
	offset := (page - 1) * limit
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

func queryLikeScope(column string, keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s LIKE ?", column), strings.Join([]string{"%", keyword, "%"}, ""))
	}
}

func getPaginationData(page int, limit float64, totalData float64) *models.Pagination {
	totalPage := math.Ceil(totalData / limit)

	return &models.Pagination{
		Page:      int64(page),
		TotalData: int64(totalData),
		TotalPage: int64(totalPage),
	}
}
