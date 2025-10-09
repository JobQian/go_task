package handler

import "go_task_4/internal/service"

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{postService: service}
}
