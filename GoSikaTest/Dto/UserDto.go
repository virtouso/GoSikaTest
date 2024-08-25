package Dto

import (
	"time"
)

type UserBulk struct {
	Users []UserDto
}

type UserDto struct {
	ID        uint
	Name      string
	Email     string
	Addresses []AddressDto
	CreatedAt time.Time
}

type AddressDto struct {
	ID      uint
	Address string
}
