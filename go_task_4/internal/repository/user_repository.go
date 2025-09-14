package repository

import (
	"go_task_4/internal/model"

	"gorm.io/gorm"
)

// UserRepository 用户仓库接口
type UserRepositoryInterface interface {
	Create(user *model.User) error
	GetByID(id int) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
	Paginate(in interface{}, page, pageSize int, out interface{}, where ...interface{}) (model.PageResult, error)
	Count() (int, error)
}

// 注入db实例
type UserRepository struct {
	db *gorm.DB
}

// 创建userRepository实例
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user model.User) error {
	err := u.db.Model(&model.User{}).Create(&user).Error
	return err
}

func (u *UserRepository) GetByID(id int) (*model.User, error) {
	user := model.User{}
	err := u.db.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepository) GetByUsername(username string) (*model.User, error) {
	user := model.User{}
	err := u.db.Model(&model.User{}).Where("username = ?", username).First(&user).Error
	return &user, err
}

func (u *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := u.db.Model(&model.User{}).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepository) Update(user model.User) error {
	err := u.db.Model(&model.User{}).Save(&user).Error
	return err
}

// 软删除
func (u *UserRepository) Delete(user model.User) error {
	err := u.db.Model(&model.User{}).Delete(&user).Error
	return err
}

func (u *UserRepository) Count() error {
	var total int64
	err := u.db.Model(&model.User{}).Count(&total).Error
	return err
}

func (u *UserRepository) Paginate(in interface{}, page, pageSize int, out interface{}, where ...interface{}) (model.PageResult, error) {
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
	query := u.db.Model(&in)

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
