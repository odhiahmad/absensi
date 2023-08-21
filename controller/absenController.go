package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/absensi/dto"
	"github.com/odhiahmad/absensi/helper"
	"github.com/odhiahmad/absensi/service"
)

type AbsenController interface {
	CreateAbsen(ctx *gin.Context)
	UpdateAbsen(ctx *gin.Context)
}

type absenController struct {
	absenService service.AbsenService
	jwtService  service.JWTService
}

func NewAbsenController(absenService service.AbsenService, jwtService service.JWTService) AbsenController {
	return &absenController{
		absenService: absenService,
		jwtService:  jwtService,
	}
}

func (c *absenController) CreateAbsen(ctx *gin.Context) {
	var absenCreateDTO dto.AbsenCreateDTO
	errDTO := ctx.ShouldBind(&absenCreateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	
		createdAbsen := c.absenService.CreateAbsen(absenCreateDTO)
		response := helper.BuildResponse(true, "!OK", createdAbsen)
		ctx.JSON(http.StatusCreated, response)
	
}

func (c *absenController) UpdateAbsen(ctx *gin.Context) {
	var absenUpdateDTO dto.AbsenUpdateDTO
	errDTO := ctx.ShouldBind(&absenUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	
		updatedAbsen := c.absenService.UpdateAbsen(absenUpdateDTO)
		response := helper.BuildResponse(true, "!OK", updatedAbsen)
		ctx.JSON(http.StatusCreated, response)
	
}
