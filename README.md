# ch1
* 1.2 示例
    os.Args[1:]
        变量
        作用:   把参数装到[]string中
        返回值: []string
    strings.Join()
        函数
        作用:   把切片里面的N个string元素, 用分隔符连接
        形参1:  []string
        形参2:  string类型, 分隔符
        返回值: string
* 1.3 示例
    os.Stdin
        变量: 是一个打开的文件句柄, fd=0
        类型: os.*File结构类. 实现了Read方法, 既是又是Reader接口类
    bufio.Newscanner()
        函数
        作用:   生成对fd(stdin, file)进行扫描的扫描器        
        形参:   io.Reader接口类 
        返回值: Scanner结构类 得到扫描器操作句柄
    scanner.Scan()
        方法
        作用:   将系统缓冲区读到自己的临时缓冲区
        返回值: bool
    scanner.Text()
        方法 
        作用:   取出临时缓冲区的内容
        返回值: string
    os.Open()
        函数      
        作用:    打开指定路径的文件,返回一个文件描述符,表示打开一个文件句柄, 通过句柄操作文件.
        形参:    是字符串，即文件路径
        返回值1: *os.File结构类,代表一个打开的文件句柄或者说对文件的抽象.
        返回值2: err
    os.ReadFile()
        函数
        作用:    一次性把整个文件读取到大内存缓冲区
        形参:    string
        返回值1: 是[]byte
        返回值2: err
    strings.Split()
        函数
        作用:   把一个string, 按照分割符分割成小的string, 放在切片中
        形参1:  string, 要分割的string
        形参2:  string, 分隔符
        返回值: []string
        strings.Split和strings.Join互为逆操作
* 1.5 示例
    fmt.Fprintf()
        函数
        作用:   把报错写入文件, stderr, stdin
        形参1:  实现io.Writer接口的实例, 可以是文件, stdin, stderr
        形参2:  输出的格式字符串, 和fmt.Printf一样
        形参3:  格式化的值
        返回值: 写入的字节数
    http.Get()
        函数
        作用:    发起http url的请求, 得到响应
        形参:    url, string类型
        返回值1: http.response结构类实对
        返回值2: err
    resp.Body
        response结构类的body属性
        类型:  io.ReadCloser,是接口类, 继承了Reader接口和Closer接口, 定义了Read方法和Close方法. 也就是说任何实现了Read和Close方法的类型, 都可以作为resp.Body的值
        作用:  是一个Reader. 可以理解为"读取句柄"或"数据流通道". 它连接这底层的socket, 类似"网络文件的读取口". 如果不去调用Read方法, 响应体内容会留在操作系统的tcp接收缓冲区, go不会自动读到内存中. 响应头是会被立即读到内存的.
        resp.Body底层是一个实现了这两个方法的具体结构体.
        resp.Body的动态类型(runtime)是 *http.body 结构体
        resp.Body的静态类型(代码层面)是io.ReadCloser
        参考类比:
        | 场景      | 读取方式                           | 通道对象类型                    |
        | --------- | --------------------------------- | -------------------------------| 
        | 打开文件   | `os.Open()` → `file.Read()`       | `*os.File`（实现了 `Reader`）   |
        | 打开响应体 | `http.Get()` → `resp.Body.Read()` | `*http.body`（实现了 `Reader`） |
    io.ReadAll()
        函数
        作用:    从操作系统的tcp接收缓冲区, 把响应体一口气读到内存, 适合小数据
        形参:    io.Reader接口类, 调Read方法, 实例必须实现read方法
        返回值1: []byte
        返回值2: err
    resp.Body.Close()
        方法
        作用: 关闭 Go 程序中的 HTTP 响应体读取通道, 释放连接资源, 释放TCP缓冲区, 归还Keep-Alive连接, 把连接归还给连接池进行复用. 或者关闭TCP socket, 释放资源.
    io.Copy()
        函数
        作用:     io.Copy(dst, src) 从src读, 并且写入dst, 边读边写，节省内存，适合大数据或流式处理场景
        形参1:    Writer接口类, 传入的实例必须实现write()方法
        形参2:    Reader接口类, 传入的实例必须实现read()方法
        返回值1:  成功拷贝多少个字节
        返回值2:  err
    strings.HasPrefix()
        函数
        作用：  判断字符串, 是否是以xxx前缀开头
        形参1： string, 被判断的字符串
        形参2： string, 前缀
        返回值：bool
    strings.Repeat()
        函数
        作用:   重复输出连续的执行字符
        形参1:  string, 重复的字符
        形参2:  int, 重复的次数
        返回值: string
* 1.6 示例
    time.Now()
        函数
        作用:   取当前本地时间, 表示一个时间点
        形参:   无形参
        返回值: Time结构类实例, 表示当前时间点
    time.Since()
        函数
        作用:   计算从时间点t到当前时间过去了多长时间, 是一个时间段
        形参:   time.Time结构类实例, 是一个时间点
        返回值: time.Duration, 取时间间隔, 默认是纳秒
    time.Duration.Seconds()
        方法
        作用:   把时间间隔换算成秒
        形参:   time.Duration 结构类实例
        返回值: float64
    ch chan<-string
        变量
        类型:  chan<-string 只写通道, 写的元素类型是string.
        ch是标识符, chan是通道类型, <- 通道的方向:只读, string 通道里的元素是string
    io.Discard
        变量
        作用: 丢弃型写入器, 写入这里的内容会被丢弃, 是一个黑洞
        类型: io.Writer接口类. 声明了Write方法. 底层实例如果实现了Write方法, 就可以赋值给io.Discard
    fmt.Sprint()
        函数
        作用:   把多个参数简单拼接成一行字符串, 每个值之间没有空格, 结尾没有\n, 没有占位符
        形参:   多个参数, 用逗号分隔
        返回值: string
    fmt.Sprintln()
        函数
        作用:   把多个参数拼成字符串并加换行, 每个值之间有空格, 结尾有\n
    fmt.Sprintf()
        函数
        作用:   格式化字符串拼接, 用占位符拼接
    http Keep Alive
        作用: 在 同一个 TCP 连接 上，可以发送多个 HTTP 请求和接收响应，而不是每个请求都需要新建立一个连接。减少 TCP 三次握手和四次断开的过程, 因为这些过程都会消耗大量的系统资源, 尤其是高并发. 
        工作: (1)首次请求时，会建立 TCP 连接;(2)如果服务器支持 Keep-Alive，并且没有关闭该功能，客户端和服务器可以继续使用这个连接;(3)客户端可以发起 多个请求，服务器会在同一个 TCP 连接上 返回多个响应，直到连接被关闭;(4)默认超时：大部分服务器和浏览器会在某个时间内保持连接活跃，超过时间后，连接会被关闭
        设置: (1)服务端: Connection: Keep-Alive 头部：表示服务器支持复用连接. 服务器通常会设置一个 最大请求数 或 最大空闲时间 来限制复用连接的生命周期. (2)客户端: 客户端会遵循服务器的超时设置，并且在请求发送时，会传递一个 Connection: Keep-Alive 头部来表示自己希望复用连接。
* 1.7 示例
    mu:=sync.Mutex
        结构体
        作用: 声明一个互斥锁的实例. 通过加锁和释放锁, 来确保并发访问的代码段在任意时刻只能由一个协程执行，从而避免了并发时的数据竞争问题。
    mu.Lock()
        方法
        作用: 一个协程要来拿互斥锁mu, 拿到后上锁, 其他协程就拿不到mu了,也就执行不到下面的逻辑
    mu.Unlock()
        方法
        作用: 当第一个协程执行到 mu.Unlock() 时，锁被释放，其他被阻塞的协程才能依次获得锁并继续执行
    
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
    http.ListenAndServe()
        函数
        作用:
        形参1:
        形参2:
        返回值: err
    
    ServerMux
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
        来了新连接只是启动了一个 goroutine，调用
        ```go
            来了新连接只是启动了一个 goroutine，调用
        ```
        这个 handler 其实就是 mux（或你自己传入的 ServeMux），它不会新建 mux，只是用已经注册好的 mux 去查找匹配路径，并分发给你注册的 handler。




    