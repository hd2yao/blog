package app

import (
    "time"

    "github.com/dgrijalva/jwt-go"

    "github.com/hd2yao/blog/global"
    "github.com/hd2yao/blog/pkg/util"
)

type Claims struct {
    AppKey    string `json:"app_key"`
    AppSecret string `json:"app_secret"`
    jwt.StandardClaims
}

// GetJWTSecret 获取该项目的 JWT Secret
func GetJWTSecret() []byte {
    return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成 JWT Token
func GenerateToken(appKey, appSecret string) (string, error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(global.JWTSetting.Expire)
    claims := Claims{
        AppKey:    util.EncodeMD5(appKey),
        AppSecret: util.EncodeMD5(appSecret),
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            Issuer:    global.JWTSetting.Issuer,
        },
    }

    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(GetJWTSecret())
    return token, err
}

// ParseToken 解析和校验 Token
func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return GetJWTSecret(), nil
    })
    if err != nil {
        return nil, err
    }
    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }
    return nil, err
}
