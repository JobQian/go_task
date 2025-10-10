package handler

import (
	"go_task_4/internal/model"
	"go_task_4/internal/response"
	"go_task_4/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(service *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: service}
}

func (co *CommentHandler) CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid comment parameters")
		return
	}
	err := co.commentService.CreateComment(&comment)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)

}

func (co *CommentHandler) UpdateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid comment parameters")
		return
	}
	err := co.commentService.UpdateComment(&comment)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (co *CommentHandler) DeleteComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid comment parameters")
		return
	}
	err := co.commentService.DeleteComment(&comment)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (co *CommentHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID parameters")
		return
	}
	comment, err := co.commentService.GetByID(int(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, comment)
}

func (co *CommentHandler) GetByUserID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID parameters")
		return
	}
	comments, err := co.commentService.GetByUserID(int(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, comments)
}

func (co *CommentHandler) GetByPostID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID parameters")
		return
	}
	comments, err := co.commentService.GetByPostID(int(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, comments)
}
