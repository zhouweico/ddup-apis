# DDUP

## 项目说明
基于 Golang 的 RESTful API 服务，使用 PostgreSQL 作为数据存储。

## 已实现功能

### 用户管理
- [x] 用户注册
- [x] 用户登录
- [x] 用户退出
- [x] JWT 认证
- [x] 获取用户详情
- [x] 更新用户信息
- [x] 修改密码
- [x] 删除用户（软删除）

### 个人资料管理
- [x] 支持多种资料类型
  - 基本信息
  - 项目经历
  - 个人项目
  - 展览经历
  - 演讲经历
  - 写作经历
  - 获奖经历
  - 特色展示
  - 工作经历
  - 志愿者经历
  - 教育经历
  - 认证证书
  - 联系方式
  - 团队信息
- [x] 资料项排序管理
- [x] 资料可见性控制
- [x] 元数据扩展支持
- [x] 附件管理

### 组织管理
- [x] 创建组织
- [x] 获取组织列表
- [x] 更新组织信息
- [x] 删除组织
- [x] 加入组织
- [x] 获取组织成员列表
- [x] 添加组织成员
- [x] 更新成员信息
- [x] 移除组织成员
- [x] 设置成员角色

### 安全特性
- [x] JWT Token 认证
- [x] 资源访问控制
- [x] 密码加密存储
- [x] CORS 跨域支持

### 系统特性
- [x] RESTful API 设计
- [x] Swagger API 文档
- [x] 统一响应格式
- [x] 分页查询支持
- [x] 环境配置管理
- [x] 数据库支持
  - PostgreSQL
  - MySQL
  - SQLite
- [x] 数据库连接池
- [x] 数据库重试机制
- [x] 健康检查
- [x] 日志系统
  - 同时输出到控制台和文件
  - 日志轮转
  - 支持不同日志级别
  - 结构化日志输出
  - 自动记录调用位置
- [x] 统一错误处理
  - 自定义错误类型
  - 错误码标准化
  - 统一错误响应格式
  - 错误日志记录  

## 技术栈
- Go 1.21+
- 支持的数据库
  - PostgreSQL 10+（推荐 11+）
  - MySQL 5.7+
  - SQLite 3+
- Docker (可选)

## 开发环境搭建
1. 数据库环境
```bash
docker run -d \
    --name ddup-postgres-dev \
    -p 5432:5432 \
    -e POSTGRES_PASSWORD=Admin@123456 \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v $HOME/data/ddup/data:/var/lib/postgresql/data \
    postgres:17.0
```

2. 创建数据库和用户

```bash
# 1. 创建用户并设置密码
docker exec -it ddup-postgres-dev psql -U postgres -c "CREATE USER ddup WITH PASSWORD 'Ddup@123456';"

# 2. 创建数据库
docker exec -it ddup-postgres-dev psql -U postgres -c "CREATE DATABASE ddup owner ddup;"


# 3. 授予权限
docker exec -it ddup-postgres-dev psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE ddup TO ddup;"
docker exec -it ddup-postgres-dev psql -U postgres -d ddup -c "GRANT ALL ON SCHEMA public TO ddup;"
```

3. 复制环境变量文件

```bash
cp .env.example .env
```

4. 修改环境变量配置

5. 运行服务
```bash
go run cmd/api/main.go
```

## 项目结构
```
.
├── cmd/              # 主要的应用程序入口
│   └── api/          # API 服务入口
├── docs/             # 文档
├── internal/         # 私有应用程序和库代码
│   ├── config/       # 配置
│   ├── db/           # 数据库
│   ├── dto/          # 数据传输对象
│   ├── errors/       # 错误处理
│   ├── handler/      # HTTP 处理器
│   ├── logger/       # 日志工具
│   ├── middleware/   # HTTP 中间件
│   ├── model/        # 数据库模型
│   ├── repository/   # 数据库操作
│   ├── service/      # 业务逻辑
│   └── utils/        # 通用工具
└── scripts/          # 脚本和工具
```

## 开发工具

### 提交代码

使用提供的脚本自动生成 Swagger 文档并提交代码：

```bash
./scripts/commit.sh "提交信息"
```

### 单元测试

```bash
# 执行所有测试
./scripts/test.sh

# 执行特定包的测试
./scripts/test.sh internal/handler
```

### API 测试

```bash
# 执行所有测试
./scripts/api_test.sh
```

## 部署

### Docker Compose 部署

1. 部署运行

```bash
# 部署（会自动调用 build.sh）
./scripts/deploy.sh

# 如果只想构建镜像
./scripts/build.sh

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down

# 重启服务
docker-compose restart
```

## 配置说明

主要配置项（.env 文件）：

- 服务配置：端口、环境等
- 数据库配置：连接信息、连接池参数等
- JWT 配置：密钥、过期时间等
- 健康检查配置：检查间隔等
- 日志配置：
  - 日志级别
  - 日志文件路径
  - 日志文件大小限制
  - 日志文件保留数量
  - 日志文件保留天数
  - 是否压缩

## 贡献

欢迎提交 Issue 和 Pull Request

## 许可证

MIT License