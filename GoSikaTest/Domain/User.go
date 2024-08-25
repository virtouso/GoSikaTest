package Domain

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Addresses []Address `gorm:"foreignKey:UserID"` // One-to-many relationship
	CreatedAt time.Time
}
