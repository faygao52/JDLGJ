package auth

import (
	"jdlgj/core"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//JWTClaims Custom claims
type JWTClaims struct {
	ID     uuid.UUID `json:"id"`
	Role   string    `json:"role"`
	Issuer string    `json:"issuer"`
	Expire int64     `json:"expire"`
	OpenID string    `json:"openid"`
	jwt.StandardClaims
}

//SignToken 生成token,uid用户id，expireSec过期秒数
func SignToken(id uuid.UUID, expireSec int, issuer string, role string, openid string) (tokenStr string, err error) {
	sec := time.Duration(expireSec)
	expire := time.Now().Add(time.Second * sec).Unix()
	claims := JWTClaims{ID: id, Role: role, Issuer: issuer, Expire: expire, OpenID: openid}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err = token.SignedString([]byte(core.GetEnv("SECRET", "secret")))

	return tokenStr, err
}

//Required check for authorisation
func Required() gin.HandlerFunc {
	return func(c *gin.Context) {

		authString := c.Request.Header.Get("Authorization")

		kv := strings.Split(authString, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Header without token"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := kv[1]

		// Parse token
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(core.GetEnv("SECRET", "secret")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Parse token failed"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Token is invalid"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(*JWTClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Token claim failed"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", claims)
	}
}
