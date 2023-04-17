package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Interfaces/models"
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
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	err := u.UserService.CreateUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, "Success")
}

func (u *UserController) GetUser(ctx *gin.Context) {
	var username string = ctx.Param("name")
	user, err := u.UserService.GetUser(&username)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) GetAllUsers(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}
	err := u.UserService.UpdateUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{
			"Error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message":"User updated successfully",
		"User":user,
	})
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	var username string = ctx.Param("name")
	err := u.UserService.DeleteUser(&username)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{
			"Error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message":"Deleted Successfully",
	})
}

func (ux *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
	userRoute := rg.Group("/user")
	userRoute.POST("/create",ux.CreateUser)
	userRoute.PATCH("/update",ux.UpdateUser)
	userRoute.GET("/getall",ux.GetAllUsers)
	userRoute.DELETE("/delete",ux.DeleteUser)
	userRoute.GET("/get/:name",ux.GetUser)

}
