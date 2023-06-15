package models

type User struct {
	Id            int          `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name          string       `gorm:"column:name" json:"name"`
	Email         string       `gorm:"column:email" json:"email"`
	Password      string       `gorm:"column:password" json:"password"`
	Phone         string       `gorm:"column:phone" json:"phone"`
	EmailVerified bool         `gorm:"column:email_verified" json:"email_verified"`
	PhoneVerified bool         `gorm:"column:phone_verified" json:"phone_verified"`
	Role          string       `gorm:"column:role" json:"role"`
	ApplicationId int          `gorm:"column:application_id; default:null" json:"application_id"`
	IsActive      bool         `gorm:"column:is_active; default:true" json:"is_active"`
	Application   *Application `gorm:"foreignKey:application_id;AssociationForeignKey:id" json:"application"`
}

func (User) TableName() string {
	return "users"
}
