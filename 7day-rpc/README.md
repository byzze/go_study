# RPC

## 总结
codec编解码器，gob解码类型，预检协议|请求头|请求体|请求头|请求体
创建一个client，开启协程接受数据，每次发送的消息封装为call，将要发送的call放入通道中，一一取出去调用write发送数据，协程接受数据解析header，body，并将call放入调用成功的通道当中。
通过反射注册方法，供client传的参数查找调用
封装client处理时间变量（client，server端都有做超时处理），根据变量将处理时间终止，防止处理时间过长损耗资源
支持http协议(get,post)
负载均衡(轮询，随机，权重，hash)
服务注册与发现（上报，心跳）

注意：
gob.Encode, gob.Decode, 方法非线程安全，且Decode读取数据时，若读取的数据为空，则会阻塞线程，可使用DecodeValue方法可以不阻塞线程

