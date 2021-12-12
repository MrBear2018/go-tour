这个节点主要完成了三件事情：
1. 初始化目录结构
2. 梳理 API 架构
3. 填充基础代码，构建可执行demo

## 初始化目录结构
煎鱼给的标准目录结构为：
```text
blog-service
├── configs
├── docs
├── global
├── internal
│   ├── dao
│   ├── middleware
│   ├── model
│   ├── routers
│   └── service
├── pkg
├── storage
├── scripts
└── third_party
```
经过分析发现
1. pkg 目录可以是本机所有工程共用的，没有必要单独维护一个pkg目录
2. storage 目录暂时没有必要，Go 编译之后也会放在 GOPATH 那里，不用单独维护
最终我的目录结构是
```text
blog-service
├── configs
├── docs
├── global
├── internal
│   ├── dao
│   ├── middleware
│   ├── model
│   ├── routers
│   └── service
├── scripts
└── third_party
```

## 梳理 API 架构
1. 基础业务功能梳理：

    标签管理：文章所归属的分类，也就是标签。我们平时都会针对文章的内容打上好几个标签，用于标识文章内容的要点要素，这样子便于读者的识别和 SEO 的收录等。

    文章管理：整个文章内容的管理，并且需要将文章和标签进行关联。

2. 设计数据库

    ```text
   blog_article      blog_tag
      文章表           标签表
       ↑ ↓              ↓
         blog_article_tag
           文章标签关联表
    ```

3. HTTP接口规划

    最核心的就是增删改查的 RESTful API 设计和编写，在 RESTful API 中 HTTP 方法对应的行为动作分别如下：
    
    GET：读取/检索动作
    
    POST：新增/新建动作
    
    PUT：更新动作，用于更新一个完整的资源，要求为幂等
    
    PATCH：更新动作，用于更新某一个资源的一个组成部分，也就是只需要更新该资源的某一项，就应该使用 PATCH 而不是 PUT，可以不幂等
    
    DELETE：删除动作
    
4. model层规划

    API 层定下来了，model层就差不多；
    
    这里需要注意的是，为了避免写重复的代码，把表格中共有的部分抽出来放在一个结构体中，命名为 model；
    
    可以提高结构体的可读性和可维护性；名字不一定非得是 model, f但是这种思想可以借鉴；
    
    GO 中没有 JAVA 的类的概念，无法创建基类来维护公共数据，可以用结构体嵌套的方式来打到同样的效果；

## 填充基础代码 
主要就是 router.go 的部分，核心是 注册我们自定义的操作方法；
    
注意：需要手动维护 Article 和 Tag 结构体的实例，后期改为依赖注入的方式，便于管理
    
-[ ] TODO 添加依赖注入框架
