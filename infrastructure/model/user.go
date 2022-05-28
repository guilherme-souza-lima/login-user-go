package model

type UserModel struct {
	ID        string `gorm:"PrimaryKey"`
	Name      string `gorm:"column:name"`
	Login     string `gorm:"column:login"`
	Password  string `gorm:"column:password"`
	Email     string `gorm:"column:email"`
	Cellphone string `gorm:"column:cellphone"`
}

func (ref *UserModel) TableName() string {
	return "users"
}
