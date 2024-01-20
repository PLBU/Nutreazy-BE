package routes

import (
	"be-nutreazy/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB) {	
	userController := controllers.NewUserController(db)

    userRoutes := router.Group("/users")
    {
        userRoutes.GET("/:id", userController.GetUserByID)
        userRoutes.GET("/", userController.GetUsers)
        userRoutes.POST("/", userController.CreateUser)
        userRoutes.PUT("/", userController.UpdateUser)
        userRoutes.DELETE("/:id", userController.DeleteUser)
    }
}