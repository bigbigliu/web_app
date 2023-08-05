FROM golang:1.20.4 AS builder

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .

RUN --mount=type=cache,mode=0777,id=go-mod,target=/go/pkg/mod \
    go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download

WORKDIR /src

COPY . /src

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o web_app

FROM alpine as prod

RUN  echo "Asia/Shanghai" > /etc/timezone \
    && rm -f /etc/localtime \
    && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime


WORKDIR /svr

COPY --from=builder /src/web_app /svr

COPY --from=builder /src/config/settings.yml /svr/conf/settings.yml

EXPOSE 8000

CMD ["./web_app", "server"]
