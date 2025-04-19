package model

type User struct {
	UserId   int64  `json:"user_id" gorm:"column:user_id;"`
	Name     string `json:"name" gorm:"column:name;"`
	Email    string `json:"email" gorm:"column:email;"`
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (User) TableName() string {
	return "users"
}
