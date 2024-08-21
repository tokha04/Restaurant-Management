package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tokha04/go-restautant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.GET("/users/login", controllers.Login())
}
