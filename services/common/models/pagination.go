package models

import "gorm.io/gorm"

type Pagination[T any] struct {
	DB *gorm.DB
}

func NewPagination[T any](db *gorm.DB) *Pagination[T] {
	return &Pagination[T]{DB: db}
}

func (p *Pagination[T]) QueryPage(page, pageSize int, query func(tx *gorm.DB) *gorm.DB) (total int64, data []T, err error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	const maxPageSize = 100
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	baseTx := query(p.DB.Model(new(T)))

	// Count total
	if err = baseTx.Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		// 返回空 slice 而不是 nil
		return total, make([]T, 0), nil
	}
	tx := baseTx.Session(&gorm.Session{})
	err = tx.Limit(pageSize).Offset((page - 1) * pageSize).Find(&data).Error
	if data == nil {
		data = make([]T, 0)
	}
	return total, data, err
}
