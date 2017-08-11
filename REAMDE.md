# GRPC

## 1. Install

```shell
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

## 2. Service definition

```go
service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string greeting = 1;
}

message HelloResponse {
  string reply = 1;
}
```



1. **Unary RPCs** where the client sends a single request to the server and gets a single response back, just like a normal function call.

   ```go
   rpc SayHello(HelloRequest) returns (HelloResponse){
   }
   ```

   ​

2. **Server streaming RPCs** where the client sends a request to the server and gets a stream to read a sequence of messages back. The client reads from the returned stream until there are no more messages.

   ```
   rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){
   }
   ```

   ​

3. **Client streaming RPCs** where the client writes a sequence of messages and sends them to the server, again using a provided stream. Once the client has finished writing the messages, it waits for the server to read them and return its response.

   ```
   rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {
   }
   ```

   ​

4. **Bidirectional streaming RPCs** where both sides send a sequence of messages using a read-write stream. The two streams operate independently, so clients and servers can read and write in whatever order they like: for example, the server could wait to receive all the client messages before writing its responses, or it could alternately read a message then write a message, or some other combination of reads and writes. The order of messages in each stream is preserved.

   ```
   rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){
   }
   ```

   ​

grpc-go 中的 [examples](https://github.com/grpc/grpc-go/blob/master/examples/route_guide/README.md), 样例完整描述了以上4中 service 定义和使用。



## 3. Grpc-go Auth

Grpc-go 支持的认证方式 参见 [grpc auth support](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md)

1. TLS
2. Google Token
3. OAuth https://github.com/grpc/grpc-go/blob/master/credentials/oauth/oauth.go




## 4. issue


在Protobuf 3的协议中，去除了对于解析字段是否会出现的判断，会统一将会出现的字段设置成默认值，这样一来就丧失了对于某些字段是否出现的判断，目前的解决方式是采用 wrapper 包装一层，使用起来比较不方便。

场景： 如果对一下用户进行部分更新的话，比如 将 desc 字段清空或者 将 actived 设置成 false，则传递这些参数在 v3 版本中就没有办法判断是用户端传送过来的，还是设置的默认值，需要应用层进行提早规避。

```shell
person {
	id   int 
	desc string 
	name string
	sex  int 
	actived bool
}
```


> Note that for scalar message fields, once a message is parsed there's no way of telling whether a field was explicitly set to the default value (for example whether a boolean was set to false) or just not set at all: you should bear this in mind when defining your message types. For example, don't have a boolean that switches on some behaviour when set to false if you don't want that behaviour to also happen by default. Also note that if a scalar message field is set to its default, the value will not be serialized on the wire.

* [issue359](https://github.com/google/protobuf/issues/359)  
* [wrappers.proto](https://github.com/google/protobuf/blob/master/src/google/protobuf/wrappers.proto#L84)






