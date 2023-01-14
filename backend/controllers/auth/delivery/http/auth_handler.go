package handler

import (
	"net/http"

	"github.com/ariopri/Let-It-Be/tree/main/backend/models"
	"github.com/ariopri/Let-It-Be/tree/main/backend/utils/jwt"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authRepo models.AuthRepository
}

func NewAuthHandler(r *gin.Engine, loginRepo models.AuthRepository) {
	h := &authHandler{authRepo: loginRepo}
	a := r.Group("siswa")
	{
		a.POST("/login", h.login)
		a.POST("/register", h.register)
	}
}

func (a *authHandler) login(c *gin.Context) {
	ctx := c.Request.Context()
	var login models.LoginUser

	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error internal",
		})

		return
	}

	userLogin, err := a.authRepo.Login(ctx, &login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error internal",
		})

		return
	}

	generateToken, err := jwt.GenerateToken(userLogin.Email, userLogin.Role)
	if err != nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"message": "failed depedency",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"auth_token": generateToken,
		"data": gin.H{
			"email": userLogin.Email,
		},
	})
}

func (a *authHandler) register(c *gin.Context) {
	ctx := c.Request.Context()
	var registerUser models.RegisterUser

	err := c.ShouldBind(&registerUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error internal",
		})

		return
	}

	userRegist, err := a.authRepo.Register(ctx, &registerUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error internal",
		})

		return
	}

	generateToken, err := jwt.GenerateToken(userRegist.Email, userRegist.Role)
	if err != nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"message": "failed depedency",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"auth_token": generateToken,
		"data": gin.H{
			"email": userRegist.Email,
		},
	})
}
