package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	InternalID int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;AutoIncrement"`
	PublicID   uuid.UUID `json:"public_id" db:"public_id"`
	CardID     int64     `json:"card_internal_id" db:"card_internal_id" `
	CardPubId  uuid.UUID `json:"card_id" db:"card_id""`
	UserID     int64     `json:"user_id" db:"user_id" gorm:"not null"`
	UserPubID  uuid.UUID `json:"user_pub_id" db:"user_pub_id" gorm:"type:uuid;not null"`
	Message    string    `json:"message" db:"message" gorm:"type:text;not null"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
