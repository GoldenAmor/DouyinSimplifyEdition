# DouyinSimplifyEdition
[字节跳动后端青训营] 极简版抖音大项目-Team：3306

# 项目结构
```
├─conf
│
├─conn
│
├─controller
│  │
│  └─vo
│
├─dao
│
├─middleware
│
├─public
│
├─repository
│
├─router
│
└─service
```
## 包说明
### conf
存放项目的配置信息，如cdn、mysql、redis一些服务的配置属性。
### conn
存放一些共用的在项目运行一开始就初始化好的变量，目前只有gorm的连接。
### controller
存放控制器
### vo
存放`controller`中需要的视图模型
### middleware
一些中间件，目前只有jwt
### public
运行时产生的一些缓存文件，需要离线任务去定时清理。
### repository
存放数据库的模型，gorm通过`repository`的结构体去初始化数据库的表和数据。
### router
定义路由
### service
为`controller`层提供数据，`repository`向`vo`的转换在这层实现。
## 调用说明
`router`->`controller`->`service`->`dao`

不允许跨层调用