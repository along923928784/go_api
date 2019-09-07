## 项目地址

https://github.com/along923928784/go_api

## 小程序扫码预览

![小程序码](https://github.com/along923928784/go_api/blob/master/image/wx.png)

## 部分接口文档
![接口文档](https://github.com/along923928784/go_api/blob/master/image/test_api.png)

## 重要: 如何运行

git clone https://github.com/along923928784/go_api.git

#### 1.Go Module管理依赖

#### 2.配置数据库

本项目依赖于任何网站项目都会使用的Mysql和Redis，所以你需要提前安装和启动这两个服务。

如果你是windows用户，可以快速的解决mysql和redis安装的问题,通过: PHPStudy。

本视频用几分钟教会你如何使用PHPStudy，https://www.bilibili.com/video/av64485001/

如果你是OSX或者linux的硬核用户，相必启动Mysql和Redis对你不是问题😁

#### 3.配置环境变量

> 设置环境变量，你可以参考singo框架的文档: https://singo.gourouting.com/quick-guide/set-env.html

由于每个用户的电脑环境不同，所以我们通过环境变量来改变着些容易变化的属性。

你需要复制项目根目录下的.env.example文件，然后建立.env文件，然后把内容帖进去

```ini
MYSQL_DSN="user:password@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local" # mysql连接串
REDIS_ADDR="127.0.0.1:6379" # redis地址
REDIS_PW="" # redis密码(可以不填)
REDIS_DB="" # redis数据库(可以不填)
GIN_MODE="debug" # 服务状态，开发环境不用改

```

#### Windows CMD 系统启动指令

```bash
set GOPROXY=https://mirrors.aliyun.com/goproxy/
set GO111MODULE=on

go run main.go
```

#### Windows Powershell 系统启动指令

```bash
$env:GOPROXY = 'https://mirrors.aliyun.com/goproxy/'
$env:GO111MODULE = 'on'

go run main.go
```

#### linux / OSX 系统启动

```bash
export GOPROXY=https://mirrors.aliyun.com/goproxy/
export GO111MODULE=on

go run main.go
```

## 神奇的接口文档

服务启动后: http://localhost:8080/swagger/index.html

接口文档位于项目swagger目录下。请阅读目录内的文档