package dto

import (

	"github.com/bigbigliu/web_app/app/orange/models"
	"github.com/bigbigliu/web_app/common/dto"
	common "github.com/bigbigliu/web_app/common/models"
)

type OrangeGetPageReq struct {
	dto.Pagination     `search:"-"`
    OrangeOrder
}

type OrangeOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:orange"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:orange"`
    Age string `form:"ageOrder"  search:"type:order;column:age;table:orange"`
    
}

func (m *OrangeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrangeInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:""`
    Age string `json:"age" comment:""`
    common.ControlBy
}

func (s *OrangeInsertReq) Generate(model *models.Orange)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Age = s.Age
}

func (s *OrangeInsertReq) GetId() interface{} {
	return s.Id
}

type OrangeUpdateReq struct {
    Id int `json:"id" comment:""` // 
    Name string `json:"name" comment:""`
    Age string `json:"age" comment:""`
    common.ControlBy
}

func (s *OrangeUpdateReq) Generate(model *models.Orange)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Age = s.Age
}

func (s *OrangeUpdateReq) GetId() interface{} {
	return s.Id
}

// OrangeGetReq 功能获取请求参数
type OrangeGetReq struct {
     Id int `uri:"id"`
}
func (s *OrangeGetReq) GetId() interface{} {
	return s.Id
}

// OrangeDeleteReq 功能删除请求参数
type OrangeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrangeDeleteReq) GetId() interface{} {
	return s.Ids
}
