package service

import (
	_ "api/util/jwt"
	jwt_secret "api/util/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var isPass = true
		var message string

		token := c.Request.Header.Get("authorization")
		if token == "" {
			message = "token not found"
			isPass = false
		} else {
			usersId, err := jwt_secret.ParseToken(token)
			c.Set("usersId", usersId)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					message = "token expired"
					isPass = false
				default:
					message = "token fail"
					isPass = false
				}
			}
		}

		if !isPass {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": message,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func WebSocketAuthMiddleware(isAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var isPass = c.IsWebsocket()
		var message string
		if isAuth {
			token := c.Request.Header.Get("authorization")
			if token == "" {
				message = "token not found"
				isPass = false
			} else {
				usersId, err := jwt_secret.ParseToken(token)
				c.Set("usersId", usersId)
				if err != nil {
					switch err.(*jwt.ValidationError).Errors {
					case jwt.ValidationErrorExpired:
						message = "token expired"
						isPass = false
					default:
						message = "token fail"
						isPass = false
					}
				}
			}
		}
		if !isPass {
			message = "Connection is invalid"
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": message,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
