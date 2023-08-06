package models

import (

	"github.com/bigbigliu/web_app/common/models"

)

type Apple struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:Name"` 
    Age string `json:"age" gorm:"type:int;comment:Age"` 
    models.ModelTime
    models.ControlBy
}

func (Apple) TableName() string {
    return "apple"
}

func (e *Apple) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Apple) GetId() interface{} {
	return e.Id
}