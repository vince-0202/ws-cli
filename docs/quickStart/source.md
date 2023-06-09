# How to use
[English](https://github.com/vince-0202/ws-cli/blob/main/docs/quickStart/source.md) | [Chinese](https://github.com/vince-0202/ws-cli/blob/main/docs/chinese/quick-start.md)
## Install
```shell
go install github.com/vince-0202/ws-cli@v0.1.0
````
Download the binary package and run it locally. When you see the startup screen, the program runs successfully.
![](https://github.com/vince-0202/ws-cli/blob/main/docs/quickStart/images/allCommands.png)
WS-CLI includes all kinds of websocket testing tools.
## running websocket server
Enter an address to quickly start a websocket server locally.
```shell
server [path]
```
![](https://github.com/vince-0202/ws-cli/blob/main/docs/quickStart/images/startWsServer.png)

## running websocket client
Enter a websocket server address to quickly start a websocket client and connect to the destination server.
```shell
client [server path]
```
![](https://github.com/vince-0202/ws-cli/blob/main/docs/quickStart/images/startWsClient.png)
Connect to the websocket server program which we started earlier.
## send a message to client
send a message to client
```shell
send 'hello ws-cli'
```
client receive the message.
![](https://github.com/vince-0202/ws-cli/blob/main/docs/quickStart/images/clientReceivedHello.png)
## send a message to server
```shell
send 'hello vince'
```
![](https://github.com/vince-0202/ws-cli/blob/main/docs/quickStart/images/serverReceivedHello.png)