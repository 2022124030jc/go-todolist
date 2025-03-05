为你的 `todolist` 项目后端编写一个项目介绍，可以帮助其他开发者快速了解项目的功能、架构和技术栈。以下是一个详细的项目介绍模板，你可以根据实际情况进行调整和补充。

---

# Todo List 后端项目介绍

## 概述

**Todo List 后端** 是一个用于管理待办事项的应用程序的后端服务。该服务提供了基本的 CRUD（创建、读取、更新、删除）操作，并支持一些高级功能如清除已完成任务等。整个项目使用 Go 语言开发，并基于 Gin 框架构建 RESTful API。

### 主要特性

- **CRUD 操作**：支持对任务（todos）的基本增删改查操作。
- **清除已完成任务**：提供接口以批量删除所有已完成的任务。
- **跨域支持**：通过 CORS 中间件支持跨域请求。
- **数据库持久化**：使用 GORM 进行数据库交互，默认支持 SQLite 和 MySQL 数据库。

### 技术栈

- **编程语言**：Go
- **Web 框架**：Gin
- **数据库 ORM**：GORM
- **数据库**：SQLite / MySQL
- **HTTP 协议**：RESTful API

## 功能说明

### 1. 创建任务 (Create)

- **API**: `POST /api/todos`
- **描述**: 创建一个新的待办事项。
- **请求体示例**:
  ```json
  {
    "mes": "Buy groceries",
    "completed": false
  }
  ```
- **响应示例**:
  ```json
  {
    "id": 1,
    "mes": "Buy groceries",
    "completed": false
  }
  ```

### 2. 获取所有任务 (Read)

- **API**: `GET /api/todos`
- **描述**: 获取所有待办事项列表。
- **响应示例**:
  ```json
  [
    {
      "id": 1,
      "mes": "Buy groceries",
      "completed": false
    },
    {
      "id": 2,
      "mes": "Finish project report",
      "completed": true
    }
  ]
  ```

### 3. 更新任务 (Update)

- **API**: `PUT /api/todos/:id` 或 `PATCH /api/todos/:id`
- **描述**: 更新指定 ID 的待办事项内容或状态。
- **请求体示例**:
  ```json
  {
    "mes": "Buy groceries and fruits",
    "completed": true
  }
  ```
- **响应示例**:
  ```json
  {
    "id": 1,
    "mes": "Buy groceries and fruits",
    "completed": true
  }
  ```

### 4. 删除任务 (Delete)

- **API**: `DELETE /api/todos/:id`
- **描述**: 删除指定 ID 的待办事项。
- **响应示例**:
  ```json
  {
    "message": "Todo deleted successfully"
  }
  ```

### 5. 清除已完成任务 (Clear Completed)

- **API**: `DELETE /api/todos/clear-completed`
- **描述**: 批量删除所有已完成的待办事项。
- **响应示例**:
  ```json
  {
    "message": "Completed todos cleared successfully"
  }
  ```

## 目录结构

```
go-todolist/
├── main.go                    # 主入口文件
├── api/                       # API 路由定义
│   └── api.go                 # 定义路由和处理函数
├── models/                    # 数据模型定义
│   └── todo.go                # Todo 结构体定义
├── database/                  # 数据库初始化和迁移
│   └── init_db.go             # 初始化数据库连接
└── middleware/                # 中间件
    └── cors.go                # CORS 中间件
```

## 环境配置

### 配置文件

项目使用环境变量来配置数据库连接和其他必要的设置。你可以在 `.env` 文件中添加以下内容：

```plaintext
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=tododb
```

### 依赖管理

项目使用 Go Modules 进行依赖管理。确保在项目根目录下运行以下命令来安装依赖：

```sh
go mod tidy
```

## 如何运行

### 1. 克隆项目

```sh
git clone https://github.com/your-username/go-todolist.git
cd go-todolist
```

### 2. 安装依赖

```sh
go mod tidy
```

### 3. 初始化数据库

根据你的数据库类型（SQLite 或 MySQL），编辑 `database/init_db.go` 文件中的数据库连接字符串。

对于 SQLite：

```go
dsn := "test.db"
db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
```

对于 MySQL：

```go
dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

### 4. 启动服务器

```sh
go run main.go
```

默认情况下，服务器将在 `localhost:8080` 上运行。

## 测试

为了确保代码的质量和稳定性，建议编写单元测试。可以使用 Go 内置的测试框架来编写和运行测试。

### 编写测试

在每个包中创建 `_test.go` 文件，并编写测试用例。例如，在 `models/todo_test.go` 中编写测试用例。

### 运行测试

```sh
go test ./...
```

## 贡献指南

欢迎任何贡献！请按照以下步骤提交 Pull Request：

1. Fork 仓库并在本地克隆。
2. 创建新分支 (`git checkout -b feature-branch`)。
3. 提交更改 (`git commit -am 'Add some feature'`)。
4. 推送分支 (`git push origin feature-branch`)。
5. 在 GitHub 上创建 Pull Request。
