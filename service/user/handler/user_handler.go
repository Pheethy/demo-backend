package handler

import (
	"demo-backend/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUs user.IUserUsecase
}

func NewUserHandler(userUs user.IUserUsecase) user.IUserHandler {
	return &userHandler{
		userUs: userUs,
	}
}

func (h *userHandler) FetchUsers(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.userUs.FetchUsers(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	resp := map[string]interface{}{
		"users": users,
	}

	c.JSON(http.StatusOK, resp)
}
