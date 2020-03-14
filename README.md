# gin_blog接口文档

| 待新增             | 功能描述                             | 预计日期  |
| ------------------ | ------------------------------------ | --------- |
| 文章评论           | 每篇文章可以进行评论                 | 2020.3.1  |
| 文章点赞点踩阅读数 | 对文章进行喜欢或不喜欢操作           | 2020.3.15 |
| 文章推荐           | 侧边栏对点赞或评论数较多文章进行推荐 | 2020.4.1  |
|                    |                                      |           |

> ## 每次改动，都会写在这里，请注意阅读
>
> # 重要

| 改动类                      | 功能描述                                                     | 改动时间  |
| --------------------------- | ------------------------------------------------------------ | --------- |
| 现将token信息接收放到header | 所有接口需要token验证的需要将token放入header而不是url中，**所有带有/api/v1的接口都需要在header中带token** | 2020.3.13 |
| 接口url改变                 | 文章和评论接口不再需要token就能获取                          | 2020.3.13 |
| 所有get请求参数变化         | 将所有get请求参数换为url拼接，其他请求处带id请求外，全部变为xxx-form形式请求 | 2020.3.14 |

**重要**

| 新增接口 | 功能目录     | 功能描述                                                  | 新增时间  |
| -------- | ------------ | --------------------------------------------------------- | --------- |
| 页面分页 | 获取接口文章 | 在url新增路径中增加？page=1，进行页面分页，一页为10条数据 | 2020.1.14 |
| tag分页  | 获取tag分页  | 在url新增路径中增加？page=1，进行页面分页，一页为10条数据 | 2020.1.14 |
| 评论接口 | 获取评论     | 新增增加一级评论，二级评论和删除评论功能                  | 2020.3.1  |



注：所有url需加上本身端口  例：

localhost:8080 /authusername=test&password=test123

所有post数据以qs.stringfy转化为url形式传递

```http
name=q2&created_by=qwj&state=1
```

## 登录获取token 

```go
get方式  /auth?username=test&password=test123
```

```go
响应
{
  "code": 200,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE1MTg3MjQ2OTMsImlzcyI6Imdpbi1ibG9nIn0.KSBY6TeavV_30kfmP7HWLRYKP5TPEDgHtABe9HCsic4"
  },
  "msg": "ok"
}
```

## 文章接口

### 获取所有文章

```go
get 请求
/articles?tag_id=1&sort=created_on
```

```go
响应
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
```

| 字段名     | 字段内容示例 | 注释                                           |
| ---------- | ------------ | ---------------------------------------------- |
| tag_id     | 标签内容     | 非必须，输入按照标签id来搜索文章               |
| state      | 发布状态     | 非必须，输入则按照发布状态来搜索文章           |
| created_by | 发布人       | 非必须，输入则按照发布人来搜索文章             |
| sort       | 排序         | 非必须，输入（created_on），则返回时间倒序结果 |
| page       | 分页         | 非必须，输入页数，获取10条记录                 |



### 文章详情

```go
get方式
/articles/{id}
/articles/7
```

```go
响应如下：
{
    "code": 200,
    "msg": "ok",
    "data": {
        "id": 6,
        "created_on": 1578107255,
        "modified_on": 1578107255,
        "deleted_on": 0,
        "tag_id": 1,
        "tag": {
            "id": 1,
            "created_on": 1518684200,
            "modified_on": 1578098690,
            "deleted_on": 0,
            "name": "edot",
            "created_by": "",
            "modified_by": "test",
            "state": 1
        },
        "comment": [
            {
                "id": 10,
                "created_on": 1581731034,
                "modified_on": 1581731034,
                "deleted_on": 1581733788,
                "content": "777",
                "article_id": 6,
                "created_by": "test",
                "parent_id": 1,
                "like_count": 0,
                "dislike_count": 0
            },
            {
                "id": 11,
                "created_on": 1581733551,
                "modified_on": 1581733551,
                "deleted_on": 0,
                "content": "1113",
                "article_id": 6,
                "created_by": "test",
                "parent_id": 1,
                "like_count": 0,
                "dislike_count": 0
            },
            {
                "id": 12,
                "created_on": 1581733555,
                "modified_on": 1581733555,
                "deleted_on": 0,
                "content": "2223",
                "article_id": 6,
                "created_by": "test",
                "parent_id": 1,
                "like_count": 0,
                "dislike_count": 0
            }
        ],
        "title": "test1",
        "desc": "1",
        "content": "test-content",
        "cover_image_url": "https://avatars0.githubusercontent.com/u/39342407?v=4",
        "created_by": "test",
        "modified_by": "",
        "state": 1
    }
}
```



### 新增文章

```go
post方式
/api/v1/articles
```

| 字段名          | 字段内容示例                                          | 注释（必需）                                               |
| --------------- | ----------------------------------------------------- | ---------------------------------------------------------- |
| title           | test1                                                 | 文章标题                                                   |
| content         | test-content                                          | 内容                                                       |
| created_by      | qwj                                                   | 创建人                                                     |
| state           | 1                                                     | 发布状态（0为草稿态，编辑好未发布存入数据库，1为直接发布） |
| cover_image_url | https://avatars0.githubusercontent.com/u/39342407?v=4 | 封面图片地址                                               |
| desc            | 1                                                     | 简述                                                       |
| tag_id          | 1                                                     | 标签ID                                                     |
| token（header） | token                                                 | 需要放在header中，（必需）                                 |

```
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```

### 修改文章

```go
put方式
/api/v1/articles/5
5为指定ID
```

| 字段名          | 字段内容示例                                          | 注释(非注释则非必需)                                       |
| --------------- | ----------------------------------------------------- | ---------------------------------------------------------- |
| title           | test1                                                 | 文章标题                                                   |
| content         | test-content                                          | 内容                                                       |
| cover_image_url | https://avatars0.githubusercontent.com/u/39342407?v=4 | 封面图片地址                                               |
| tag_id          | 1                                                     | 标签ID                                                     |
| state           | 1                                                     | 发布状态（0为草稿态，编辑好未发布存入数据库，1为直接发布） |
| modified_by     | qwj                                                   | 修改人（必需）                                             |
| desc            | 1                                                     | 简述                                                       |
| token（header） | token                                                 | 需要放在header中，（必需）                                 |

```go
响应
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```

### 删除文章

```go
delete 方式 需要在header中带token
/api/v1/articles/3
3为文章ID
```

```go
响应
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```

## tag接口

### 获取所有tag

```go
get方式
/tags
```

```go
响应
{
    "code": 200,
    "msg": "ok",
    "data": {
        "lists": [
            {
                "id": 1,
                "created_on": 1518684200,
                "modified_on": 1578098690,
                "deleted_on": 0,
                "name": "edot",
                "created_by": "",
                "modified_by": "test",
                "state": 1
            }
        ],
        "total": 1
    }
}
```

### tag分页

```go
get 请求
/tags?page={输入页数,如1}
每页为10个记录
```

### 增加tag

```go
post
/api/v1/tags
```

| 字段名     | 字段值示例 | 注释(必需) |
| ---------- | ---------- | ---------- |
| name       | edit       | tag名      |
| created_by | qwj        | 创建人     |
| state      | 1          | 状态       |

```go
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```

### 修改tag

```go
put
/api/v1/tags/2
2为tag的id
```

| 字段名     | 字段值示例 | 注释         |
| ---------- | ---------- | ------------ |
| name       | test       | tag名(必需)  |
| modfied_by | test       | 修改者(必需) |
| state      | 1          | 状态         |

### 删除tag

```go
delete
/api/v1/tags/2
(2为id,删除后，数据仍存在于数据库中,但会有delete_on属性表示为删除)
```

```go
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```

## 评论接口

### 获取评论

```go
get
/comment?articleId={articleId}
根据articleId即文章id获取该文章下的评论
```

| 字段名    | 字段值示例 | 注释                               |
| --------- | ---------- | ---------------------------------- |
| articleId | 7          | 根据文章id获取该文章下的评论，必需 |

```go
{
    "code": 200,
    "msg": "ok",
    "data": {
        "lists": [
            {
                "id": 11,
                "created_on": 1581733551,
                "modified_on": 1581733551,
                "deleted_on": 0,
                "content": "1113",
                "article_id": 6,
                "created_by": "test",
                "parent_id": 1,
                "like_count": 0,
                "dislike_count": 0
            },
            {
                "id": 12,
                "created_on": 1581733555,
                "modified_on": 1581733555,
                "deleted_on": 0,
                "content": "2223",
                "article_id": 6,
                "created_by": "test",
                "parent_id": 1,
                "like_count": 0,
                "dislike_count": 0
            }
        ]
    }
}
```

### 删除评论

```go
delete
/comment/{id}
```

```go
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```

### 新增评论

```go
post
/comment
```

| 字段名     | 字段值示例 | 注释（必需）                                |
| ---------- | ---------- | ------------------------------------------- |
| content    | hhh        | 评论内容                                    |
| created_by | qwj        | 创建人，（评论不需要登录，相当于留言）      |
| parent_id  | 0          | 父级评论，为0则为1级评论，否则则为评论id如7 |
| article_id | 7          | 文章id                                      |

```go
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```

