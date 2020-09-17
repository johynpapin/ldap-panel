package main

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
	"log"
	"net/http"
	"os"
)

type signInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type updatePasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

type user struct {
	Username string `json:"username"`
}

func main() {
	// LDAP

	ldapURL := os.Getenv("LDAP_PANEL_LDAP_URL")
	baseDN := os.Getenv("LDAP_PANEL_LDAP_BASE_DN")

	// Auth

	key := os.Getenv("LDAP_PANEL_SECRET_KEY")

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "ldap-panel",
		Key:   []byte(key),
		Authenticator: func(ctx *gin.Context) (interface{}, error) {
			var request signInRequest
			if err := ctx.ShouldBind(&request); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			l, err := ldap.DialURL(ldapURL)
			if err != nil {
				log.Printf("ldap.DialURL() Error: %v\n", err.Error())
				return nil, jwt.ErrFailedAuthentication
			}

			err = l.Bind(fmt.Sprintf("uid=%s,%s", request.Username, baseDN), request.Password)
			if err != nil {
				log.Printf("l.Bind() Error: %v\n", err.Error())
				return nil, jwt.ErrFailedAuthentication
			}

			return &user{
				Username: request.Username,
			}, nil
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user); ok {
				return jwt.MapClaims{
					"identity": v.Username,
				}
			}

			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			log.Println("claims")
			log.Println(claims)

			return &user{
				Username: claims["identity"].(string),
			}
		},
	})
	if err != nil {
		log.Fatalf("JWT Error: %v\n", err.Error())
	}

	if authMiddleware.MiddlewareInit() != nil {
		log.Fatalf("authMiddleware.MiddlewareInit() Error: %v\n", err.Error())
	}

	// Routes

	router := gin.Default()

	router.POST("/auth/sign-in", authMiddleware.LoginHandler)
	router.GET("/auth/refresh", authMiddleware.MiddlewareFunc(), authMiddleware.RefreshHandler)

	router.POST("/update-password", authMiddleware.MiddlewareFunc(), func(ctx *gin.Context) {
		var request updatePasswordRequest
		if err := ctx.ShouldBind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		rawUser, _ := ctx.Get("identity")
		user := rawUser.(*user)

		l, err := ldap.DialURL(ldapURL)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = l.Bind(os.Getenv("LDAP_PANEL_LDAP_USERNAME"), os.Getenv("LDAP_PANEL_LDAP_PASSWORD"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		passwordModifyRequest := ldap.NewPasswordModifyRequest(
			fmt.Sprintf("uid=%s,%s", user.Username, baseDN), "", request.Password)
		if _, err = l.PasswordModify(passwordModifyRequest); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	})

	router.Run()
}
