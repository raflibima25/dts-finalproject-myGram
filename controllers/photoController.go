package controllers

import (
	"final-project-mygram/database"
	"final-project-mygram/models"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(ctx *gin.Context) {
	var photo models.Photo

	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	username := userData["username"].(string)

	_ = ctx.ShouldBind(&photo)
	file, err := ctx.FormFile("photo_url")
	if err == nil {

		charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
		stringRandom := make([]byte, rand.Intn(100))
		for i := range stringRandom {
			stringRandom[i] = charset[rand.Intn(len(charset))]
		}

		ext := strings.Split(file.Filename, ".")[1]

		log.Println("file ext ->", ext)

		if ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "webp" {
			dst := "./assets/" + username + "-" + string(stringRandom) + "." + "jpg"
			ctx.SaveUploadedFile(file, dst)

			urlPhoto := "http://" + ctx.Request.Host + "/img/" + username + "-" + string(stringRandom) + "." + "jpg"
			photo.PhotoUrl = urlPhoto
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "File is not image/photo",
			})
			return
		}

	}

	photo.UserID = userID

	err = db.Debug().Create(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func GetAllPhoto(ctx *gin.Context) {
	var photos []models.Photo

	db := database.GetDB()

	if _, ok := ctx.GetQuery("user_id"); ok {
		user_id, err := strconv.Atoi(ctx.Query("user_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Input user_id with number",
			})
			return
		}

		err = db.Debug().Preload("Comments").Order("id").Where("user_id = ?", user_id).Find(&photos).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

	} else {

		err := db.Debug().Preload("Comments").Order("id").Find(&photos).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": photos,
	})
}

func GetOnePhoto(ctx *gin.Context) {
	var photo models.Photo

	db := database.GetDB()

	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Preload("Comments").Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func UpdatePhoto(ctx *gin.Context) {
	var photo, findPhoto models.Photo

	db := database.GetDB()

	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Where("id = ?", photoID).First(&findPhoto).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Photo not found",
		})
		return
	}

	_ = ctx.ShouldBind(&photo)

	photo = models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: findPhoto.PhotoUrl,
	}

	photo.ID = uint(photoID)
	photo.CreatedAt = findPhoto.CreatedAt
	photo.UserID = findPhoto.UserID

	err = db.Debug().Model(&photo).Where("id = ?", photoID).Updates(photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Preload("Comments").Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func DeletePhoto(ctx *gin.Context) {
	var photo models.Photo

	db := database.GetDB()

	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Where("id = ?", photoID).First(&photo).Delete(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Photo with title '%s' successfully deleted", photo.Title),
	})
}
