## http client 请求server
执行 
```
kubectl -n mtls-demo apply -f redis-client-no-tls.yaml
kubectl -n mtls-demo apply -f redis-server-no-tls.yaml
kubectl -n mtls-demo port-forward svc/redis-client-normal 8080:8080
curl http://localhost:8080/
```
**结论：正常**
## http client 请求server，server开启mTLS，client不开启mTLS
修改redis-server-normal的host指向redis-server-tls
报错信息如下
```
Connection reset; nested exception is redis.clients.jedis.exceptions.JedisConnectionException: java.net.SocketException: Connection reset
```
**结论：连接不上**

## http client 请求server，client开启mTLS，server不开启mTLS
结论：异常，报错信息如下
```
Caused by: java.net.SocketTimeoutException: Read timed out
	at java.base/java.net.SocketInputStream.socketRead0(Native Method) ~[na:na]
	at java.base/java.net.SocketInputStream.socketRead(Unknown Source) ~[na:na]
	at java.base/java.net.SocketInputStream.read(Unknown Source) ~[na:na]
	at java.base/java.net.SocketInputStream.read(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketInputRecord.read(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketInputRecord.readHeader(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketInputRecord.decode(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLTransport.decode(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketImpl.decode(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketImpl.readHandshakeRecord(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketImpl.startHandshake(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketImpl.ensureNegotiated(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketImpl$AppOutputStream.write(Unknown Source) ~[na:na]
	at redis.clients.jedis.util.RedisOutputStream.flushBuffer(RedisOutputStream.java:52)
```
## http client 请求server，client和server都开启mTLS
**结论：正常**

## http client 请求网格中的http server，server不开启mTLS
**结论：正常**
## http client 请求网格中的http server，server开启mTLS 宽松模式
**结论：正常**
## http client 请求网格中的http server，server开启mTLS 严格模式
**结论：请求正常**，因为Istio 默认情况下允许非网格客户端（即没有 Sidecar 的客户端）通过普通的 HTTP 访问网格内的服务。这是因为非网格客户端无法参与 mTLS 握手，因此无法使用 mTLS 访问网格内的服务。如果你想要禁止非网格客户端访问网格内的服务，可以通过配置 Istio 的 ingressgateway 来实现。

## http client 请求网格中的http server，client开启mTLS，server不开启mTLS

**结论：异常**，报错信息如下
```
Caused by: java.net.SocketTimeoutException: Read timed out
	at java.base/java.net.SocketInputStream.socketRead0(Native Method) ~[na:na]
	at java.base/java.net.SocketInputStream.socketRead(Unknown Source) ~[na:na]
	at java.base/java.net.SocketInputStream.read(Unknown Source) ~[na:na]
	at java.base/java.net.SocketInputStream.read(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketInputRecord.read(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketInputRecord.readHeader(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketInputRecord.decode(Unknown Source) ~[na:na]
```
## http client 请求网格中的http server，client和server都开启mTLS，严格模式
**结论：异常**，报错信息如下
```
Caused by: java.net.SocketTimeoutException: Read timed out
	at java.base/java.net.SocketInputStream.socketRead0(Native Method) ~[na:na]
	at java.base/java.net.SocketInputStream.socketRead(Unknown Source) ~[na:na]
	at java.base/java.net.SocketInputStream.read(Unknown Source) ~[na:na]
	at java.base/java.net.SocketInputStream.read(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketInputRecord.read(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketInputRecord.readHeader(Unknown Source) ~[na:na]
	at java.base/sun.security.ssl.SSLSocketInputRecord.decode(Unknown Source) ~[na:na]
```


