package token

import (
	"errors"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	_secret        = "lplms-token-secret"
	_issuer        = "lplms-token-issuer"
	ExpireDuration = time.Hour * 4
)

var AuthSecret = []byte(_secret)

type AuthClaims struct {
	User AuthUser
	jwt.StandardClaims
}

type AuthUser struct {
	UID       string
	ProjectID string
}

func GenerateToken(u *entity.User) (string, error) {
	if u == nil {
		return "", errors.New("empty user")
	}
	expirationTime := time.Now().Add(ExpireDuration)
	claims := &AuthClaims{
		User: AuthUser{
			UID:       u.UID,
			ProjectID: u.ProjectID,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    _issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenStr, err := token.SignedString(AuthSecret); err != nil {
		return "", err
	} else {
		return tokenStr, nil
	}
}

func ParseToken(token string) (*AuthClaims, error) {
	tk, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		return AuthSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tk.Claims.(*AuthClaims); ok && tk.Valid {
		return claims, nil
	}
	return nil, errors.New("err parse token")
}
