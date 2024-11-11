package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/delivery/http/handler"
	userRepo "vietime-backend/internal/repo/user"
	signUp "vietime-backend/internal/use-case/sign-up"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(db *mongo.Database, gin *gin.Engine) {
	userRP := userRepo.NewUserRepository(db)

	signUpUsecase := signUp.NewSignupUsecase(userRP)

	h := handler.NewHandler(signUpUsecase)

	publicRouter := gin.Group("")

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	publicRouter.GET("/swagger/*any", swaggerHandler)
	publicRouter.POST("/api/signup", h.SignUp)
}
