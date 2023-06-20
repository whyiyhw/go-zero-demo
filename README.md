# go-zero-demo

## 构建文档
- 请查看 doc 目录

## 项目目录

- 遵循 go-zero 社区原则下 ， 项目目录结构如下

```
├── service
│   ├── demo
│   │   ├── api
│   │   │   ├── doc
│   │   │   │── etc   
│   │   │   ├── internal
│   │   │   ├── demo.api
│   │   │   ├── demo.go
│   │   ├── model
│   │   ├── dao
```

- 在生成目录下加入 doc 目录，用于存放 api 文件，如 user.api,这些文件会通过 demo.api 进行 import
- model 与 dao 目录下存放数据库模型，通过 sql 文件自动生成，不需要手动创建

## 项目基本功能说明

### 验证器 （本地化支持）
- 目前 `go-zero` 支持的 验证器很有限，我们选择接入第三方的验证器，目前选择的是 `go-playground/validator`
- 使用 `goctl` 的模板功能，替换生成的代码，自动生成验证器的代码

```shell
# 先进行模板初始化
goctl  template init 
# Templates are generated in C:\Users\Administrator\.goctl\1.5.3
# 对于mac 应该实在 ~/.goctl/1.5.3/ 目录下
# 然后对模板进行替换，比如现在我们的验证器，每次都需要手动去写，我们可以把这个写入到模板
# 替换的模板我放在了 项目一级目录 template 下面
```

### 统一全局响应码 （已实现）
- 在 common 中 加入 response 文件夹，用于存放统一的响应码以及相关的方法
- 加入 xerr 目录，所有的错误响应实现 了 error 接口，这样可以统一处理错误响应

### 日志
- 默认为 console 日志，后续都是容器部署，可以直接采集容器日志，所以不需要文件日志

### 中间件

#### access_log  记录访问&响应日志 （已实现）

#### [`jwt: Auth`认证](./doc/04用户详情.md) （已实现）


### [单元测试](./doc/05单元测试.md) （已实现）

### [数据迁移](./doc/06数据迁移.md) （已实现）

### 队列

### 定时任务

### 监控（waiting） （待实现）

## 引入包说明

### model 操作

- `gorm.io/gorm` 与 `gorm.io/gen` 来进行 sql 操作

### migrate 数据迁移

- 使用了原生的 `gorm` 自动迁移来处理

### 验证器

- [github.com/go-playground/validator/v10](https://github.com/go-playground/validator)

### docker-compose 编译&部署

```shell
docker compose build && docker compose up -d
```

## 工具的安装

- `gentool` && `goctl`
```shell
go install gorm.io/gen/tools/gentool@latest

go install github.com/zeromicro/go-zero/tools/goctl@latest
```

```tool
gentool -db "postgres" -dsn "host=localhost user=root password=123456 dbname=demo port=18886 sslmode=disable TimeZone=Asia/Shanghai" -outPath "./service/demo/dao"
```