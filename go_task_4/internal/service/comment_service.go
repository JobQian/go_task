package service

import (
	"errors"
	"go_task_4/internal/model"
	"go_task_4/internal/repository"
	"go_task_4/pkg/utils"
)

type CommentService struct {
	commentrepository *repository.CommentRepository
}

func NewCommentService(commentrepository *repository.CommentRepository) *CommentService {
	return &CommentService{commentrepository: commentrepository}
}

func (c *CommentService) CreateComment(comment *model.Comment) error {
	if err := utils.CheckNotEmpty("Content", comment.Content); err != nil {
		return err
	}
	if comment.UserId == 0 {
		return errors.New("invalid UserId")
	}
	if comment.PostId == 0 {
		return errors.New("invalid PostId")
	}
	err := c.commentrepository.Create(comment)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentService) UpdateComment(comment *model.Comment) error {
	if comment.ID == 0 {
		return errors.New("invalid ID")
	}
	err := c.commentrepository.Update(*comment)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentService) DeleteComment(comment *model.Comment) error {
	if comment.ID == 0 {
		return errors.New("invalid ID")
	}
	err := c.commentrepository.Delete(*comment)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentService) GetByID(id int) (*model.Comment, error) {
	if id == 0 {
		return &model.Comment{}, errors.New("invalid ID")
	}
	comment, err := c.commentrepository.GetByID(id)
	if err != nil {
		return &model.Comment{}, err
	}
	return comment, nil
}
func (c *CommentService) GetByUserID(id int) ([]model.Comment, error) {
	if id == 0 {
		return []model.Comment{}, errors.New("invalid ID")
	}
	comments, err := c.commentrepository.GetByUserID(id)
	if err != nil {
		return []model.Comment{}, err
	}
	return comments, nil
}

func (c *CommentService) GetByPostID(id int) ([]model.Comment, error) {
	if id == 0 {
		return []model.Comment{}, errors.New("invalid ID")
	}
	comments, err := c.commentrepository.GetByPostID(id)
	if err != nil {
		return []model.Comment{}, err
	}
	return comments, nil
}
