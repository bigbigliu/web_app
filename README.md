# web_app
go-admin精简api服务

## 开发环境
* 开发语言: Golang
* web 框架: go-admin(gin + gorm)
* 包管理: go mod

## 环境配置
* 安装配置 Golang 运行环境
* 设置go mod 代理
```
export GO111MODULE=auto
export GOPROXY=https://goproxy.cn,direct
```
* 更新依赖包
```go
go mod tidy

不要执行 go get -u !!!
不要执行 go get -u !!!
不要执行 go get -u !!!
```

## 运行
```go
go run main.go server
或
make run
```

## 生成swagger文档
```api
swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin
或
make swag-1
```

sys_columns.sql和sys_tables.sql必须内置到你的数据库中，才能进行代码生成

