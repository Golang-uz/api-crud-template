package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realtemirov/api-crud-template/config"
	"github.com/realtemirov/api-crud-template/services"
	"github.com/rs/zerolog"
)

const (
	ErrorBadRequest = "bad request"
	ErrorInternal   = "internal error"
	ErrorNotFound   = "not found"
	Success         = "success"
	OK              = "ok"
	InvalidID       = "invalid id"
	InvalidEMail    = "invalid email"
	InvalidUsernam  = "invalid username"
)

type Handler struct {
	log         zerolog.Logger
	cnf         *config.Config
	UserHandler *userHandler
	PostHandler *postHandler
}

func NewHandler(cnf *config.Config, services *services.Services, log zerolog.Logger) *Handler {
	return &Handler{
		log:         log,
		cnf:         cnf,
		UserHandler: NewUserHandler(services.UserService, log),
		PostHandler: NewPostHandler(services.PostService, log),
	}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

func (h *Handler) Default(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"api": "test",
	})
}

func response(c *gin.Context, code int, err error, data interface{}, message string) {
	if err == nil {
		err = fmt.Errorf("")
	}
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"error":   err.Error(),
		"data":    data,
	})
}

// func errorResponse(c *gin.Context, code int, err error, message string) {
// 	c.JSON(code, gin.H{
// 		"code":    code,
// 		"message": message,
// 		"error":   err,
// 		"data":    nil,
// 	})
// 	// }, models.Response{
// 	// 	Code:    code,
// 	// 	Data:    nil,
// 	// 	Message: message,
// 	// 	Error:   err,
// 	// })
// }
