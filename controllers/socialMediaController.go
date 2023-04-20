package controllers

import (
	"final-project-mygram/database"
	"final-project-mygram/helpers"
	"final-project-mygram/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := ctx.ShouldBind(&socialMedia); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	socialMedia.UserID = userID

	err := db.Debug().Create(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, socialMedia)

}

func GetAllSocialMedia(ctx *gin.Context) {
	var socialMedia []models.SocialMedia

	db := database.GetDB()

	if _, ok := ctx.GetQuery("user_id"); ok {
		user_id, err := strconv.Atoi(ctx.Query("user_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Input user_id with number",
			})
			return
		}

		err = db.Debug().Order("id").Where("user_id = ?", user_id).Find(&socialMedia).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if len(socialMedia) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("user_id %d doesn't have social media", user_id),
			})
			return
		}
	} else {
		err := db.Debug().Order("id").Find(&socialMedia).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": socialMedia,
	})
}

func GetOneSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	db := database.GetDB()

	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameter",
		})
		return
	}

	err = db.Debug().Where("id = ?", socialMediaID).First(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Social media not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func UpdateSocialMedia(ctx *gin.Context) {
	var socialMedia, findSocialMedia models.SocialMedia

	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	db := database.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := ctx.ShouldBind(&socialMedia); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	err = db.Debug().Where("id = ?", socialMediaID).First(&findSocialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Social media not found",
		})
		return
	}

	socialMedia = models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	socialMedia.ID = uint(socialMediaID)
	socialMedia.CreatedAt = findSocialMedia.CreatedAt
	socialMedia.UserID = findSocialMedia.UserID

	err = db.Debug().Model(&socialMedia).Where("id = ?", socialMediaID).Updates(socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func DeleteSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	db := database.GetDB()

	err = db.Debug().Where("id = ?", socialMediaID).First(&socialMedia).Delete(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Social Media %s successfully deleted", socialMedia.Name),
	})
}
