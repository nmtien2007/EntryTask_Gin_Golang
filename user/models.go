package user

// using function to set table name is not plural
//func (User) TableName() string {
//	return "user"
//}

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Company struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Users []User `gorm:"foreignKey:CompanyId"`
}

type User struct {
	Id        int32     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CompanyId int       `json:"company_id" gorm:"column:company_id"`
	Company   Company   `json:"company"`
	Groups    []*Group  `json:"groups" gorm:"many2many:user_group;"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

type Group struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Users        []*User        `gorm:"many2many:user_group;" json:"users"`
	Perms []*Perm `gorm:"many2many:group_perm;" json:"permissions"`
}

type Perm struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Groups    []*Group  `gorm:"many2many:group_perm;" json:"groups"`
}
