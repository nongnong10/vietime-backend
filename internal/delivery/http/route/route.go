package route

import (
	config "vietime-backend/config"
	"vietime-backend/internal/delivery/http/middleware"
	"vietime-backend/internal/handler"
	cardRepo "vietime-backend/internal/repo/card"
	deckRepo "vietime-backend/internal/repo/deck"
	userRepo "vietime-backend/internal/repo/user"

	cardUC "vietime-backend/internal/use-case/card"
	deckUC "vietime-backend/internal/use-case/deck"
	signupUC "vietime-backend/internal/use-case/sign-up"
	userUC "vietime-backend/internal/use-case/user"

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
// @description					Description for what is this security definition being used
func Setup(db *mongo.Database, gin *gin.Engine) {
	userRepository := userRepo.NewUserRepository(db)
	cardRepository := cardRepo.NewCardRepository(db)
	deckRepository := deckRepo.NewDeckRepository(db)

	signUpUsecase := signupUC.NewSignUpUseCase(userRepository)
	cardUseCase := cardUC.NewCardUseCase(cardRepository)
	userUseCase := userUC.NewUserUseCase(userRepository)
	deckUseCase := deckUC.NewDeckUsecase(deckRepository, userRepository)

	h := handler.NewHandler(signUpUsecase, cardUseCase, userUseCase, deckUseCase)

	publicRouter := gin.Group("")

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	publicRouter.GET("/swagger/*any", swaggerHandler)
	publicRouter.POST("/api/signup", h.SignUp)
	publicRouter.POST("/api/login", h.Login)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(config.E.AccessTokenSecret))

	// Card
	protectedRouter.POST("/api/card/create", h.CreateCard)
	protectedRouter.PUT("/api/card/update", h.UpdateCard)

	// Deck
	protectedRouter.POST("/api/deck/create", h.CreateDeck)
	protectedRouter.PUT("/api/deck/update", h.UpdateDeck)

	// User
	protectedRouter.PUT("/api/user/update", h.UpdateUser)
}
