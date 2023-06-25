# blog

```
blog-service       
├── configs        // 配置文件
├── docs           // 文档集合
├── global         // 全局变量
├── internal       // 内部模块
│   ├── dao        // 数据访问层
│   ├── middleware // 中间件
│   ├── model      // 模型层
│   ├── routers    // 路由
│   └── service    // 核心业务逻辑
├── pkg            // 模块包
├── storage        // 临时文件
├── scripts        // 脚本
└── third_party    // 第三方工具
```
+ 标签管理

| 功能     | HTTP方法 | 路径        |
|--------|--------|-----------|
| 新增标签   | POST   | /tags     |
| 删除指定标签 | DELETE | /tags/:id |
| 更新指定标签 | PUT    | /tags/:id |
| 获取标签列表 | GET    | /tags     |

+ 文章管理

| 功能     | HTTP方法 | 路径            |
|--------|--------|---------------|
| 新增文章   | POST   | /articles     |
| 删除指定文章 | DELETE | /articles/:id |
| 更新指定文章 | PUT    | /articles/:id |
| 获取指定文章 | GET    | /articles/:id |
| 获取文章列表 | GET    | /articles     |