package models

import "time"

type User struct {
	IdUser            uint `gorm:"primaryKey" json:"id_user"`
	Fullname          string `json:"fullname" valid:"type(string), required"`
	Email             string `json:"email" valid:"type(string), required"`
	Password          string `json:"password" valid:"type(string), required"`
	Address           string `json:"address" valid:"type(string), required"`
	Birthdate         string `json:"birthdate" valid:"type(string), required"`
	Phone             string `json:"phone" valid:"type(string), required"`
	Gender            string `json:"gender" valid:"type(string), required"`
	Image             string `json:"image" valid:"type(string), required"`
	Role              string `json:"role" valid:"type(string), required"`
	Verified          bool `json:"verified" valid:"type(bool), required"`
	Created_at        time.Time
	Updated_at        time.Time
	Reset_Pass_Token  string `json:"reset_pass_token" valid:"type(string), required"`
	Reset_Pass_Expire string `json:"reset_pass_expire" valid:"type(string), required"`
}

type Users []User
