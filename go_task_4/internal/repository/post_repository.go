package repository

import (
	"go_task_4/internal/model"

	"gorm.io/gorm"
)

type PostRepositoryInterface interface {
	Create(post *model.Post) error
	GetByID(id int) (*model.Post, error)
	GetByUserID(userid int) ([]*model.Post, error)
	FindAll() ([]*model.Post, error)
	Update(post *model.Post) error
	Delete(id int) error
	Count() (int, error)
}

// 注入db实例
type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (p *PostRepository) Create(post *model.Post) error {
	err := p.db.Model(&model.Post{}).Create(post).Error
	return err
}

// 预加载用户
func (p *PostRepository) GetByID(id int) (*model.Post, error) {
	post := model.Post{}
	err := p.db.Model(&model.Post{}).Preload("User").Where("id = ?", id).First(&post).Error
	return &post, err
}

func (p *PostRepository) GetByUserID(id int) (*[]model.Post, error) {
	posts := []model.Post{}
	err := p.db.Model(&model.Post{}).Where("user_id = ?", id).Find(&posts).Error
	return &posts, err
}

func (p *PostRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	// Find all posts and preload the user for each post
	err := p.db.Preload("User").Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (p *PostRepository) Update(post model.Post) error {
	err := p.db.Model(&model.Post{}).Save(&post).Error
	return err
}

// 软删除
func (p *PostRepository) Delete(post model.Post) error {
	err := p.db.Model(&model.User{}).Delete(&post).Error
	return err
}

func (p *PostRepository) Count() error {
	var total int64
	err := p.db.Model(&model.Post{}).Count(&total).Error
	return err
}
