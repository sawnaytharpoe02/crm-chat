package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	BaseModel

	ID        primitive.ObjectID `bson:"_id,omitempty"`
	SenderID  primitive.ObjectID `bson:"sender_id"`
	Content   string             `bson:"content"`
	FilePath  string             `bson:"file_path,omitempty"`
	IsFile    bool               `bson:"is_file"`
	IsBroadcast bool             `bson:"is_broadcast"`
	Recipients []primitive.ObjectID `bson:"recipients,omitempty"`
}
