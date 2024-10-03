package entities

type Task struct {
	Id          string `gorm:"column:id;primary_key"`
	IdUser      string `gorm:"column:id_user"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Status      string `gorm:"column:status"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime:mill"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoCreateTime:mill;autoUpdateTime:mill"`
	User        User   `gorm:"foreignKey:id_user;references:id"`
}

func (t *Task) TableName() string {
	return "tasks"
}
