package models


type Visit struct {
	Id      int  `gorm:"primary_key"json:"id"`
	Location int `json:"location"`
	User int `json:"user"`
	Visited_at int `json:"visited_at"`
	Mark int `json:"mark"`
}

type Visits_list struct {
	Visits []Visit
}

func (u *Visit) TableName() string {
	return "visits"
}
