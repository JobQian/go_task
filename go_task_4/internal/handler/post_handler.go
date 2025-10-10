package handler

import (
	"go_task_4/internal/model"
	"go_task_4/internal/response"
	"go_task_4/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{postService: service}
}

func (p *PostHandler) CreatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post parameters")
		return
	}
	err := p.postService.CreatePost(&post)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (p *PostHandler) UpdatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post parameters")
		return
	}
	err := p.postService.UpdatePost(&post)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (p *PostHandler) DeletePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post parameters")
		return
	}
	err := p.postService.DeletePost(&post)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (p *PostHandler) FindAllPost(c *gin.Context) {
	posts, err := p.postService.FindAllPost()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, posts)
}

func (p *PostHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID parameters")
		return
	}
	post, err := p.postService.GetByID(int(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, post)
}

func (p *PostHandler) GetByUserID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID parameters")
		return
	}
	posts, err := p.postService.GetByUserID(int(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, posts)
}
