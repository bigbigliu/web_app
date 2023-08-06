package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"github.com/bigbigliu/web_app/app/apple/models"
	"github.com/bigbigliu/web_app/app/apple/service/dto"
	cDto "github.com/bigbigliu/web_app/common/dto"
)

type Apple struct {
	service.Service
}

// GetPage 获取Apple列表
func (e *Apple) GetPage(c *dto.AppleGetPageReq, list *[]models.Apple, count *int64) error {
	var err error
	var data models.Apple

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AppleService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Apple对象
func (e *Apple) Get(d *dto.AppleGetReq, model *models.Apple) error {
	var data models.Apple

	err := e.Orm.Model(&data).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在")
		e.Log.Errorf("Service GetApple error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Apple对象
func (e *Apple) Insert(c *dto.AppleInsertReq) error {
    var err error
    var data models.Apple
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AppleService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Apple对象
func (e *Apple) Update(c *dto.AppleUpdateReq) error {
    var err error
    var data = models.Apple{}
    e.Orm.First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("AppleService Save error:%s \r\n", err)
        return err
    }
    return nil
}

// Remove 删除Apple
func (e *Apple) Remove(d *dto.AppleDeleteReq) error {
	var data models.Apple

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveApple error:%s \r\n", err)
        return err
    }
	return nil
}
