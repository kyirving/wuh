package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

// JWTProvider JWT 提供器
type JWTProvider struct {
	secret        []byte        // 密钥
	issuer        string        // 签发者
	expire        time.Duration // 过期时间
	refreshExpire time.Duration // 刷新过期时间
}

type JWTToken struct {
	Token         string `json:"token"`
	Expire        int64  `json:"expire"`
	RefreshToken  string `json:"refreshToken"`
	RefreshExpire int64  `json:"refreshExpire"`
}

// NewJWTProvider 创建 JWT 提供器
func NewJWTProvider(secret []byte, issuer string, expire time.Duration, refreshExpire time.Duration) *JWTProvider {
	return &JWTProvider{
		secret:        secret,
		issuer:        issuer,
		expire:        expire,
		refreshExpire: refreshExpire,
	}
}

// GenerateToken 生成 JWT 令牌
func (j *JWTProvider) GenerateToken(userId int64, roles []string) (*JWTToken, error) {
	claims := jwt.MapClaims{
		"sub":   userId,
		"iss":   j.issuer,
		"roles": roles,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(j.expire).Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"sub": userId,
		"iss": j.issuer,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(j.refreshExpire).Unix(),
	}

	Token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := Token1.SignedString(j.secret)
	if err != nil {
		return nil, err
	}

	Token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := Token2.SignedString(j.secret)
	if err != nil {
		return nil, err
	}

	return &JWTToken{
		Token:         token,
		Expire:        cast.ToInt64(claims["exp"]),
		RefreshToken:  refreshToken,
		RefreshExpire: cast.ToInt64(refreshClaims["exp"]),
	}, nil
}

// ParseToken 解析 JWT 令牌
func (j *JWTProvider) ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	}, jwt.WithIssuer(j.issuer))
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// RefreshToken 刷新 JWT 令牌
func (j *JWTProvider) RefreshToken(tokenString string) (*JWTToken, error) {
	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	userId, ok := claims["sub"].(float64)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	roles, ok := claims["roles"].([]interface{})
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// 生成新的令牌
	newToken, err := j.GenerateToken(int64(userId), cast.ToStringSlice(roles))
	if err != nil {
		return nil, err
	}

	return newToken, nil
}
