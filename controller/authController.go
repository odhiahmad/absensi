package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/absensi/dto"
	"github.com/odhiahmad/absensi/entity"
	"github.com/odhiahmad/absensi/helper"
	"github.com/odhiahmad/absensi/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return

	}
	authResult := c.authService.VerifyCredential(loginDTO.Username, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(v.Username)
		v.Token = generatedToken
		response := helper.BuildResponse(true, "Berhail Login!", v)
		ctx.JSON(http.StatusOK, response)
		return

	}
	response := helper.BuildResponse(false, "Invalid credential", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

}
