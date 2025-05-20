package models

import "time"

const (
	BUYER  = "buyer"
	SELLER = "seller"
)

type User struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      int       `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Address   Address   `json:"address"`
	Cart      Cart      `json:"cart"`
	Orders    []Order   `json:"orders"`
	Payment   []Payment `json:"payment"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"user_type" gorm:"default:buyer"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn  time.Time `json:"voided_on" gorm:"default:NULL"`
}
