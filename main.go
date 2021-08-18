package main

import (
	"github.com/ardi7923/test-sagara/config"
	"github.com/ardi7923/test-sagara/controller"
	"github.com/ardi7923/test-sagara/repository"
	"github.com/ardi7923/test-sagara/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetUpDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	authService       service.AuthService          = service.NewAuthService(userRepository)
	jwtService        service.JWTService           = service.NewJWTService()
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	productService    service.ProductService       = service.NewProductService(productRepository)
	productController controller.ProductController = controller.NewProductController(productService, jwtService)
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed To Load ENV")
	}

	defer config.CloseDatabaseConnection(db)
	router := gin.New()
	baseUrl := "api/v1"

	// authentication Route
	authRoutes := router.Group(baseUrl + "/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	// Product Route
	productRoutes := router.Group(baseUrl + "/product")
	{
		productRoutes.GET("/", productController.All)
	}

	router.Run(":8080")
}
