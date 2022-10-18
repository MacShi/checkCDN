# checkCDN
### server

#### 编辑配置文件

```
#server启动的服务
[app]
  debug = false
  host = "0.0.0.0"
  port = 8088
  #用户访问server的服务的认证信息
  token = "b947f276-e761-431f-928d-30ba93a397c1"

#server可以调用的client
[[client.infos]]
  host = "x.x.x.x"
  port = "80"
  #server调用client的认证信息
  token = "b947f276-e761-431f-928d-30ba93a397c1"
[[client.infos]]
  host = "x.x.x.x"
  port = "8087"
  token = "b947f276-e761-431f-928d-30ba93a397c1"
```

#### 使用方法
| Path | /apis/ping |
| ------ | ------ |
|Method|get|
|params|host(string,必须)|
|params|token(string,必须)
```bigquery
http://127.0.0.1:8088/apis/ping?token=b947f276-e761-431f-928d-30ba93a397c1&host=www.baidu.com

{
  "code": 0,
  "msg": "success",
  "data": {
    "have_cdn": "Yes",
    "ip_addr": [
      "110.242.68.4",
      "103.235.46.40"
    ],
    "loss_rate": [
      "0%",
      "0%"
    ]
  }
}
```



