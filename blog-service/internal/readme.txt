本文件夹用于存放工程的内部模块
dao:数据访问层，即 database access objects, 所有与数据相关的操作都会在 dao 层进行, 例如 Mysql, Elastic Search 等;
middleware: HTTP 中间件;
model: 模型层, 用于存放 model 对象;
routers: 路由相关的逻辑;
service: 项目核心业务逻辑;


model 层对数据库进行抽象；
dao 层主要做数据持久相关的工作，负责与数据的增删改查，不涉及业务逻辑，只是负责根据某个条件获得指定数据；
service 层主要负责业务模块的应用逻辑应用设计，主要负责实现业务逻辑；如果涉及到对数据的操作，则要调用已经定义的dao层接口；
controller 层负责具体的业务模块流程的控制，和与前端的交互，这一层需要操作 service；（只负责处理大业务模块的流程排序等简单的操作）