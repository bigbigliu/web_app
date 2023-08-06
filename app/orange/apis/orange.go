package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"github.com/bigbigliu/web_app/app/orange/models"
	"github.com/bigbigliu/web_app/app/orange/service"
	"github.com/bigbigliu/web_app/app/orange/service/dto"
)

type Orange struct {
	api.Api
}

// GetPage 获取Orange列表
// @Summary 获取Orange列表
// @Description 获取Orange列表
// @Tags Orange
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Orange}} "{"code": 200, "data": [...]}"
// @Router /api/v1/orange/list [get]
// @Security Bearer
func (e Orange) GetPage(c *gin.Context) {
    req := dto.OrangeGetPageReq{}
    s := service.Orange{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	list := make([]models.Orange, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Orange失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Orange
// @Summary 获取Orange
// @Description 获取Orange
// @Tags Orange
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Orange} "{"code": 200, "data": [...]}"
// @Router /api/v1/orange/{id} [get]
// @Security Bearer
func (e Orange) Get(c *gin.Context) {
	req := dto.OrangeGetReq{}
	s := service.Orange{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Orange

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Orange失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Orange
// @Summary 创建Orange
// @Description 创建Orange
// @Tags Orange
// @Accept application/json
// @Product application/json
// @Param data body dto.OrangeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/orange/create [post]
// @Security Bearer
func (e Orange) Insert(c *gin.Context) {
    req := dto.OrangeInsertReq{}
    s := service.Orange{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建Orange失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Orange
// @Summary 修改Orange
// @Description 修改Orange
// @Tags Orange
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrangeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/orange/{id}/edit [put]
// @Security Bearer
func (e Orange) Update(c *gin.Context) {
    req := dto.OrangeUpdateReq{}
    s := service.Orange{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))

	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改Orange失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Orange
// @Summary 删除Orange
// @Description 删除Orange
// @Tags Orange
// @Param data body dto.OrangeDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/orange [delete]
// @Security Bearer
func (e Orange) Delete(c *gin.Context) {
    s := service.Orange{}
    req := dto.OrangeDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除Orange失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
