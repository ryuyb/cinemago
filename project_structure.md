# 目录结构设计

```text
.
├── cmd/                     # 主程序入口
│   └── app/                 # API服务入口
│       └── main.go          # 应用主入口
├── configs/                 # 配置文件
│   ├── config.yml           # 基础配置
│   ├── config.dev.yml       # 开发环境配置
│   └── config.prod.yml      # 生产环境配置
├── docs/                    # 文档
│   ├── docs.go              # 文档生成
│   ├── swagger.json         # Swagger JSON
│   └── swagger.yaml         # Swagger YAML
├── internal/                # 内部模块
│   ├── config/              # 配置管理
│   │   ├── cobra.go         # Cobra配置
│   │   ├── config.go        # 配置加载逻辑
│   │   ├── types.go         # 配置类型定义
│   │   └── wire.go          # Wire配置
│   ├── handler/             # 请求处理器
│   │   ├── UserHandler.go   # 用户处理器
│   │   └── wire.go          # Wire配置
│   ├── logger/              # 日志模块
│   │   ├── logger.go        # 日志实现
│   │   └── wire.go          # Wire配置
│   ├── middleware/          # 中间件
│   │   ├── error_handler.go # 错误处理
│   │   ├── fiber.go         # Fiber中间件
│   │   └── not_found.go     # 404处理
│   ├── model/               # 数据模型
│   │   └── ent/             # Ent ORM
│   │       ├── client.go    # 数据库客户端
│   │       ├── generate.go  # 生成代码
│   │       ├── user.go      # 用户模型
│   │       └── ...          # 其他Ent相关文件
│   ├── repository/          # 数据访问层
│   │   ├── user_repository.go # 用户存储库
│   │   └wire.go            # Wire配置
│   ├── router/              # 路由
│   │   ├── router.go        # 主路由
│   │   ├user_router.go      # 用户路由
│   │   └wire.go            # Wire配置
│   ├── server/              # 服务器
│   │   ├── server.go        # 服务器实现
│   │   └wire.go            # Wire配置
│   ├── service/             # 业务逻辑
│   │   ├── user_service.go  # 用户服务
│   │   └wire.go            # Wire配置
│   └── wire/                # 依赖注入
│       ├── wire.go          # Wire配置
│       └wire_gen.go         # 生成代码
├── pkg/                     # 可复用模块
│   └── logger/              # 日志模块
│       ├── logger.go        # 日志实现
│       └logger_test.go      # 日志测试
├── .air.toml                # 热重载配置
├── Dockerfile               # Docker配置
├── go.mod                   # Go模块文件
├── go.sum                   # 依赖校验
├── Makefile                 # 构建脚本
└── README.md                # 项目说明
```

## 详细目录说明

### cmd/

- 包含程序的入口点
- main.go是应用程序的启动文件，负责：
  - 初始化配置
  - 依赖注入
  - 启动服务器

### configs/

- 存放应用配置文件
- 支持多环境配置(dev/prod)
- 使用viper加载配置

### internal/

- 应用核心实现代码
- config/: 配置加载和解析
- handler/: HTTP请求处理
- logger/: 日志实现
- middleware/: 中间件
- model/: 数据模型定义
  - ent/: Ent ORM实现
- repository/: 数据访问层
- router/: 路由定义
- server/: 服务器实现
- service/: 业务逻辑层
- wire/: 依赖注入配置

### pkg/

- 可复用的公共组件
- logger/: 日志模块实现

## 技术栈说明

1. Web框架: Fiber
2. ORM: Entgo
3. 日志: Zap
   - `zap` 已经使用`fiberzap`集成进`fiber`自带的日志框架中
   - 可以直接使用`github.com/gofiber/fiber/v2/log`记录日志
4. 依赖注入: Wire
5. 配置管理: Viper
