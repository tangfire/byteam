package app

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type adminClaims struct {
	AdminID  uint   `json:"adminId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *Server) ensureAdmin() error {
	var count int64
	if err := s.db.Model(&Admin{}).Where("username = ?", s.cfg.AdminUsername).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(s.cfg.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.db.Create(&Admin{
		Username:     s.cfg.AdminUsername,
		PasswordHash: string(hash),
	}).Error
}

func (s *Server) login(c *gin.Context) {
	payload, ok := bindJSON[loginRequest](c)
	if !ok {
		return
	}

	var admin Admin
	err := s.db.Where("username = ?", payload.Username).First(&admin).Error
	if err != nil || bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(payload.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token, err := s.signToken(admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"admin": gin.H{
			"id":       admin.ID,
			"username": admin.Username,
		},
	})
}

func (s *Server) me(c *gin.Context) {
	admin := currentAdmin(c)
	c.JSON(http.StatusOK, gin.H{"admin": admin})
}

func (s *Server) signToken(admin Admin) (string, error) {
	now := time.Now()
	claims := adminClaims{
		AdminID:  admin.ID,
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   admin.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.cfg.TokenTTL)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.cfg.JWTSecret))
}

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimSpace(strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer "))
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		var claims adminClaims
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(_ *jwt.Token) (any, error) {
			return []byte(s.cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		var admin Admin
		if err := s.db.First(&admin, claims.AdminID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "admin not found"})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Set("admin", gin.H{"id": admin.ID, "username": admin.Username})
		c.Next()
	}
}

func currentAdmin(c *gin.Context) any {
	admin, exists := c.Get("admin")
	if !exists {
		return nil
	}
	return admin
}
