package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserClick struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID        string             `json:"user_id" bson:"user_id"`
	WalletAddress string             `json:"wallet_address" bson:"wallet_address"`
	Date          string             `json:"date" bson:"date"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserClickRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
type UserClickResponse struct {
	Status bool `json:"status"`
}
