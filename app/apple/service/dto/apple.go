package dto

import (

	"github.com/bigbigliu/web_app/app/apple/models"
	"github.com/bigbigliu/web_app/common/dto"
	common "github.com/bigbigliu/web_app/common/models"
)

type AppleGetPageReq struct {
	dto.Pagination     `search:"-"`
    AppleOrder
}

type AppleOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:apple"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:apple"`
    Age string `form:"ageOrder"  search:"type:order;column:age;table:apple"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:apple"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:apple"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:apple"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:apple"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:apple"`
    
}

func (m *AppleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AppleInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:""`
    Age string `json:"age" comment:""`
    common.ControlBy
}

func (s *AppleInsertReq) Generate(model *models.Apple)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Age = s.Age
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AppleInsertReq) GetId() interface{} {
	return s.Id
}

type AppleUpdateReq struct {
    Id int `json:"id" comment:""` // 
    Name string `json:"name" comment:""`
    Age string `json:"age" comment:""`
    common.ControlBy
}

func (s *AppleUpdateReq) Generate(model *models.Apple)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Age = s.Age
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AppleUpdateReq) GetId() interface{} {
	return s.Id
}

// AppleGetReq 功能获取请求参数
type AppleGetReq struct {
     Id int `uri:"id"`
}
func (s *AppleGetReq) GetId() interface{} {
	return s.Id
}

// AppleDeleteReq 功能删除请求参数
type AppleDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AppleDeleteReq) GetId() interface{} {
	return s.Ids
}
