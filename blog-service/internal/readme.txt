本文件夹用于存放工程的内部模块
dao:数据访问层，即 database access objects, 所有与数据相关的操作都会在 dao 层进行, 例如 Mysql, Elastic Search 等;
middleware: HTTP 中间件;
model: 模型层, 用于存放 model 对象;
routers: 路由相关的逻辑;
service: 项目核心业务逻辑;