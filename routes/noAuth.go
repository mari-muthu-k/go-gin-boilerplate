package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mari-muthu-k/gin-template/controller"
)


func NoAuthGroupRoutes(r *gin.RouterGroup) {
	r.POST("/login", controller.Login)
	r.POST("/sign-up", controller.SignUp)

}