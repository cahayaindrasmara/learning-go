package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here
		// Ambil cookie session_token dari request
		sessionToken, err := ctx.Cookie("session_token")
		if err != nil {
			contentType := ctx.GetHeader("Content-Type")
			if contentType == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			} else {
				ctx.Redirect(http.StatusSeeOther, "/login")
			}
			ctx.Abort()
			return
		}

		// Parsing token menggunakan jwt-go
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(sessionToken, claims, func(token *jwt.Token) (interface{}, error) {
			// Ganti "your-secret-key" dengan kunci rahasia Anda
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing token"})
			}
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// Masukkan UserID ke dalam context
		ctx.Set("id", claims.UserID)

		// Lanjutkan ke handler berikutnya

		ctx.Next()
	})
}
