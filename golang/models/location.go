package models


type Location struct {
	Id      int  `gorm:"primary_key"json:"id"`
	Place string `gorm:"type:text" json:"place""`
	Country string `json:"country"`
	City string `json:"city"`
	Distance int `json:"distance"`
}

func (u *Location) TableName() string {
	return "locations"
}

type Locations_list struct {
	Locations []Location
}

