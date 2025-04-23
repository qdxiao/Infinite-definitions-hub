package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"infni-backend/apps/api/internal/config"
)

// GenerateToken 生成JWT token
func GenerateToken(userId int64, username string, cfg config.Config) (string, error) {
	// 设置过期时间为7天
	expirationTime := time.Now().Add(time.Duration(cfg.Jwt.Expire) * time.Second)

	// 创建claims
	claims := jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	// 使用JWT secret签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.Jwt.Secret))
}
