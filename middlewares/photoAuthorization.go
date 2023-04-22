package middlewares

import (
	"final-project-mygram/database"
	"final-project-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var photo models.Photo
		db := database.GetDB()

		photoID, err := strconv.Atoi(ctx.Param("photoID"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		err = db.Debug().Select("user_id").First(&photo, uint(photoID)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Photo doesn't exist",
			})
			return
		}

		if photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data photo",
			})
			return
		}

		ctx.Next()
	}
}
