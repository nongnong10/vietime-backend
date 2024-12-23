package route

import (
	config "vietime-backend/config"
	"vietime-backend/internal/delivery/http/middleware"
	"vietime-backend/internal/handler"
	cardRepo "vietime-backend/internal/repo/card"
	deckRepo "vietime-backend/internal/repo/deck"
	factRepo "vietime-backend/internal/repo/fact"
	userRepo "vietime-backend/internal/repo/user"

	cardUC "vietime-backend/internal/use-case/card"
	deckUC "vietime-backend/internal/use-case/deck"
	factUC "vietime-backend/internal/use-case/fact"
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
	factRepository := factRepo.NewFactRepository(db)

	signUpUsecase := signupUC.NewSignUpUseCase(userRepository)
	cardUseCase := cardUC.NewCardUseCase(cardRepository)
	userUseCase := userUC.NewUserUseCase(userRepository)
	deckUseCase := deckUC.NewDeckUsecase(deckRepository, userRepository, cardRepository)
	factUseCase := factUC.NewFactUseCase(factRepository)

	h := handler.NewHandler(signUpUsecase, cardUseCase, userUseCase, deckUseCase, factUseCase)

	publicRouter := gin.Group("")

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	publicRouter.GET("/swagger/*any", swaggerHandler)

	// Login/Signup
	publicRouter.POST("/api/signup", h.SignUp)
	publicRouter.POST("/api/login", h.Login)
	publicRouter.POST("/api/refresh", h.RefreshToken)

	// Mobile
	publicRouter.POST("/api/login-get-all", h.LogInGetAllData)
	publicRouter.POST("/api/signup-get-all", h.SignUpGetAllData)
	publicRouter.POST("/api/get-all", h.GetAllData)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(config.E.AccessTokenSecret))

	// Card
	protectedRouter.POST("/api/card/create", h.CreateCard)
	protectedRouter.POST("/api/card/create_list", h.CreateManyCards)
	protectedRouter.PUT("/api/card/update", h.UpdateCard)
	protectedRouter.POST("/api/card/copy", h.CopyCardToDeck)
	protectedRouter.PUT("/api/card/review", h.UpdateReviewCards)
	protectedRouter.DELETE("/api/card/delete", h.DeleteCard)

	// Deck
	protectedRouter.POST("/api/deck/create", h.CreateDeck)
	protectedRouter.PUT("/api/deck/update", h.UpdateDeck)
	publicRouter.PUT("/api/deck/view", h.UpdateViewDeck)
	protectedRouter.POST("/api/deck/copy", h.CopyDeck)
	protectedRouter.DELETE("/api/deck/delete", h.DeleteDeck)
	protectedRouter.GET("/api/deck/review-cards", h.GetDeckWithReviewCards)

	// User
	protectedRouter.PUT("/api/user/update", h.UpdateUser)

	// Fact
	protectedRouter.POST("/api/fact/create", h.CreateFact)
	publicRouter.GET("/api/fact", h.GetFact)
}
