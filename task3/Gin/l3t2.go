package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func getDSN() string {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	return user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:user"` // admin / user
}

var jwtSecret = []byte("secret-key")

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func generateToken(userID uint, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid Authorization header"})
			c.Abort()
			return
		}
		tokenStr := parts[1]
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		claims := token.Claims.(*Claims)
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, exists := c.Get("role")
		if !exists || r != role {
			c.JSON(403, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}


func loginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var user User
		if err := db.Where("username=? AND password=?", req.Username, req.Password).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		token, _ := generateToken(user.ID, user.Role)
		c.JSON(200, gin.H{"token": token, "user": gin.H{"id": user.ID, "username": user.Username, "role": user.Role}})
	}
}

func getUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, users)
	}
}

func createUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&user)
		c.JSON(200, user)
	}
}

var allowedExts = []string{".jpg", ".png", ".pdf"}
var maxFileSize int64 = 5 << 20 // 5MB

func uploadFileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if file.Size > maxFileSize {
			c.JSON(400, gin.H{"error": "File too large"})
			return
		}

		ext := strings.ToLower(filepath.Ext(file.Filename))
		valid := false
		for _, e := range allowedExts {
			if e == ext {
				valid = true
				break
			}
		}
		if !valid {
			c.JSON(400, gin.H{"error": "Invalid file type"})
			return
		}

		dst := filepath.Join("uploads", filepath.Base(file.Filename))
		c.SaveUploadedFile(file, dst)
		c.JSON(200, gin.H{"message": "File uploaded", "filename": file.Filename})
	}
}

func main() {
	os.MkdirAll("./uploads", 0755)

	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}
	db.AutoMigrate(&User{})

	r := gin.Default()

	r.POST("/login", loginHandler(db))

	r.POST("/upload", uploadFileHandler())

	api := r.Group("/api")
	api.Use(AuthMiddleware(), RequireRole("admin"))
	{
		api.GET("/users", getUsersHandler(db))
		api.POST("/users", createUserHandler(db))
	}

	r.Run(":8080")
}
