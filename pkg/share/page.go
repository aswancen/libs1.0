package share

import "gorm.io/gorm"

// Paginate 分页器(offset从1开始)
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	if pageSize > 10000 {
		pageSize = 10000
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
