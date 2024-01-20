// controllers/user_controller.go
package controllers

import (
	"be-nutreazy/models"
	"be-nutreazy/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
    userService *services.UserService
}

func NewUserController(db *gorm.DB) *UserController {
    return &UserController{
        userService: services.NewUserService(db),
    }
}

func (c *UserController) GetUsers(ctx *gin.Context) {
    users, err := c.userService.GetUsers()
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Internal Server Error"})
        return
    }
    ctx.JSON(200, users)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
        ctx.JSON(400, gin.H{"error": "id is not int"})
        return
    }
    user, err := c.userService.GetUserByID(userID)
    if err != nil {
        ctx.JSON(404, gin.H{"error": "User not found"})
        return
    }
    ctx.JSON(200, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    newUser, err := c.userService.CreateUser(user)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Internal Server Error"})
        return
    }

    ctx.JSON(201, newUser)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
    var user models.User

    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    updatedUser, err := c.userService.UpdateUser(user)
    if err != nil {
        ctx.JSON(404, gin.H{"error": "User not found"})
        return
    }

    ctx.JSON(200, updatedUser)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
    userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
        ctx.JSON(400, gin.H{"error": "id is not int"})
        return
    }
    err = c.userService.DeleteUser(userID)
    if err != nil {
        ctx.JSON(404, gin.H{"error": "User not found"})
        return
    }

    ctx.JSON(204, nil)
}
