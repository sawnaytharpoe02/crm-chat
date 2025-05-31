package main

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	socketio "github.com/googollee/go-socket.io"
	"go.mongodb.org/mongo-driver/mongo"
)

// Configuration
const (
	AdminSecretKey    = "your_admin_secret_key_here"
	CustomerSecretKey = "your_customer_secret_key_here"
	UploadDir        = "./uploads"
	MaxUploadSize    = 5 << 20 // 5MB
)

var (
	mongoClient *mongo.Client
	userCollection *mongo.Collection
	messageCollection *mongo.Collection
)


func main() {

	app := fiber.New()

	// Socket.IO server setup
	server := socketio.NewServer(nil)

	// Authentication for Scoket.IO connections
	server.OnConnect("/", func(conn socketio.Conn) error {
		token := conn.URL().Query().Get("token")
		if token == "" {
			return errors.New("authorization token is required")
		}

		// Validate token and get role
		role, userID, err := validateTokenAndGetRole(token)
		if err != nil {
			return err
		}

		// Store user info in connection context
		conn.SetContext(map[string]interface{}{
			"role": role,
			"userID": userID,
		})

		// Join admin room if admin
		if role == "admin" {
			conn.Join("admin_room")
		}

		return nil
	})

	app.Listen(":3000")
}


func validateTokenAndGetRole(tokenString string) (string, string, error) {
	// First try with admin key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AdminSecretKey), nil
	})

	if err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, ok := claims["sub"].(string); ok {
				return "admin", sub, nil
			}
		}
	}

	// If admin key failed, try with customer key
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(CustomerSecretKey), nil
	})

	if err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, ok := claims["sub"].(string); ok {
				return "customer", sub, nil
			}
		}
	}

	return "", "", errors.New("invalid token")
}
