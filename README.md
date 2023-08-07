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
```bash
go mod tidy

不要执行 go get -u !!!
不要执行 go get -u !!!
不要执行 go get -u !!!
```

## 运行
```bash
go run main.go server
或
make run
```

## 生成swagger文档
```bash
swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin
或
make swag-1
```

## 构建docker镜像
```bash
make docker-build
```

## 自动生成代码
```bash
# sys_columns.sql和sys_tables.sql必须内置到你的数据库中，才能进行代码生成
# 数据库新建表之后运行下列命令，自动生成crud代码
wget --no-check-certificate --quiet \
  --method GET \
  --timeout=0 \
  --header '' \
   'http://localhost:8000/gen/code?table=pink&app_name=pink'
```

