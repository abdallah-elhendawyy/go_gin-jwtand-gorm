package midleware

import (
	"fmt"
	"jwt/config"
	"jwt/mosels"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ReqAuth(c *gin.Context) {
	fmt.Println("iam in middle")

	//get cokiee

	tokenString, err := c.Cookie("author")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode//validate it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("hmacSampleSecr02003404003030202et"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//check expiry

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user of token

		var user mosels.User
		config.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//atach the req
		c.Set("id", user.ID)

		//continue
		c.Next()

		fmt.Println(claims["foo"], claims["nbf"])

	} else {
		fmt.Println(err)
	}

}
