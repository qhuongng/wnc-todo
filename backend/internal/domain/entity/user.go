package entity

// mapping for database
type User struct {
	Id          int64  `gorm:"column:id;primaryKey" json:"id"`
	Username    string `gorm:"column:username" json:"username"`
	Password    string `gorm:"column:password" json:"password"`
	RefeshToken string `gorm:"column:refresh_token" json:"refresh_token"`
}

func (User) TableName() string {
	return "users"
}
