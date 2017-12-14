package models

import "encoding/json"

import "github.com/jinzhu/gorm"

type Entity interface {
	GetId() int
	CreateNew(db *gorm.DB)
	GetTableInsert() string
	GetValueForInsert() string
}

type Entity_list struct {
	EntityList []Entity
}

type EntityListI interface {

}

type EntityList interface {
	json.Unmarshaler
	GetList() []Entity
}

func JsonDecode(jsonStr []byte, v EntityList) error {
	return json.Unmarshal(jsonStr, v)
}
