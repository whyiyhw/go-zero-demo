FROM golang:1.20-alpine AS builder

# 为镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/build
WORKDIR /build

RUN  apk add tzdata

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执行文件 app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./service/demo/api/demo.go

# 创建一个小镜像
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

COPY service/demo/api/etc /etc

# 从builder镜像中把/build/app 拷贝到当前目录
COPY --from=builder /build/app /app
#COPY ./app /app

EXPOSE 8888

CMD ["/app"]