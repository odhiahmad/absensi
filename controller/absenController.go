package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/apiuser/dto"
	"github.com/odhiahmad/apiuser/helper"
	"github.com/odhiahmad/apiuser/service"
)

type AbsenController interface {
	CreateAbsen(ctx *gin.Context)
	UpdateAbsen(ctx *gin.Context)
}

type userController struct {
	userService service.AbsenService
	jwtService  service.JWTService
}

func NewAbsenController(userService service.AbsenService, jwtService service.JWTService) AbsenController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) CreateAbsen(ctx *gin.Context) {
	var userCreateDTO dto.AbsenCreateDTO
	errDTO := ctx.ShouldBind(&userCreateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.userService.IsDuplicateAbsenname(userCreateDTO.Absenname) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate username", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdAbsen := c.userService.CreateAbsen(userCreateDTO)
		response := helper.BuildResponse(true, "!OK", createdAbsen)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *userController) UpdateAbsen(ctx *gin.Context) {
	var userUpdateDTO dto.AbsenUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.userService.IsDuplicateAbsenname(userUpdateDTO.Absenname) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate username", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		updatedAbsen := c.userService.UpdateAbsen(userUpdateDTO)
		response := helper.BuildResponse(true, "!OK", updatedAbsen)
		ctx.JSON(http.StatusCreated, response)
	}
}
