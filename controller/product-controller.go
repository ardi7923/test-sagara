package controller

import (
	"net/http"

	"github.com/ardi7923/test-sagara/entity"
	"github.com/ardi7923/test-sagara/helper"
	"github.com/ardi7923/test-sagara/service"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	All(context *gin.Context)
}

type productController struct {
	productService service.ProductService
	jwtService     service.JWTService
}

func NewProductController(productService service.ProductService, jwtService service.JWTService) ProductController {
	return &productController{
		productService: productService,
		jwtService:     jwtService,
	}
}

func (c *productController) All(context *gin.Context) {
	var products []entity.Product = c.productService.All()
	res := helper.BuildResponse(true, "OK", products)
	context.JSON(http.StatusOK, res)

}
