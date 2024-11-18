package route

import (
	"vietime-backend/internal/handler"
	userRepo "vietime-backend/internal/repo/user"
	signupUC "vietime-backend/internal/use-case/sign-up"

	_ "vietime-backend/docs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			VietCard Backend API
//	@version		1.0
//	@description	Backend server for VietCard application
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	hynduf@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func Setup(db *mongo.Database, gin *gin.Engine) {
	userRP := userRepo.NewUserRepository(db)

	signUpUsecase := signupUC.NewSignUpUseCase(userRP)

	h := handler.NewHandler(signUpUsecase)

	publicRouter := gin.Group("")

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	publicRouter.GET("/swagger/*any", swaggerHandler)
	publicRouter.POST("/api/signup", h.SignUp)
	publicRouter.POST("/api/login", h.Login)
}
