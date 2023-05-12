# WS-CLI使用手册
下载二进制包并在本地运行，看见启动画面表示程序运行成功。

![](/docs/quickStart/images/allCommands.png)
包括各类websocket测试工具。
## 开启一个websocket服务器
输入一个地址在本地快速启动一个websocket服务器
```shell
server [path]
```
![](/docs/quickStart/images/startWsServer.png)

## 开启一个websocket客户端
输入一个websocket服务端地址，快速启动一个websocket客户端并连接到目的服务器。
```shell
client [server path]
```
![](/docs/quickStart/images/startWsClient.png)
连接到前面启动到websocket服务端程序
## 向客户端发送一条消息
向客户端发送一条消息
```shell
send 'hello ws-cli'
```
客户端接收消息
![](/docs/quickStart/images/clientReceivedHello.png)
## 向服务端发送一条消息
```shell
send 'hello vince'
```
![](/docs/quickStart/images/serverReceivedHello.png)