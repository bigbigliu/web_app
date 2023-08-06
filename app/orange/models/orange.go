package models

import (

	"github.com/bigbigliu/web_app/common/models"

)

type Orange struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(255);comment:Name"` 
    Age string `json:"age" gorm:"type:int;comment:Age"` 
    models.ModelTime
    models.ControlBy
}

func (Orange) TableName() string {
    return "orange"
}

func (e *Orange) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Orange) GetId() interface{} {
	return e.Id
}