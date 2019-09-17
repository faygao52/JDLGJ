package controllers

import (
	"jdlgj/auth"
	"jdlgj/core"
	"jdlgj/models"
	"jdlgj/repository"

	"net/http"

	"github.com/gin-gonic/gin"
)

//LoginByWechat login user by wechat code
func LoginByWechat(c *gin.Context) {
	var user models.WcUser

	code := c.Query("code")
	wxSession, wxErr := models.WxLogin(code)
	if wxErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Wechat login failed"})
		return
	}

	resource, err := repository.FindBy("open_id", &user, wxSession.OpenID)
	if err != nil {
		user = models.WcUser{OpenID: wxSession.OpenID}
		resource = repository.Create(&user)
	}
	userResource, ok := resource.(models.WcUserResource)
	if ok {
		token, err := auth.SignToken(userResource.ID, 86400, "wechat", "user", userResource.OpenID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "JWT token signed failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": token, "openId": userResource.OpenID})

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Type convert error"})
		return
	}
}

//Register new user
func Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingUser models.User
	resource, err := repository.FindBy("username", &existingUser, user.Username)
	if err != nil {
		var password = core.HashAndSalt(user.Password)
		resource = repository.Create(&models.User{Username: user.Username, Name: user.Name, Password: password, Role: user.Role})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "message": "Username already exist"})
		return
	}
	userResource, ok := resource.(models.UserResource)
	if ok {
		token, err := auth.SignToken(userResource.ID, 86400, "react", userResource.Role, "")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "JWT token signed failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user": userResource, "token": token})

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Type convert error"})
		return
	}
}

// LoginJSON stuff
type LoginJSON struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//Login let user login via dashboard
func Login(c *gin.Context) {
	var user models.User
	var loginForm LoginJSON

	if err := c.BindJSON(&loginForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resource, err := repository.FindBy("username", &user, loginForm.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusNotFound, "message": "Username doesn't exist"})
		return
	}
	userResource, ok := resource.(models.UserResource)
	if ok {
		if !core.ComparePasswords(user.Password, loginForm.Password) {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusUnauthorized, "message": "Incorrect password"})
			return
		}

		token, err := auth.SignToken(userResource.ID, 86400, "react", userResource.Role, "")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "JWT token signed failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user": userResource, "token": token})

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Type convert error"})
		return
	}
}
