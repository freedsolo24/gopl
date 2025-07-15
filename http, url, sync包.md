# http 包
http.Get()
    函数
    作用: 发起http url的请求, 得到响应
    形参:  url, string类型
    返回值1: http.response结构类实对
    返回值2: err

http.ListenAndServe()
    函数
    作用:   监听
    形参1:  string, 监听的地址和端口
    形参2:  handler操作函数
    返回值: err

resp.Body
    response结构类的body属性
    类型:  io.ReadCloser,是接口类, 继承了Reader接口和Closer接口, 定义了Read方法和Close方法. 也就是说任何实现了Read和Close方法的类型, 都可以作为resp.Body的值
    作用:  是一个Reader. 可以理解为"读取句柄"或"数据流通道". 它连接这底层的socket, 类似"网络文件的读取口". 如果不去调用Read方法, 响应体内容会留在操作系统的tcp接收缓冲区, go不会自动读到内存中. 响应头是会被立即读到内存的.
    resp.Body 底层是一个实现了这两个方法的具体结构体.
    resp.Body 的动态类型(runtime)是 *http.body 结构体
    resp.Body 的静态类型(代码层面)是io.ReadCloser
    参考类比:
    | 场景      | 读取方式                           | 通道对象类型                    |
    | --------- | --------------------------------- | -------------------------------| 
    | 打开文件   | `os.Open()` → `file.Read()`       | `*os.File`（实现了 `Reader`）   |
    | 打开响应体 | `http.Get()` → `resp.Body.Read()` | `*http.body`（实现了 `Reader`） |
    resp.Body 只能读取一次, 如果先用 json.NewDecoder(resp.Body).Decode(&comic) 读取响应体，然后又用 io.ReadAll(resp.Body) 读取，导致 body 为空 

resp.Body.Close()
    方法
    作用: 关闭 Go 程序中的 HTTP 响应体读取通道, 释放连接资源, 释放TCP缓冲区, 归还Keep-Alive连接, 把连接归还给连接池进行复用. 或者关闭TCP socket, 释放资源.
    (1) 如果不调用 Close，可能会导致资源泄漏（如文件描述符未释放），尤其是在高并发场景下，可能耗尽系统资源
    (2) 在高负载程序中，积累的未关闭连接可能导致系统资源耗尽（例如 "too many open files" 错误）
    (3) HTTP/1.1 默认启用 keep-alive，允许重用 TCP 连接以提高性能. 不调用 Close 可能阻止连接被放回连接池，降低效率

http.StatusOK
    http 包定义的常量
    StatusOK = 200

resp.StatusCode
    成员属性
    Response 结构体里面的 StatusCode 成员属性

resp.Status
    成员属性
    Response 结构体里面的 Status 成员属性, 用来描述 http 响应状态

http.HandleFunc()
    函数
    作用: 注册一个处理器函数handler, 当用户访问路径url开头的时候, 会调用这个handler来处理这个http请求.
    形参1: string, 匹配的路由前缀
    形参2: handler函数的签名. func(http.ResponseWriter, *http.Request)
            w http.ResponseWriter  
            接口类.赋值给他的实例需要右Write方法.用于向客户端写http响应, 可以通过w.Write() 或 fmt.Fprintf(w,...) 把数据返回给浏览器
            r *http.Request
            结构类.表示客户端的请求,里面包含所有的http信息
    无返回值
```go
http.HandleFunc("/", handler)
func handler(w http.ResponseWriter, r *http.Request) {}                  // 我定义了一个普通的handler函数
type HandlerFunc func(ResponseWriter, *Request)                          // Go标准库把它封装成HandlerFunc类型
type Handler interface { ServeHTTP(ResponseWriter, *Request) }           // handler接口声明ServeHTTP方法
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) { f(w, r) } // HandlerFunc类型的实例, 实现了ServeHTTP方法, 就实现了Handler接口

http.HandleFunc("/", handler)
// 我传参给HandleFunc函数
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
if use121 {
    DefaultServeMux.mux121.handleFunc(pattern, handler)
// ServeMux实例, 调handlerFunc方法,传"/"和handler函数
func (mux *serveMux121) handleFunc(pattern string, handler func(ResponseWriter, *Request)) { mux.handle(pattern, HandlerFunc(handler)) }
// 把我写的handler函数转成HandlerFunc类型,就是转成Handler接口
// ServeMux实例, 调handle方法, 传 "/" , 我的函数已经是Handler接口了
func (mux *serveMux121) handle(pattern string, handler Handler) { e := muxEntry{h: handler, pattern: pattern} 
            mux.m[pattern] = e
}
// 把我定义的handler函数, 存到了mux实例的m键对应的值
http.ListenAndServe("localhost:8000", nil)
// 我开始监听
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}
// server是一个实例, 调ListenAndServe方法
func (s *Server) ListenAndServe() error { 
    ln, err := net.Listen("tcp", addr)
    return s.Serve(ln)
        }
// 把tcp和地址传给Listen函数, 返回ln, 在调用Serve
func (s *Server) Serve(l net.Listener) error { 
    for {
            rw, err := l.Accept()
            c := s.newConn(rw)
            connCtx := ctx
            go c.serve(connCtx)
        }
    }
// 每个连接新建一个 goroutine，执行 conn.serve
func (c *conn) serve(ctx context.Context) {
    serverHandler{c.server}.ServeHTTP(w, w.req)
    // serverHandler 是一个结构体, 里面的属性是之前定义的Sever结构体,里面的属性有{地址}
    // 实现了ServeHTTP的方法
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
    handler := sh.srv.Handler
    if handler == nil { handler = DefaultServeMux }  // 就是我们用 HandleFunc 注册的 mux
    handler.ServeHTTP(rw, req)
}
// http.ListenAndServe(addr, nil) 传入的是 nil，则使用默认的 DefaultServeMux
// DefaultServeMux 就是你用 http.HandleFunc() 注册 handler 的地方
// 这个handler就是我写的handler, 因为我写的handler被变成了HandlerFunc类型，这个类型又实现了Handler接口，可以调ServeHTTP方法
// 这个handler底层是ServeMux
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
        var h Handler 
        h, _ = mux.mux121.findHandler(r)
        h.ServeHTTP(w, r)
    }
// 从 mux 路由表中查找匹配的 Handler，调用对应的 handler 处理请求
// 这时候的h就是我注册的handler
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
f(w, r) }
// 我定义的handler最终被调用
```


http Keep Alive
    作用: 在 同一个 TCP 连接 上，可以发送多个 HTTP 请求和接收响应，而不是每个请求都需要新建立一个连接。减少 TCP 三次握手和四次断开的过程, 因为这些过程都会消耗大量的系统资源, 尤其是高并发. 
    工作: (1)首次请求时，会建立 TCP 连接;(2)如果服务器支持 Keep-Alive，并且没有关闭该功能，客户端和服务器可以继续使用这个连接;(3)客户端可以发起 多个请求，服务器会在同一个 TCP 连接上 返回多个响应，直到连接被关闭;(4)默认超时：大部分服务器和浏览器会在某个时间内保持连接活跃，超过时间后，连接会被关闭
    设置: (1)服务端: Connection: Keep-Alive 头部：表示服务器支持复用连接. 服务器通常会设置一个 最大请求数 或 最大空闲时间 来限制复用连接的生命周期. (2)客户端: 客户端会遵循服务器的超时设置，并且在请求发送时，会传递一个 Connection: Keep-Alive 头部来表示自己希望复用连接。

# sync 包
mu:=sync.Mutex
    结构体
    作用: 声明一个互斥锁的实例. 通过加锁和释放锁, 来确保并发访问的代码段在任意时刻只能由一个协程执行，从而避免了并发时的数据竞争问题。
mu.Lock()
    方法
    作用: 一个协程要来拿互斥锁mu, 拿到后上锁, 其他协程就拿不到mu了,也就执行不到下面的逻辑
mu.Unlock()
    方法
    作用: 当第一个协程执行到 mu.Unlock() 时，锁被释放，其他被阻塞的协程才能依次获得锁并继续执行

ServerMux 是标准库中 http 包的路由器（Router），它是一个 多路复用器（Multiplexer），ServeMux 是根据 HTTP 请求的 路径（URL） 来决定调用哪个 handler。
ServeMux 会根据你注册的路径和 Handler 映射，在 接收到 HTTP 请求时，查找请求的路径对应的 Handler，然后调用该 Handler 的 ServeHTTP 方法来处理请求。
当我调用 ```http.HandleFunc("/", handler)```，背后发生的是把我写的 handler 函数，注册进 mux 的 map
```mux.m[pattern] = muxEntry{h: handler, pattern: pattern}```
注册发生时，mux（即 DefaultServeMux）就已经存在，它是常驻的
```go
http.ListenAndServe(":8000", nil)
if handler == nil {
    handler = DefaultServeMux
}
```
所以 ServeHTTP 方法会用 默认 mux 来分发请求。
来了新连接只是启动了一个 goroutine，调用这个 handler 其实就是 mux（或你自己传入的 ServeMux），它不会新建 mux，只是用已经注册好的 mux 去查找匹配路径，并分发给你注册的 handler。

ch chan<-string
    变量
    类型:  chan<-string 只写通道, 写的元素类型是string.
    ch是标识符, chan是通道类型, <- 通道的方向:只读, string 通道里的元素是string

# url 包
url.QueryEscape()
    函数
    作用: 将传入的字符串 []string 构建成 HTTP 请求的查询参数. 它会把不能直接出现在 URL 查询参数中的字符（如空格、中文、特殊符号）进行转义.将输入字符串 s 编码为 URL 查询参数安全的格式
    ' '     + 或 %20
    ?       %3F
    &       %26
    =       %3D
    :       %3A
    /       %2F
    例如: url.QueryEscape("repo:golang/go is:open json decoder") -> "repo%3Agolang%2Fgo+is%3Aopen+json+decoder"
    Escape: 在编程中, 指将特殊字符转换为安全形式

url.PathEscape 与 url.QueryEscape 区别
    (1) url.QueryEscape 用于查询参数（?q=...），对更多字符进行编码（如 + 编码为 %2B）
    (2) url.PathEscape 用于 URL 路径（如 /path/to/resource），对保留字符的编码更宽松（如 + 不编码）。

url.Values()
    类型 map[string][]string 的别名
    作用: 用于构建和管理 URL 的查询参数（即 URL 中 ? 后面的部分）
    键是字符串 值是字符串切片
v.Set()
    方法
    作用: 用于构建 url 查询参数. 将某个键设置为单一值，覆盖该键的任何现有值. 如果键已存在，Set 会替换其值（不像 Add，后者追加值
    设置键值对: 如 q, per_page, page, sort, order
v.Add()
    方法
    作用: 用于构建 url 查询参数. 如果键已存在，Add 会追加值
v.Encode()
    方法
    作用: 用于构建 url 查询参数. 将 url.Values 转换为 URL 编码的查询字符串
url.Parse()
    函数
    作用: 将字符串 URL 解析为 url.URL 结构体，方便操作 URL 的各部分（如 Scheme, Host, Path, RawQuery）
    解析基础 URL
u.RawQuery
    RawQuery 是 url 结构体中的成员属性
    作用: 设置或获取 URL 的查询部分（? 后面的内容）

```go
q := url.QueryEscape(strings.Join(terms, " "))
v := url.Values{}
v.Set("q", q)
v.Set("per_page", "100")
v.Set("sort", sortBy)
v.Set("order", order)
// q=json+decoder+repo:golang/go+is:open+is:issue, per_page=100, sort=created, order=asc
u:=url.Parse("https://api.github.com/search/issues")
u.RawQuery=v.Encode()
```

r.URL.Path
    r:    http.Request结构类的实例
    URL:  Request结构类里面的一个属性, 也是结构类
    Path: 是URL结构类里面的一个属性, 类型是string
```bash
请求地址：http://localhost:8080/hello/world?x=1#top
r.URL.Path      = "/hello/world"
r.URL.RawQuery  = "x=1"
r.URL.Fragment  = "top"
```

r.Header
    作用:  属性, 类型map[string][]string

r.ParseForm()
    方法
    作用:   解析来自 URL 的 query 参数（比如 ?name=david&age=18）和 POST 表单数据，统一放到 r.Form 这个 map 里
    形参:   无形参
    返回值: err

r.Form()
    是一个 map[string][]string，可以遍历所有表单字段
    类型: map[string][]string
            key是string, value是切片里面是string元素, 表示多个字段, 每个字段有很对的值
    例如: 
```go
    http://localhost:8000/?name=David&hobby=Go&hobby=Music
    ["name"]["David"]
    ["hobby"]["Go" "Music"]
```