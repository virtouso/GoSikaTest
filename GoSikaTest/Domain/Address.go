package Domain

type Address struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	Address string
}
