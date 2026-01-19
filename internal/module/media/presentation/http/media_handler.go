package http

import "github.com/gin-gonic/gin"

type MediaHandler interface {
}

func NewMediaHandler() MediaHandler {
	return &mediaHandler{}
}

type mediaHandler struct {
}

func (h *mediaHandler) Paginate(c *gin.Context) {
}
