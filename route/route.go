package route

import (
	"demo-backend/service/user"

	"github.com/gin-gonic/gin"
)

type Route struct {
	e *gin.Engine
}

// Constructor
func NewRoute(e *gin.Engine) *Route {
	return &Route{e}
}

func (r *Route) RegisterUser(handler user.IUserHandler) {
	r.e.GET("/v1/user", handler.FetchUsers)
}
