## 进行项目配置
+ [初始化项目目录](https://book.eddycjy.com/golang/gin/api-01.html)
+ conf：用于存储配置文件
+ middleware：应用中间件
+ models：应用数据库模型
+ pkg：第三方包
+ routers 路由逻辑处理(接口编写)
+ runtime：应用运行时数据
+ sql： sql语句用于创建表

## 初始化
1. 初始化pkg\setting\setting.go  ： 获取app.ini配置文件信息
2. 配置pkg\e\code.go 和msg.go的错误处理包
3. 编写工具包pkg\util 获取分页
4. 编写models 定义数据模型

## 项目启动编写路由文件
编写main函数测试路由启动


接口文档
```
登录获取token get请求
/auth?username=test&password=test123
{
  "code": 200,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE1MTg3MjQ2OTMsImlzcyI6Imdpbi1ibG9nIn0.KSBY6TeavV_30kfmP7HWLRYKP5TPEDgHtABe9HCsic4"
  },
  "msg": "ok"
}

//获取全部文章 get请求
/api/v1/articles?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjA5OGY2YmNkNDYyMWQzNzNjYWRlNGU4MzI2MjdiNGY2IiwicGFzc3dvcmQiOiJjYzAzZTc0N2E2YWZiYmNiZjhiZTc2NjhhY2ZlYmVlNSIsImV4cCI6MTU3ODEwOTA1MiwiaXNzIjoiZ2luLWJsb2cifQ.sZmJLB81ihQ_9-cfTzrf8I9fo-RrzXloi1TYl4idZtI
{
    "code": 200,
    "msg": "ok",
    "data": {
        "lists": [
            {
                "id": 2,
                "created_on": 1518700920,
                "modified_on": 0,
                "deleted_on": 0,
                "tag_id": 1,
                "tag": {
                    "id": 1,
                    "created_on": 1518684200,
                    "modified_on": 0,
                    "deleted_on": 0,
                    "name": "",
                    "created_by": "",
                    "modified_by": "",
                    "state": 1
                },
                "title": "123",
                "desc": "1",
                "content": "123456",
                "cover_image_url": "",
                "created_by": "",
                "modified_by": "",
                "state": 1
            }
        ],
        "total": 1
    }
}

//修改tag  put 访问  删除 则 delete访问
/api/v1/tags/1?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjA5OGY2YmNkNDYyMWQzNzNjYWRlNGU4MzI2MjdiNGY2IiwicGFzc3dvcmQiOiJjYzAzZTc0N2E2YWZiYmNiZjhiZTc2NjhhY2ZlYmVlNSIsImV4cCI6MTU3ODEwOTA1MiwiaXNzIjoiZ2luLWJsb2cifQ.sZmJLB81ihQ_9-cfTzrf8I9fo-RrzXloi1TYl4idZtI
name=edit
modified_by=test
state=1
{
    "code": 200,
    "msg": "ok",
    "data": null
}

//删除文章 delete访问
http://localhost:8080/api/v1/articles/2?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjA5OGY2YmNkNDYyMWQzNzNjYWRlNGU4MzI2MjdiNGY2IiwicGFzc3dvcmQiOiJjYzAzZTc0N2E2YWZiYmNiZjhiZTc2NjhhY2ZlYmVlNSIsImV4cCI6MTU3ODEwOTA1MiwiaXNzIjoiZ2luLWJsb2cifQ.sZmJLB81ihQ_9-cfTzrf8I9fo-RrzXloi1TYl4idZtI
{
    "code": 200,
    "msg": "ok",
    "data": null
}

//新增文章 post
tag_id:1
title:test1
desc:1
content:test-content
created_by:qwj
state:1
cover_image_url:https://avatars0.githubusercontent.com/u/39342407?v=4
/api/v1/articles?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjA5OGY2YmNkNDYyMWQzNzNjYWRlNGU4MzI2MjdiNGY2IiwicGFzc3dvcmQiOiJjYzAzZTc0N2E2YWZiYmNiZjhiZTc2NjhhY2ZlYmVlNSIsImV4cCI6MTU3ODEwOTA1MiwiaXNzIjoiZ2luLWJsb2cifQ.sZmJLB81ihQ_9-cfTzrf8I9fo-RrzXloi1TYl4idZtI
{
    "code": 200,
    "msg": "ok",
    "data": null
}