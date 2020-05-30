## Go语言开发基于websocket的聊天（IM）系统

本项目是由[gim](https://github.com/alberliu/gim)项目fork而来，在此基础上进行了二次开发，目的在于实践`IM`系统的`单用户多设备支持，离线消息同步`。

开发的目标是快速梳理[gim](https://github.com/alberliu/gim)的核心流程，并仿写实现`单用户多设备支持，离线消息同步`的逻辑层。在开发的过程中，
砍掉了gRPC、TCP服务端、Redis缓存等模块。同时，在`go.mod`中仅依赖了几个必须的第三方包，其余均由纯go实现。

> 新增功能点

- 1.新增了`web`模块，提供`RESTFul API`接口，可以用来注册设备、创建群组、添加用户等。

- 2.`web`服务支持静态文件处理，在`example`目录中提供一个H5客户端，可以调试服务端`websocket`，收发消息。

### Development

```sh
# clone the project
git clone git@github.com:ZuoFuhong/go-IM.git

# update dependency
go mod tidy

# build the project
make

# init the database
./doc/init_table.sql
```

### License

The project is licensed under the MIT license.
