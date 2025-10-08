package repository

import (
	"go_task_4/internal/model"

	"gorm.io/gorm"
)

type CommentRepositoryInterface interface {
	Create(comment *model.Comment) error
	GetByID(id int) (*model.Comment, error)
	GetByUserID(userid int) ([]*model.Comment, error)
	GetByPostID(postid int) ([]*model.Comment, error)
	Update(comment *model.Comment) error
	Delete(id int) error
	Count() (int, error)
}

// 注入db实例
type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (c *CommentRepository) Create(comment *model.Comment) error {
	err := c.db.Model(&model.Comment{}).Create(comment).Error
	return err
}

func (c *CommentRepository) GetByID(id int) (*model.Comment, error) {
	comment := model.Comment{}
	err := c.db.Model(&model.Comment{}).Where("id = ?", id).First(&comment).Error
	return &comment, err
}

func (c *CommentRepository) GetByUserID(id int) (*[]model.Comment, error) {
	comments := []model.Comment{}
	err := c.db.Model(&model.Comment{}).Where("user_id = ?", id).Find(&comments).Error
	return &comments, err
}

func (c *CommentRepository) GetByPostID(id int) (*[]model.Comment, error) {
	comments := []model.Comment{}
	err := c.db.Model(&model.Comment{}).Where("post_id = ?", id).Find(&comments).Error
	return &comments, err
}

func (c *CommentRepository) Update(comment model.Comment) error {
	err := c.db.Model(&model.Comment{}).Save(&comment).Error
	return err
}

// 软删除
func (c *CommentRepository) Delete(comment model.Comment) error {
	err := c.db.Model(&model.User{}).Delete(&comment).Error
	return err
}

func (c *CommentRepository) Count() error {
	var total int64
	err := c.db.Model(&model.Comment{}).Count(&total).Error
	return err
}
