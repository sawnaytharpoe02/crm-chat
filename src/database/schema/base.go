package schema

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time			`bson:"updated_at"`
	DeletedAt gorm.DeletedAt	`bson:"deleted_at,omitempty"`
	Version int 				`bson:"version,default:1"`
}
