package controller

import (
	"net/http"
	"strconv"

	"github.com/ardi7923/test-sagara/entity"
	"github.com/ardi7923/test-sagara/helper"
	"github.com/ardi7923/test-sagara/request"
	"github.com/ardi7923/test-sagara/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
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
	var loginRequest request.LoginRequest
	errRequest := ctx.ShouldBind(&loginRequest)
	if errRequest != nil {
		res := helper.BuildErrorResponse("Failed to proccess request", errRequest.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authResult := c.authService.VerifyCredential(loginRequest.Username, loginRequest.Password)
	if v, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.Id, 10))
		v.Token = generateToken
		res := helper.BuildResponse(true, "OK", v)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := helper.BuildErrorResponse("Please check your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)

}

func (c *authController) Register(ctx *gin.Context) {
	var registerRequest request.UserCreateRequest
	errRequest := ctx.ShouldBind(&registerRequest)
	if errRequest != nil {
		res := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
		return
	}

	if c.authService.IsDuplicateUsername(registerRequest.Username) {
		res := helper.BuildErrorResponse("Failed to proccess request", "Duplicate Username", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		createdUser := c.authService.CreateUser(registerRequest)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.Id, 10))
		createdUser.Token = token
		res := helper.BuildResponse(true, "OK", createdUser)
		ctx.JSON(http.StatusCreated, res)
	}
}
