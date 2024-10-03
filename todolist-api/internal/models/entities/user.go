package entities

type User struct {
	Id        string `gorm:"column:id;primary_key;"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	Token     string `gorm:"column:token"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:mill"`
	Task      []Task `gorm:"foreignKey:id_user;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
