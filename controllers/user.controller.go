package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Interfaces/services"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (u *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (u *UserController) GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (ux *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
	userRoute := rg.Group("/user")
	userRoute.POST("/create",ux.CreateUser)
	userRoute.PATCH("/update",ux.UpdateUser)
	userRoute.GET("/getall",ux.GetAllUsers)
	userRoute.DELETE("/delete",ux.DeleteUser)
	userRoute.GET("/get/:name",ux.GetUser)

}
