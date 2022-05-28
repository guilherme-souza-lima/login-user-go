package model

type UserModel struct {
	ID        string `gorm:"PrimaryKey"`
	Name      string `gorm:"column:name"`
	Login     string `gorm:"unique"`
	Password  string `gorm:"column:password"`
	Email     string `gorm:"unique"`
	Cellphone string `gorm:"unique"`
}

func (ref *UserModel) TableName() string {
	return "users"
}
