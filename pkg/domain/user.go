package domain

type User struct {
	Id         uint64 `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Firstname  string `json:"firstname" validate:"required,min=4,max=16"`
	Lastname   string `json:"lastname" validate:"required,min=4,max=16"`
	Password   string `json:"password" validate:"required,min=6,max=16"`
	Email      string `json:"email" validate:"email,required"`
	Otp        string `json:"otp"`
	Isverified bool   `json:"isverified" gorm:"default:false"`
	Profile    string `json:"profile"`
}
