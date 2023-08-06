package tools

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"github.com/bigbigliu/web_app/app/other/models/tools"
)

type Gen struct {
	api.Api
}

func (e Gen) GenCodeNew(c *gin.Context) {
	e.Context = c
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(500, err, "数据库连接获取失败")
		return
	}
	table_name := c.Query("table")

	tablesList := strings.Split(table_name, ",")
	if len(tablesList) != 1 {
		e.Error(400, nil, "一次只能生成一张表的代码")
		return
	}

	for i := 0; i < len(tablesList); i++ {

		data, err := genTableInit(db, tablesList, i, c)
		if err != nil {
			log.Errorf("genTableInit error, %s", err.Error())
			e.Error(500, err, "")
			return
		}

		_, err = data.Create(db)
		if err != nil {
			log.Errorf("Create error, %s", err.Error())
			e.Error(500, err, "")
			return
		}
	}

	table := tools.SysTables{}

	table.TBName = table_name
	tab, _ := table.Get(db, false)

	e.NOActionsGen(c, tab)

	e.OK("", "Code generated successfully！")
}

func (e Gen) NOActionsGen(c *gin.Context, tab tools.SysTables) {
	e.Context = c
	log := e.GetLogger()
	tab.MLTBName = strings.Replace(tab.TBName, "_", "-", -1)

	basePath := "template/v4/"
	routerFile := ""

	switch tab.IsAuth {
	case 1:
		routerFile = basePath + "no_actions/router_check_role.go.template"
	case 2:
		routerFile = basePath + "no_actions/router_no_check_role.go.template"
	}

	t1, err := template.ParseFiles(basePath + "model.go.template")
	if err != nil {
		log.Error(err)
		e.Error(500, err, fmt.Sprintf("model模版读取失败！错误详情：%s", err.Error()))
		return
	}
	t2, err := template.ParseFiles(basePath + "no_actions/apis.go.template")
	if err != nil {
		log.Error(err)
		e.Error(500, err, fmt.Sprintf("api模版读取失败！错误详情：%s", err.Error()))
		return
	}
	t3, err := template.ParseFiles(routerFile)
	if err != nil {
		log.Error(err)
		e.Error(500, err, fmt.Sprintf("路由模版失败！错误详情：%s", err.Error()))
		return
	}

	t6, err := template.ParseFiles(basePath + "dto.go.template")
	if err != nil {
		log.Error(err)
		e.Error(500, err, fmt.Sprintf("dto模版解析失败失败！错误详情：%s", err.Error()))
		return
	}
	t7, err := template.ParseFiles(basePath + "no_actions/service.go.template")
	if err != nil {
		log.Error(err)
		e.Error(500, err, fmt.Sprintf("service模版失败！错误详情：%s", err.Error()))
		return
	}

	_ = pkg.PathCreate("./app/" + tab.PackageName + "/apis/")
	_ = pkg.PathCreate("./app/" + tab.PackageName + "/models/")
	_ = pkg.PathCreate("./app/" + tab.PackageName + "/router/")
	_ = pkg.PathCreate("./app/" + tab.PackageName + "/service/dto/")
	_ = pkg.PathCreate(config.GenConfig.FrontPath + "/api/" + tab.PackageName + "/")
	if err != nil {
		log.Error(err)
		e.Error(500, err, fmt.Sprintf("views目录创建失败！错误详情：%s", err.Error()))
		return
	}

	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	var b6 bytes.Buffer
	err = t6.Execute(&b6, tab)
	var b7 bytes.Buffer
	err = t7.Execute(&b7, tab)
	pkg.FileCreate(b1, "./app/"+tab.PackageName+"/models/"+tab.TBName+".go")
	pkg.FileCreate(b2, "./app/"+tab.PackageName+"/apis/"+tab.TBName+".go")
	pkg.FileCreate(b3, "./app/"+tab.PackageName+"/router/"+tab.TBName+".go")
	pkg.FileCreate(b6, "./app/"+tab.PackageName+"/service/dto/"+tab.TBName+".go")
	pkg.FileCreate(b7, "./app/"+tab.PackageName+"/service/"+tab.TBName+".go")

}

func (e Gen) genApiToFile(c *gin.Context, tab tools.SysTables) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	basePath := "template/"

	t1, err := template.ParseFiles(basePath + "api_migrate.template")
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf("数据迁移模版解析失败！错误详情：%s", err.Error()))
		return
	}
	i := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	var b1 bytes.Buffer
	err = t1.Execute(&b1, struct {
		tools.SysTables
		GenerateTime string
	}{tab, i})

	pkg.FileCreate(b1, "./cmd/migrate/migration/version-local/"+i+"_migrate.go")

}
