package persistence

import (
	"time"

	customgorm "github.com/tonybka/go-base-ddd/infrastructure/custom_gorm"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        customgorm.CustomTypeUUIDv1 `gorm:"primarykey;default:(UUID_TO_BIN(UUID()));"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
