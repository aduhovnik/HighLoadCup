package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id      int    `gorm:"primary_key" json:"id"`
	Email   string `json:"email"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Gender string `json:"gender"`
	Birth_date int `json:"birth_date""`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetId() int {
	return u.Id
}

func (u *User) GetRef() User{
	return *u
}

type Users_list struct {
	Users []User
}

func (u *User) CreateNew(db *gorm.DB) {
	db.Create(u)
}

func (u *User) GetTableInsert() string {
	return "INSERT INTO users (id, email, first_name, last_name, gender, birth_date) VALUES %s"
}

func (u * User) GetValueForInsert() string {
	return ""
}

func (u *Users_list) GetList() []Entity {
	return u.GetList()
}

/*func (c *Users_list) UnmarshalJSON(j []byte) error {
	//fmt.Print("ASDSA")
	err := json.Unmarshal(j, &c)
	if err != nil {
		return err
	}
	return nil
}*/