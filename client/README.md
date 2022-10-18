# checkCDN
### client

#### 编辑配置文件

```
[app]
  #指定监听地址
  host = "0.0.0.0"
  #指定监听端口
  port = 8087
  #设置token
  token = "b947f276-e761-431f-928d-30ba93a397c1"
```

#### 使用方法
| Path | /apis/ping |
| ------ | ------ |
|Method|get|
|params|host(string,必须)|
|params|token(string,必须)
```bigquery
http://127.0.0.1:8087/apis/ping?host=www.baidu.com&token=b947f276-e761-431f-928d-30ba93a397c1


{
"code": 0,
"msg": "success",
"data": "\r\n正在 Ping www.baidu.com [39.156.66.18] 具有 32 字节的数据:\r\n来自 39.156.66.18 的回复: 字节=32 时间=19ms TTL=52\r\n\r\n39.156.66.18 的 Ping 统计信息:\r\n 数据包: 已发送 = 1，已接收 = 1，丢失 = 0 (0% 丢失)，\r\n往返行程的估计时间(以毫秒为单位):\r\n 最短 = 19ms，最长 = 19ms，平均 = 19ms\r\n"
}

```



