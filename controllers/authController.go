package controllers

import (
	"jdlgj/auth"
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
		token, err := auth.SignWxToken(userResource.ID, 86400, "wechat", "user", userResource.OpenID)
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

//Login login user
// func Login(c *gin.Context) {
// 	c.JSON(http.StatusOK)

// }
