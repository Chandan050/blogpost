package controller

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Chandan050/blogwebsite/database"
	"github.com/Chandan050/blogwebsite/models"
	"github.com/Chandan050/blogwebsite/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// validation for email
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9. %+\-]+\.[a-z0-9.%+\-]`)

	return Re.MatchString(email)
}

// it will assign the user input data into database and it will validated
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]interface{}
		var userData models.User

		// user input data is binding to data map
		if err := c.ShouldBindJSON(&data); err != nil {
			fmt.Println("unable to parse body", err.Error())
			return
		}

		//Check if password is less than 6 charactres
		if len((data["password"]).(string)) <= 6 {

			c.JSON(http.StatusBadRequest, gin.H{"error": "passeord", "message": "Validation failed"})
			// return fmt.Errorf("error while len is than 6")
			return

		}
		if !ValidateEmail(strings.TrimSpace(data["email"].(string))) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email", "message": "invalid email address so failed"})
			// return fmt.Errorf("email is not perfect")
			return

		}
		//check if email already exist in database

		database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)

		if userData.Id != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email", "message": "invalid email address so failed"})
			// return fmt.Errorf("id already exits")
			return
		}

		user := models.User{
			Fname: data["fname"].(string),
			Lname: data["lname"].(string),
			Phone: data["phone"].(string),
			Email: strings.TrimSpace(data["email"].(string)),
		}

		user.SetPassword(data["password"].(string))
		err := database.DB.Create(&user)
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{"user": user, "message": "Account created succesfull"})
	}

	// return nil
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]string

		if err := c.ShouldBindJSON(&data); err != nil {
			fmt.Println("unable to parse the body")
		}

		var user models.User

		// it is search query to find user by using email id
		// select * from db where email = "db.email";

		database.DB.Where("email=?", data["email"]).First(&user)
		log.Println(data["email"])

		// if email id exists it will returns the is else it will return "0, nil - "
		if user.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not  exists", "message": "email adrres doesnt exits, kingly login an account"})
			return
		}

		// check password
		if err := user.ComparePassword(data["password"]); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "password", "message": "plz enter correct password"})
			return
		}

		token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "error")
		}

		c.SetCookie("jwt", token, int(time.Now().Add(time.Hour*24).Unix()), "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "user succesfully loged in", "user": user})
		c.Header("Authorization", token) // set jwt token in header

	}
}

type Claims struct {
	jwt.StandardClaims
}
