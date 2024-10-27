package web

import (
	_ "github.com/Nokeni/GODS/docs"
	"github.com/Nokeni/GODS/internal/web/api/handlers"
	"github.com/Nokeni/GODS/internal/web/api/middlewares"
	"github.com/Nokeni/GODS/internal/web/api/repositories"
	apiroutes "github.com/Nokeni/GODS/internal/web/api/routes"
	"github.com/Nokeni/GODS/internal/web/api/services"
	"github.com/Nokeni/GODS/internal/web/common/dtos"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title           GODS API
// @version         1.0
// @description     GODS, for gods.

// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func NewHTTPServer(database *gorm.DB) (*gin.Engine, error) {
	router := gin.Default()
	if err := router.SetTrustedProxies([]string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "127.0.0.1"}); err != nil {
		return nil, err
	}

	// Set up the api repositories
	userRepository := repositories.NewUserRepository(database)
	groupRepository := repositories.NewGroupRepository(database)
	userGroupRepository := repositories.NewUserGroupRepository(database)

	// Set up the api services
	userService := services.NewUserService(userRepository)
	groupService := services.NewGroupService(groupRepository)
	userGroupService := services.NewUserGroupService(userGroupRepository)
	authService := services.NewAuthService(userRepository)

	// Set up the api handlers
	userHandler := handlers.NewUserHandler(userService)
	groupHandler := handlers.NewGroupHandler(groupService)
	userGroupHandler := handlers.NewUserGroupHandler(userGroupService)
	authHandler := handlers.NewAuthHandler(authService)

	// Create the admin user and group
	adminUser, _ := userService.Create(&dtos.CreateUserDTO{Name: viper.GetString("ADMIN_NAME"), Email: viper.GetString("ADMIN_EMAIL"), Password: viper.GetString("ADMIN_PASSWORD")})
	adminGroup, _ := groupService.Create(&dtos.CreateGroupDTO{Name: "admin"})
	userGroupService.AddUserToGroup(adminUser.ID, adminGroup.ID)

	// Set up API routes
	apiroutes.RegisterAPIRoutes(
		router,
		userHandler,
		groupHandler,
		userGroupHandler,
		authHandler,
		middlewares.AuthMiddleware(),
		middlewares.AdminMiddleware(userService),
	)

	// Set up UI routes
	// uiroutes.SetupUIRoutes(router, nil)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router, nil
}
