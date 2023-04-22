package controllers

import (
	"final-project-mygram/database"
	"final-project-mygram/helpers"
	"final-project-mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJson = "application/json"

// RegisterUser godoc
// @Summary Register User
// @Description Register user for my gram
// @Tags User
// @Accept json
// @Produce json
// @Param UserRegister body models.RequestUserRegister true "User Register"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ResponseFailed
// @Router /user/register [post]
func RegisterUser(ctx *gin.Context) {
	var user models.User

	db := database.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"age":      user.Age,
		"email":    user.Email,
		"username": user.Username,
	})
}

// LoginUser godoc
// @Summary Login User
// @Description Login user for have token (jwt)
// @Tags User
// @Accept json
// @Produce json
// @Param UserLogin body models.RequestUserLogin true "User Login"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailed
// @Router /user/login [post]
func LoginUser(ctx *gin.Context) {
	var user models.User

	db := database.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	}

	password := user.Password

	err := db.Debug().Where("username = ?", user.Username).Take(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
