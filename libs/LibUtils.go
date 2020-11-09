package libs

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var countGame uint64 = 0
var countUser uint64 = 0

func GenerateId(count *uint64) string{
	s := strconv.FormatUint(*count,10) + time.Now().String()
	*count++
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	id := fmt.Sprintf("%x", bs)
	return id
}

func GenerateGameId() string{
	return GenerateId(&countGame)
}
func GenerateUserId() string {
	return GenerateId(&countUser)
}

func GenerateGameToken(idUser string, idGame string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"idUser": idUser,
		"idGame": idGame,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("networkEngine"))
	return tokenString, err
}
