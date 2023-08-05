PROJECT:=go-admin

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -o web_app .

# make build-linux
build-linux:
	@docker build -t web_app:latest .
	@echo "build successful"

stop:
    # delete go-admin-api container
	@if [ $(shell docker ps -aq --filter name=go-admin --filter publish=8000) ]; then docker-compose down; fi
	#@if [ $(shell docker ps -aq --filter name=go-admin --filter publish=8000) ]; then docker rm -f go-admin; fi
	#@echo "go-admin stop success"


deploy:

	#@git checkout master
	#@git pull origin master
	make build-linux
	make run

swag-1:
	# 生成swagger文档
	swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

swag-2:
	# 生成swagger文档
	swag i -g init_router.go -dir app/admin/router --instanceName admin --parseDependency -o docs/admin

# 生成app代码命令示例,生成app并注册app路由
create-app:
	go run main.go app -n apple