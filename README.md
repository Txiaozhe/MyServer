# MyServer
An echo demo.

* error信息：

```go
UnKnow           = -1 //未知错误
ErrSucceed       = 0  // 成功
ErrInvalidParam  = 1  // 参数错误
ErrInvalidPerm   = 2  // 权限错误
ErrInvalidUser   = 3  // 用户不存在
ErrInactiveUser  = 4  // 用户被禁用
ErrOperationPass = 5  // 操作密码错误
ErrUserExists    = 6  // 用户已存在
ErrInvalidCode   = 7  // 验证码错误
ErrSendCode      = 8  // 验证码发送失败
ErrGroupExists   = 9  // 用户色团已存在
ErrIncorrectPass = 10 // 密码错误
ErrCache         = 11 // 存入缓存出错
ErrConv          = 12 // 类型转换出错
ErrNoChange      = 13 // 未改变
ErrIsInGroup     = 14 // 已经存在于该色团
ErrMysqlNotFound = 15 // 数据库没找到


ErrSession       = 400 // Session 错误
ErrLoginRequired = 401 // 未登录
ErrActive        = 402 // 用户未激活
ErrUpgrade       = 403 // ws连接错误

ErrMysql         = 500 // MySQL 错误

ErrMongoDB       = 600 // MongoDB 错误
```

* 登录：post   /user/login

  params

```json
{
  "name": "",
  "password": ""
}
```
​       return

```json
{
  "status": 0
}
```

* 注册：post   /user/register

  params:

```json
{
  "name": "",
  "password": "",
  "repeat": ""
}
```
​      return: 

```json
{
  "status": 0
}
```

* 登出：post  /user/logout

  return:

```json
{
  "status": 0
}
```

* 修改用户信息： post  /user/reviseUserInfo

​	params:

```json
{
  "name":"",
  "age":0,
  "city":"保定",
  "sex":0, //0代表男，1代表女
  "birthday":"2017-05-21"
}
```
​       return:

```json
{
  "status": 0
}
```

* 获取用户信息：get  /user/getUserInfo

  return: 

```json
{
  "status":0,
  "data": {
  	"id":0,
  	"name":"",
  	"age":0,
  	"city":"保定",
  	"sex":0, //0代表男，1代表女
  	"birthday":"2017-05-21"
  }
}
```

