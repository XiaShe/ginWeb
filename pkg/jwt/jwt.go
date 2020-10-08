package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("侠奢侠奢你最帅")

/*
	MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
	jwt包自带的jwt.StandardClaims只包含了官方字段
	我们这里需要额外记录一个 userid 与 username 字段，所以要自定义结构体
	如果想要保存更多信息，都可以添加到这个结构体中
*/
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userid int64, username string) (string, error) {
	// 创建一个我们自己的声明的数据实例
	c := MyClaims{
		UserID:   userid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "My-Web",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	log.Println(token)
	log.Println(token.SignedString(mySecret))
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token /  将 tokenString 解析到 mc 中
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
