package controllers

import (
	"jwt/config"
	"jwt/mosels"
	"net/http"
	"time"

	//"os/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	//get email and pass from req
	var body struct {
		Email string
		pass  string
	}
	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"eror": "faild to read body",
		})
		return
	}

	//hash pass

	hash, err := bcrypt.GenerateFromPassword([]byte(body.pass), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"eror": "faild to hash password",
		})
	}

	//createuser

	user := mosels.User{Email: body.Email, Password: string(hash)}
	resulet := config.DB.Create(&user)
	if resulet.Error != nil {
		c.JSON(400, gin.H{
			"eror": "faild to create user",
		})
		return
	}

	//respond
	c.JSON(200, gin.H{
		"id": user.ID,
	})
}

func Signin(c *gin.Context) {
	//get email from req
	var body struct {
		Email string
		pass  string
	}
	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"eror": "faild to read body",
		})
		return
	}
	//look fir req user

	var user mosels.User
	config.DB.First(&user, "Email= ?", body.Email)
	if user.ID == 0 {
		c.JSON(400, gin.H{
			"eror": "invaild emailor pass",
		})
		return
	}

	//compare pass

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.pass))
	if err != nil {
		c.JSON(400, gin.H{
			"eror": "invaild pass",
		})
		return
	}

	//generate token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("hmacSampleSecr02003404003030202et"))
	if err != nil {
		c.JSON(400, gin.H{
			"eror": "invaild create token",
		})
		return
	}

	//sent back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("author", tokenString, 3600*30*24, "", "", false, true)

	c.JSON(200, gin.H{})
}

func Addproduct(c *gin.Context) {
	var body struct {
		Name  string
		Price float64
	}
	userValue, _ := c.Get("id")
	id, _ := userValue.(uint)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}
	newproduct := mosels.Product{UserID: id, Name: body.Name, Price: body.Price}

	config.DB.Create(&newproduct)

	c.JSON(http.StatusOK, gin.H{

		"message": "Product added successfully",
	})
}

func Getproducts(c *gin.Context) {
	var products []mosels.Product
	config.DB.Find(&products)
	c.JSON(200, gin.H{
		"products": products,
	})

}

func GetUserProductes(c *gin.Context) {
	id := c.Param("id")
	var products []mosels.Product
	config.DB.Where("user_id = ?", id).Find(&products)
	c.JSON(200, gin.H{
		"products": products,
	})

}
