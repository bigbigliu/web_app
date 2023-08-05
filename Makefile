PROJECT:=web_app
GOPATH:=$(shell go env GOPATH)
.PHONY : build

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -o web_app .

swag-1:
	# 生成swagger文档
	swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

swag-2:
	# 生成swagger文档
	swag i -g init_router.go -dir app/admin/router --instanceName admin --parseDependency -o docs/admin

# 生成app代码命令示例,生成app并注册app路由
create-app:
	go run main.go app -n apple

docker-build:
	docker build -t web_app:v1 .

run:
	go run main.go server
