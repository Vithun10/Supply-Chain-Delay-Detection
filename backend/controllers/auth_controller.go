package controllers

import (
	"net/http"
	"supply-chain-monitor/services"

	"github.com/gin-gonic/gin"
)

// AuthController handles register and login HTTP requests.
type AuthController struct {
	authSvc *services.AuthService
}

// NewAuthController creates a new AuthController.
func NewAuthController(authSvc *services.AuthService) *AuthController {
	return &AuthController{authSvc: authSvc}
}

// Register handles POST /register
func (ctrl *AuthController) Register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if body.Username == "" || body.Password == "" || body.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username, password and role are required"})
		return
	}

	err := ctrl.authSvc.Register(body.Username, body.Password, body.Role)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

// Login handles POST /login
func (ctrl *AuthController) Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if body.Username == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	token, err := ctrl.authSvc.Login(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}
