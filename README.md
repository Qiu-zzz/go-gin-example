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

