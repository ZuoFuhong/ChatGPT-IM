## ChatGPT 聊天机器人的 IM 系统

实时通讯 IM 系统接入 ChatGPT completion API 的绝佳实践案例，在会话中接入 OpenAI 最新发布的 Chat completion API，可以回答各种问题，像 ChatGPT 一样灵性。

![架构设计](./doc/im/ChatGPT-IM.jpg)

**功能特性**

- 1.`A` 新增好友列表
- 2.`A` 新增 web 模块，提供 `RESTFul API` 接口，可以用来注册设备、创建群组、添加用户、添加好友等。
- 3.`A` Web 聊天室在 `frontend/web-im` 目录中提供一个 H5 实现的聊天室（仅测试过：chrome 浏览器）。
- 4.`A` 使用嵌入数据库 Bolt 作为数据源，免去用户环境搭建，同时内置演示数据，方便体验
- 5.`A` 新增接入 OpenAI 云端 API 服务，实现类似 ChatGPT 的聊天机器人，可以回答各种问题
- 6.`A` 新增接入 OpenAI 最新发布的 Chat completion API，回答问题像 ChatGPT 一样灵性
- 7.`A` 新增接入 OpenAI 最新发布的 Whisper API 实现语音转文字
- 8.`A` 新增 ChatGPT IM 桌面端，使用 tauri 构建的桌面应用程序

### 演示 Demo

**B站视频**：https://www.bilibili.com/video/BV1eM41147jN/

**Web 网页端**
![web-im](./doc/im/cover-openai.jpg)

测试数据：

```shell
# 熊大
http://localhost:63342/go-IM/example/web-im/index.html?uid=1629770111088857088

# 熊二
http://localhost:63342/go-IM/example/web-im/index.html?uid=1629770216865009664

# 光头强
http://localhost:63342/go-IM/example/web-im/index.html?uid=1629769779311022080
```

**桌面应用**

![app-im](./doc/im/desktop-app.jpg)

### 开发

1.Server 端

```sh
cd backend
# 默认运行在 127.0.0.1:8080
go run main.go
```

2.Web 端

```sh
cd frontend/web-im

# 浏览器打开 index.html
http://localhost:63342/go-IM/example/web-im/index.html?uid=1629770111088857088
```

3.桌面端

```sh
# 开发模式
cargo tauri dev

# 构建安装包
cargo tauri build
```

### 博客

- [OpenAI Chat completion API 入门指南](https://www.cnblogs.com/marszuo/p/17177286.html)

### License

This project is licensed under the [Apache 2.0 license](https://github.com/ZuoFuhong/ChatGPT-IM/blob/master/LICENSE).