package libs

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"github.com/l-giuliani/networkEngine/context"
	"github.com/l-giuliani/networkEngine/components"
)

var countGame uint64 = 0
var countUser uint64 = 0

func GetGamePool() *components.GamePool{
	return context.GetGamePool()
}

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

func ParseGameToken(tokenString string) (string, string, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("networkEngine"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["idUser"].(string), claims["idGame"].(string), nil
	} else {
		return "", "", err
	}
}