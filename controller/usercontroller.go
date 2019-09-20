package controller

import (
	"fondo-go/auth"
	"fondo-go/common"
	"fondo-go/config"
	"fondo-go/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func login(c *gin.Context) {

	type LoginRequest struct {
		Username string
		Password string
	}

	var (
		err          error
		loginRequest LoginRequest
	)
	err = c.ShouldBindJSON(&loginRequest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, common.Result{
			Code:      40001,
			Message:   "",
			Timestamp: time.Now().Unix(),
		})
	}

	user, err := db.FindUserByUsername(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Result{
			Code:      40001,
			Message:   "user not exist",
			Timestamp: time.Now().Unix(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusUnauthorized, common.Result{
				Code:      400101,
				Message:   "bad credential",
				Timestamp: time.Now().Unix(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, common.Result{
				Code:      500000,
				Message:   err.Error(),
				Timestamp: time.Now().Unix(),
			})
		}
		return
	}

	token, err := auth.SignToken(auth.Claims{User: &user}, time.Duration(config.Conf.Env.Token.ExpireIn),
		config.Conf.Env.Token.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"expires_in":   config.Conf.Env.Token.ExpireIn,
	})
}

func createUser(c *gin.Context) {
	result := common.Result{
		Code: 1,
		Data: db.User{
			Username: "username",
			Password: "password",
		},
		Message:   "success",
		Timestamp: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, result)
}

func findById(c *gin.Context) {
	result := common.Result{
		Code: 1,
		Data: db.User{
			Username: "username",
			Password: "password",
		},
		Message:   "success",
		Timestamp: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, result)
}

func UserController(router *gin.Engine) {
	router.POST("/user/login", login)
	router.POST("/users", createUser)
	router.GET("/users/:id", findById)
}
