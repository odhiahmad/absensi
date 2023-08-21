package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Absen struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID  uuid.UUID `gorm:"null"`
	User    *User     `gorm:"foreignKey:UserID;references:ID"`
	Nama     string    `gorm:"type:varchar(255)" json:"nama"`
	Lat       string     `gorm:"not null; size: 255" json:"lat"`
	Long                  string     `gorm:"not null; size: 255" json:"long"`
	CreatedAt             time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt             time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt             *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (u *Absen) Prepare() error {
	u.ID = uuid.NewV4()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
