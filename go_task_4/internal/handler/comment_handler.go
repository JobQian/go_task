package handler

import "go_task_4/internal/service"

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(service *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: service}
}
