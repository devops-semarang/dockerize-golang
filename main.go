package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv"
)

type User struct {
	gorm.Model
	UserId   string `gorm:"unique"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

func main() {
	// err := godotenv.Load(".env") // Load environment variables from a .env file
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
	}
	db.AutoMigrate(&User{})
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// Register User
	router.POST("/register", func(c *gin.Context) {
		// Generate a new UUID for the user ID
		id := uuid.New()

		// Bind the JSON data to the User struct
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Hash the user's password before storing it
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
			return
		}

		// Update the user's password with the hashed value
		user.UserId = id.String()
		user.Password = hashedPassword

		// Create the user in the database
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User Created Successfully!",
		})
	})
	// Login User
	router.POST("/login", func(c *gin.Context) {
		// Bind the JSON data to the User struct
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get the user from the database
		var dbUser User
		if err := db.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist"})
			return
		}

		// Check if the password is correct
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		jwt_token := generateJWT(dbUser.Username, dbUser.UserId)
		c.JSON(http.StatusOK, gin.H{
			"message":   "User Logged In Successfully!",
			"jwt_token": jwt_token,
		})
	})
	fmt.Print("Server running on port " + os.Getenv("APP_PORT"))
	router.Run(":" + os.Getenv("APP_PORT"))
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func generateJWT(username string, userId string) string {
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Id:        userId,
		Issuer:    username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt_token, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
	}
	return jwt_token
}
