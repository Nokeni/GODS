package middlewares

import (
	"net/http"

	"github.com/Nokeni/GODS/internal/web/api/services"
	"github.com/gin-gonic/gin"
)

// AdminMiddleware checks if the authenticated user is an admin.
func AdminMiddleware(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user from the context
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		// Get the user from the database
		uid, ok := userID.(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		user, err := userService.Get(uint(uid))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		// Check if the user is in the "admin" group
		isAdmin := false
		for _, group := range user.Groups {
			if group.Name == "admin" {
				isAdmin = true
				break
			}
		}
		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			c.Abort()
			return
		}

		// Continue to the next handler
		c.Next()
	}
}
