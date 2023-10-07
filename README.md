# go-chat
该项目是是要使用golang开发的基于websocket的简单聊天室

项目接口文档 https://apifox.com/apidoc/shared-a4e0cb5e-3af5-4733-8d2e-edced8e98c62

pkg目录是已经封装好的工具包直接使用即可

在config/app.yaml目录下，新建如下配置文件
```yaml
Serve:
  addr: ":8084"
  read_timeout: 5s #读超时的时间
  write_timeout: 5s #写超时的时间
  default_timeout: 30s #优雅退出的时间
Mysql:
  user: "root"
  password: "123456"
  host: "127.0.0.1"
  port: 3306
  dbName: "chat_app"
Log: # 日志配置文件
  Level: debug
  LogSavePath: "storage/Applogs/" # 日志保存路径
  HighLevelFile: error          # 高级别日志文件名
  LowLevelFile: info          # 低级别文件名
  LogFileExt: .log          # 日志文件后缀
  MaxSize: 10 # 日志文件最大大小
  MaxAge: 30 # 日志文件最大保存天数
  MaxBackups: 7 # 最大备份数
  Compress: false # 是否压缩
SMTPInfo:
  host: "smtp.qq.com"
  port: 465
  userName: "***@qq.com"
  password: "***" #密码
  isSSL: true
  from: "***@qq.com" #发件人
  to: #接收人
    - "***@qq.com"
Redis:
  addr: "127.0.0.1:6379"
  password: "***"
  poolSize: 20
  CacheTime: 1h
Work: #工作池配置
  TaskChanCapacity: 50 # 任务队列容量
  WorkerChanCapacity: 50 # 工作队列容量
  WorkerNum: 10 # 工作池数量
Rule:
  delUserTime: "60s"
  delCodeTime: "120s"
  accountMaxNums: 5
  DefaultAccountAvatar: "http://lycmall.lyc666.xyz/chat/first.jpg"
  DefaultClientTimeout: "20s"
  FileMaxSize: 1024000
  DefaultInsertDataNum : 50
Auto:
  retry:
    timeDuration: 5s
    timeCount: 3
Token:
  Key: "***"
  AccessTokenExpire: 168h
  RefreshTokenExpire: 720h
  AccountTokenExpire: 24h
  AuthType: "x_token"
  AuthKey: "Claims"
Page:
  DefaultPageSize: 10
  MaxPageSize: 200
  PageKey: page
  PageSizeKey: page_size
AliyunOSS:
  Endpoint: "http://oss-cn-hangzhou.aliyuncs.com"
  AccessKeyId: "***"
  AccessKeySecret: "***"
  BucketName: "lycmall2"
  BucketUrl: "chat"
  BasePath: "***"
Rocketmq:
  Addr: "192.168.28.3"
  Port: 9876


```

项目功能：

  1.一个用户可以有多种职责，各个账号之间的信息是独立的
  
  2.好友之间可以进行好友申请

  3.好友之间可以发送消息，消息可以是文本，图片，视频之类的。消息支持pin，撤回，回复等相关操作

  4.支持消息的已读未读功能

  5.支持离线消息

  6.用户长时间未进行操作，则将其强制下线

  7.系统消息进行多设备推送


项目亮点：
  1.使用单用户多账号，实现用户的解耦
  
  2.支持多设备同时在线，消息同步推送

  3.使用了协程池技术，做一些任务的异步化

  4.邮箱模块使用map + 定时器

  5.使用令牌桶对ip和api进行限流

  6.代码风格简洁，清爽。单元测试规范

  
项目待优化点：
  1.数据库的表结构设计不太合理，例如消息表的设计
  
  2.ws服务可以使用RPC通信





  
  
  
