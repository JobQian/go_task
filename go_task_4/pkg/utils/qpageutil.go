package utils

import (
	"go_task_4/internal/model"

	"gorm.io/gorm"
)

// 通用分页查询
func Paginate(db *gorm.DB, in interface{}, page, pageSize int, out interface{}, where ...interface{}) (model.PageResult, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100 // 限制最大分页
	}

	var total int64
	query := db.Model(&in)

	// 如果有 where 条件，应用它
	if len(where) > 0 {
		query = query.Where(where[0], where[1:]...)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return model.PageResult{}, err
	}

	// 查数据
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&out).Error; err != nil {
		return model.PageResult{}, err
	}

	// 返回结果
	return model.PageResult{
		Data:       out,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}, nil
}
