package middlewares

import (
	"final-project-mygram/database"
	"final-project-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var comment models.Comment
		db := database.GetDB()

		commentID, err := strconv.Atoi(ctx.Param("commentID"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		err = db.Debug().Select("user_id").First(&comment, uint(commentID)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Comment doesn't exist",
			})
			return
		}

		if comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this comment",
			})
			return
		}

		ctx.Next()
	}
}
