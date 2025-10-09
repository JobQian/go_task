package service

import (
	"errors"
	"go_task_4/internal/model"
	"go_task_4/internal/repository"
	"go_task_4/pkg/utils"
)

type PostService struct {
	postrepository *repository.PostRepository
}

func NewPostService(postrepository *repository.PostRepository) *PostService {
	return &PostService{postrepository: postrepository}
}

func (p *PostService) CreatePost(post *model.Post) error {
	if err := utils.CheckNotEmpty("Title", post.Title); err != nil {
		return err
	}
	if err := utils.CheckNotEmpty("Content", post.Content); err != nil {
		return err
	}
	if post.UserId == 0 {
		return errors.New("invalid userid")
	}
	err := p.postrepository.Create(post)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostService) UpdatePost(post *model.Post) error {
	if post.ID == 0 {
		return errors.New("invalid ID")
	}
	err := p.postrepository.Update(*post)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostService) DeletePost(post *model.Post) error {
	if post.ID == 0 {
		return errors.New("invalid ID")
	}
	err := p.postrepository.Delete(*post)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostService) FindAllPost() ([]model.Post, error) {
	posts, err := p.postrepository.FindAll()
	if err != nil {
		return []model.Post{}, err
	}
	return posts, nil
}

func (p *PostService) CountAllPost() (int64, error) {
	total, err := p.postrepository.Count()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (p *PostService) GetByID(id int) (*model.Post, error) {
	if id == 0 {
		return &model.Post{}, errors.New("invalid ID")
	}
	post, err := p.postrepository.GetByID(id)
	if err != nil {
		return &model.Post{}, err
	}
	return post, nil
}
func (p *PostService) GetByUserID(id int) ([]model.Post, error) {
	if id == 0 {
		return []model.Post{}, errors.New("invalid ID")
	}
	posts, err := p.postrepository.GetByUserID(id)
	if err != nil {
		return []model.Post{}, err
	}
	return posts, nil
}
