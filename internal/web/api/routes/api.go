package routes

import (
	"github.com/Nokeni/GODS/internal/web/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(
	router *gin.Engine,
	userHandler handlers.UserHandler,
	groupHandler handlers.GroupHandler,
	userGroupHandler handlers.UserGroupHandler,
	authHandler handlers.AuthHandler,
	authMiddleware gin.HandlerFunc,
	adminMiddleware gin.HandlerFunc,
) {
	api := router.Group("/api")
	{
		userRoutes := api.Group("/users", authMiddleware, adminMiddleware)
		{
			userRoutes.GET("/", userHandler.GetAll)
			userRoutes.GET("/:id", userHandler.Get)
			userRoutes.POST("/", userHandler.Create)
			userRoutes.PUT("/:id", userHandler.Update)
			userRoutes.DELETE("/:id", userHandler.Delete)
		}

		groupRoutes := api.Group("/groups", authMiddleware, adminMiddleware)
		{
			groupRoutes.GET("/", groupHandler.GetAll)
			groupRoutes.GET("/:id", groupHandler.Get)
			groupRoutes.POST("/", groupHandler.Create)
			groupRoutes.PUT("/:id", groupHandler.Update)
			groupRoutes.DELETE("/:id", groupHandler.Delete)
		}

		userGroupRoutes := api.Group("/users-groups", authMiddleware, adminMiddleware)
		{
			userGroupRoutes.POST("/:groupId/users/:userId", userGroupHandler.AddUserToGroup)
			userGroupRoutes.DELETE("/:groupId/users/:userId", userGroupHandler.RemoveUserFromGroup)
			userGroupRoutes.GET("/users/:userId", userGroupHandler.GetUserGroups)
			userGroupRoutes.GET("/:groupId/users", userGroupHandler.GetGroupUsers)
		}

		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/login", authHandler.Login)
			authRoutes.POST("/signup", authHandler.Signup)
		}
	}
}
