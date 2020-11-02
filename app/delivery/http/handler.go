package http

import (
	"net/http"

	"github.com/alexvelfr/go-template/app"
	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	uc app.Usecase
}

// NewHandler ...
func NewHandler(uc app.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

// HelloWorld ...
func (h *Handler) HelloWorld(c *gin.Context) {
	h.uc.HelloWorld(c.Request.Context())
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
