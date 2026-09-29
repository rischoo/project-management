package models

import (
	"github.com/google/uuid"
	"github.com/rischoo/project-management/models/types"
)

type CardPosition struct {
	InternalID int64           `json:"internal_id" gorm:"primaryKey;AutoIncrement"`
	PublicID   uuid.UUID       `json:"public_id" gorm:"uuid;not null"`
	ListID     int64           `json:"list_internal_id" gorm:"column:list_internal_id";not null`
	CardOrder  types.UUIDArray `json:"card_order" gorm:"type:uuid[]"`
}
