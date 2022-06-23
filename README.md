# fast-go

> 一个基于Go语言的快速Web开发框架

## 目录结构

```
.
├── app                         应用目录
│   ├── cmd                     命令目录
│   ├── common                  公共文件库
│   ├── constant                常量库
│   ├── dao                     数据访问库
│   ├── handler                 助手文件库
│   ├── http                    Http访问库
│   │ ├── controller            控制器目录
│   │ ├── middleware            中间件目录
│   │ └── request               请求验证器目录
│   ├── model                   模型文件库
│   │ ├── dto                   数据传输对象库
│   │ ├── entity                实体类库
│   │ └── vo                    回调对象库
│   └── service                 接口服务库
├── bootstrap                   引导文件库
├── config                      配置文件库
├── database                    迁移文件及工厂文件库
│   ├── factory                 工厂文件库
│   ├── migrations              迁移文件库
│   └── seeder                  种子文件库
├── main.go                     应用入口文件
├── pkg                         内核代码库
│   ├── app                     应用代码库
│   ├── auth                    认证代码库
│   ├── cache                   缓存代码库
│   ├── config                  配置代码库
│   ├── console                 控制台代码库
│   ├── database                数据库操作代码库
│   ├── file                    文件操作代码库
│   ├── heplers                 助手代码库
│   ├── limiter                 限流代码库
│   ├── logger                  日志代码库
│   ├── migrate                 迁移操作代码库
│   ├── paginator               分页操作代码库
│   ├── redis                   Redis操作代码库
│   ├── seed                    种子操作代码库
│   ├── str                     字符串操作代码库
│   └── utils                   工具代码库
├── public                      公开目录
├── resources                   资源文件库
├── routes                      路由文件库
│   ├── api                     API路由文件库
│   └── web                     Web路由文件库
├── storage                     文件资料库
│   └── logs                    日志文件库
├── tests                       单元测试库
└── tmp                         编译文件库    

```

## 所用扩展包

| 包/扩展                                                       | 作用说明         |
|------------------------------------------------------------|--------------|
| [gin](https://github.com/gin-gonic/gin)                    | 基础Web框架      |
| [air](https://github.com/cosmtrek/air)                     | 热重载          |
| [cast](https://github.com/spf13/cast)                      | 类型转换         |
| [testify](https://github.com/stretchr/testify)             | 单元测试         |
| [gorm](https://github.com/go-gorm/gorm)                    | ORM框架        |
| [gorm-mysql](https://github.com/go-gorm/mysql)             | GORM-MySQL依赖 |
| [viper](https://github.com/spf13/viper)                    | 配置文件处理       |
| [jwt](https://github.com/dgrijalva/jwt-go)                 | JWT          |
| [go.uuid](https://github.com/satori/go.uuid)               | UUID生成器      |
| [zap](https://github.com/uber-go/zap)                      | 日志处理扩展       |
| [lumberjack](https://github.com/natefinch/lumberjack)      | 日志滚动方案       |
| [go-redis](https://github.com/go-redis/redis)              | Redis操作      |
| [limiter](https://github.com/ulule/limiter)                | 限流器          |
| [ansi](https://github.com/mgutz/ansi)                      | 终端高亮         |
| [cobra](https://github.com/spf13/cobra)                    | 命令行处理        |
| [strcase](https://github.com/iancoleman/strcase)           | 处理大小写        |
| [go-pluralize](https://github.com/gertd/go-pluralize)      | 处理英文单复数      |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 参数验证器        |
| [faker](https://github.com/bxcodec/faker)                  | 假数据填充        |

## 环境说明

- Go
    - 版本: 1.18.x
- MySQL
    - 版本: 5.7.x
- Redis
    - 版本: 6.x

## 运行说明

- 1.首先拉取项目到本地，然后进入项目目录，执行如下命令：` cp .yml.example .yml`
- 2.进入项目根目录执行：`go mod tidy`，整理依赖
- 3.配置项目下的 `.yml` 文件,填写项目数据库信息，如果没有数据库信息，请先创建数据库，然后在 `.yml` 文件中填写数据库信息
- 4.根据`.yml`配置的数据库名称到数据库服务中创建对应数据库
- 5.如果本机环境有安装 air，则进入项目根目录，可以通过如下命令进行热重载运行项目：` air run` 或者 ` air run -d`(常驻进程，开发时不建议)
- 6.如果本机环境没有安装 air，则进入项目根目录，可以通过如下命令运行项目：` go run main.go`
- 7.也可通过框架封装的命令：` go run main.go serve`，运行项目，并启动服务

## 命令说明
- 1.查看所有命令：项目根目录执行：`go run main.go -h`
- 2.查看子命令举例：项目根目录执行：`go run main.go cache -h` 即可查看 cache 命令的帮助信息

## 开发说明
- 所有包名尽可能使用1个单词命名，如果包名中有多个单词，则使用下划线分隔，如：`http_response`
- 每个包中的文件名尽可能使用单个单词命名，如果文件名中有多个单词，则使用下划线分隔，如：`http_response.go`
- 尽量不要使用复数命名包、文件、函数、方法等名，除非语义化需求
- 结构体遵循大驼峰命名法，如：`User`、`UserInfo`
- 变量名尽可能使用小驼峰命名法，如：`user`、`userInfo`
- 函数名需注意：如果需对外开放访问权限则使用大驼峰命名法，如：`GetUser`、`GetUserInfo`,如果仅对内部函数访问权限则使用小驼峰命名法，如：`getUser`、`getUserInfo`

## Commit提交规范

> 格式: type(scope) : subject

#### type: 本次 commit 的类型，诸如 bugfix、docs、style 等，参考如下

- feat：添加新功能
- fix：修补缺陷
- docs：修改文档
- style：修改格式
- refactor：重构
- perf：优化
- test：增加测试
- chore：构建过程或辅助工具的变动
- revert：回滚到上一个版本

#### scope: 本次 commit 波及的范围(一般为文件名)

#### subject: 简明扼要的阐述下本次 commit 的主旨(结尾无需添加标点)