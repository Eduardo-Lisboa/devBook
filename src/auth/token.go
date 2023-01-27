package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(UsuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["UsuarioID"] = UsuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("secret"))
}
