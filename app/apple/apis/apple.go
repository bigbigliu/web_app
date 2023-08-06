package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"github.com/bigbigliu/web_app/app/apple/models"
	"github.com/bigbigliu/web_app/app/apple/service"
	"github.com/bigbigliu/web_app/app/apple/service/dto"
)

type Apple struct {
	api.Api
}

// GetPage 获取Apple列表
// @Summary 获取Apple列表
// @Description 获取Apple列表
// @Tags Apple
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Apple}} "{"code": 200, "data": [...]}"
// @Router /api/v1/apple/list [get]
// @Security Bearer
func (e Apple) GetPage(c *gin.Context) {
    req := dto.AppleGetPageReq{}
    s := service.Apple{}
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

	list := make([]models.Apple, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Apple失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Apple
// @Summary 获取Apple
// @Description 获取Apple
// @Tags Apple
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Apple} "{"code": 200, "data": [...]}"
// @Router /api/v1/apple/{id} [get]
// @Security Bearer
func (e Apple) Get(c *gin.Context) {
	req := dto.AppleGetReq{}
	s := service.Apple{}
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
	var object models.Apple

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Apple失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Apple
// @Summary 创建Apple
// @Description 创建Apple
// @Tags Apple
// @Accept application/json
// @Product application/json
// @Param data body dto.AppleInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/apple/create [post]
// @Security Bearer
func (e Apple) Insert(c *gin.Context) {
    req := dto.AppleInsertReq{}
    s := service.Apple{}
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
		e.Error(500, err, fmt.Sprintf("创建Apple失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Apple
// @Summary 修改Apple
// @Description 修改Apple
// @Tags Apple
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AppleUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/apple/edit [put]
// @Security Bearer
func (e Apple) Update(c *gin.Context) {
    req := dto.AppleUpdateReq{}
    s := service.Apple{}
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

	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改Apple失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Apple
// @Summary 删除Apple
// @Description 删除Apple
// @Tags Apple
// @Param data body dto.AppleDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/apple [delete]
// @Security Bearer
func (e Apple) Delete(c *gin.Context) {
    s := service.Apple{}
    req := dto.AppleDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Apple失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
