package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("super-secret-key")
var APIKEY = "12345"

/*func Load() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	APIKEY = os.Getenv("API_KEY")
	SECRET = []byte(os.Getenv("SECRET_KEY"))
}*/

func CreateJWT(role string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["role"] = role

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return tokenStr, nil
}

func ValidateJwtOnHandler(roles []string, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return SECRET, nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized"))
			}

			if token.Valid {
				for _, role := range roles {
					if token.Claims.(jwt.MapClaims)["role"] == role {
						handler(w, r)
						return
					}
				}
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

func ValidateJWT(roles []string, inputToken string) bool {

	type Claims struct {
		Roles []string `json:"roles"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(inputToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})
	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {

		for _, role := range roles {
			if claims.Roles[0] == role {
				return true
			}
		}
	}
	return false
}

func GetJWT(w http.ResponseWriter, r *http.Request) {

	if r.Header["Role"] != nil {
		if r.Header["Access"] != nil {
			if r.Header["Access"][0] != APIKEY {
				fmt.Fprint(w, "APIKEY")
				return
			} else {
				token, err := CreateJWT(r.Header["Role"][0])
				if err != nil {
					return
				}
				fmt.Fprint(w, token)
			}
		}
	} else {
		fmt.Fprint(w, "role needed")
	}
}
