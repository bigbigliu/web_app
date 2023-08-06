package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"github.com/bigbigliu/web_app/app/orange/models"
	"github.com/bigbigliu/web_app/app/orange/service/dto"
	cDto "github.com/bigbigliu/web_app/common/dto"
)

type Orange struct {
	service.Service
}

// GetPage 获取Orange列表
func (e *Orange) GetPage(c *dto.OrangeGetPageReq, list *[]models.Orange, count *int64) error {
	var err error
	var data models.Orange

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("OrangeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Orange对象
func (e *Orange) Get(d *dto.OrangeGetReq, model *models.Orange) error {
	var data models.Orange

	err := e.Orm.Model(&data).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在")
		e.Log.Errorf("Service GetOrange error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Orange对象
func (e *Orange) Insert(c *dto.OrangeInsertReq) error {
    var err error
    var data models.Orange
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("OrangeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Orange对象
func (e *Orange) Update(c *dto.OrangeUpdateReq) error {
    var err error
    var data = models.Orange{}
    e.Orm.First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("OrangeService Save error:%s \r\n", err)
        return err
    }
    return nil
}

// Remove 删除Orange
func (e *Orange) Remove(d *dto.OrangeDeleteReq) error {
	var data models.Orange

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveOrange error:%s \r\n", err)
        return err
    }
	return nil
}
