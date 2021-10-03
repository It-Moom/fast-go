# fast-go

> 一个基于Go语言的快速Web开发框架

## 目录结构

```
.
├── README.md                   框架介绍
├── app                         应用目录
│   ├── common                  公共文件库
│   ├── constant                常量库
│   ├── dao                     数据访问库
│   ├── handler                 助手文件库
│   ├── http                    Http访问库
│   ├── model                   模型文件库
│   └── service                 接口服务库
├── bootstrap                   启动文件库
│   └── README.md
├── config                      配置文件库
│   ├── README.md
│   ├── conf                    配置文件
│   └── conf_property           配置文件属性
├── database                    迁移文件及工厂文件库
│   └── README.md
├── go.mod
├── go.sum
├── main.go                     应用入口文件
├── pkg                         内核代码库
│   └──  README.md
├── resources                   资源文件库
│   └── README.md
├── routes                      路由文件库
│   └── README.md
├── storage                     文件资料库
│   └── README.md
└── tests                       单元测试库
    └──  README.md

```

## 所用扩展包

|  包/扩展   | 作用说明  |
|  ----  | ----  |
| [gin](https://github.com/gin-gonic/gin)  | 基础Web框架 |
| [air](https://github.com/cosmtrek/air)  | 热重载 |
| [testify](https://github.com/stretchr/testify) | 单元测试 |
| [gorm](https://github.com/go-gorm/gorm) | ORM框架 |
| [gorm-mysql](https://github.com/go-gorm/mysql) | GORM-MySQL依赖 |
| [viper](https://github.com/spf13/viper) | 配置文件处理 |
| [jwt](https://github.com/dgrijalva/jwt-go) | JWT |
| [go.uuid](https://github.com/satori/go.uuid) | UUID生成器 |

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