package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin   UserRole = "admin"
	RoleCustomer UserRole = "customer"
)

type User struct {
	gorm.Model
	BaseModel
	
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	SecretKey string            `bson:"secret_key"`
	Role      UserRole          `bson:"role"`
	CreatedAt time.Time         `bson:"created_at"`
}
