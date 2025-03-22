package user

import "github.com/gin-gonic/gin"

// Port
type IUserHandler interface {
	FetchUsers(c *gin.Context)
}
