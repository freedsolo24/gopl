```bash
var m map[string]int        # 只是声明, 不立即分配内存. 不可直接赋值, 没有底层数据结构. 得到一个空map, 后面在动态赋值. 作用: 延迟初始化, 全局变量声明
m["a"] = 1                  # panic
m = make(map[string]int)    # 在这里初始化
---
n := make(map[string]int)   # 是初始化, 可以直接赋值, 有底层数据结构, 可以字面量赋值 map[...]{...}, 动态增加键值对
n["a"] = 1 

var s []int                 # 只是声明切片, 不立即分配内存. 不可直接赋值, 没有底层数据结构. 
s = append(), s = make()    # 切片可以append自动分配内存
---
# {...}字面量定义 必须是复合类型, 复合类型前面必须加类型: 数组,切片,map,结构体
---
var title []struct{Title string}  # 匿名结构体定义
---
type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue `json:"items"`
}
# 核心点是(1)内存效率, (2)共享底层数据, (3)空值处理
# (1) Items是一个切片, 切片里每一个元素指向Issue结构体的指针. Items切片只存储指向这些实例的指针, 当这个结构体被操作时, 只传递指针节省内存.
# (2) 指针允许多个地方引用同一个Issue实例, 如果某个地方修改 Issue 字段, 所有引用该指针的地方都会反应这个变化
# (3) []*Issue 允许 Issue 为 nil, 表示元素为初始化或无效. []Issue 每个元素都会初始化为 Issue 的零值, 无法表示"缺失状态"
---
type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue `json:"items"`
}
var result IssuesSearchResult
# 声明并初始化结构体变量. 自动分配零值内存空间. TotalCount 初始化为0, Items 初始化为 nil
# 显示初始化
result := IssuesSearchResult {
    TotalCount: 0,
    Items:      []*Issue{}
}
# 使用 new 初始化, 返回的是指针
result := new(IssuesSearchResult)
# 短格式声明
result := IssuesSearchResult{}
# 如果结构体包含 slice 或 map 字段，且需要非 nil 值，则需要手动初始化这些字段
result := IssuesSearchResult {
    Items: make([]*Issue, 0), // 初始化为空切片
}
# 数组, 结构体是值类型, 不需要 make(), 可以使用 new()
# slice, map, channel 引用类型, 才需要 make(). make() 的本质是构造标头值(len, cap)
```
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
        形参1:  实现io.Writer接口的实例, 可以是文件, stdin, stderr, bytes.buffer
        形参2:  输出的格式字符串, 和fmt.Printf一样
        形参3:  格式化的值
        返回值: 写入的字节数
    fmt.Fprintf()和fmt.Printf()的区别
        Printf:  输出的是终端
        Fprintf: 输出的是实现了Writer接口的实例
    http.Get()
        函数
        作用: 发起http url的请求, 得到响应
        形参:  url, string类型
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
        resp.Body 只能读取一次, 如果先用 json.NewDecoder(resp.Body).Decode(&comic) 读取响应体，然后又用 io.ReadAll(resp.Body) 读取，导致 body 为空 
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
        作用:   监听
        形参1:  string, 监听的地址和端口
        形参2:  handler操作函数
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
        来了新连接只是启动了一个 goroutine，调用这个 handler 其实就是 mux（或你自己传入的 ServeMux），它不会新建 mux，只是用已经注册好的 mux 去查找匹配路径，并分发给你注册的 handler。

* 1.7 示例体
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
# ch2
* 1.6示例
    strconv.ParseFloat
        函数
        作用: 把第一个string类型的形参, 解析成float类型的数值
        形参1: string类型
        形参2: 指定返回值是64bit, 还是32bit
        返回值1: float类型
        返回值2: err
    func init() { ... }
        init函数
        作用: 包中的init函数在包导入的时候, 最先执行. 在此函数中用来初始化一个数据表的变量pc[265]byte, 让pc这个数组变量执行后里面有值
        init函数不能被调用和引用
# ch3
* 3.5.4示例
    strings.LastIndex()
        函数
        作用: 判断字符串中, 指定字符的最后一个索引号
        形参1: 字符串
        形参2: 字符
        返回值: 索引id
    strings.Contains(s, substr string) bool
        函数
        作用: 字符串s, 是否包含子串
    strings.Count(s, sep string) int
        函数
        作用: 字符串s, 有几个字串的个数 
    strings.Fields(s string) []string
        函数
        作用: 把字符串s, 针对一个或多个空格, 进行分割, 返回[]string
    strings.HasPrefix(s, prefix string) bool
        函数
        作用: 字符串s, 是否有前缀子串
    strings.Index(s, sep string) int
        函数
        作用: 字符串s, 里面的sep子串, 索引id
    strings.Join(a []string, sep string) string
        函数
        作用: 把sep字串放在字符串切片里面
    bytes.Contains(b, subslice []byte) bool
        函数
        作用: 字节流b, 是否包含子字节流
    bytes.Count(s, sep []byte) int
        函数
        作用: 
    bytes.Fields(s []byte) [][]byte
        函数
    bytes.HasPrefix(s, prefix []byte) bool
        函数
    bytes.Index(s, sep []byte) int
        函数
    bytes.Join(s [][]byte, sep []byte) []byte
        函数
    bytes.Buffer
        结构体
        作用: 提供了一个可增长的字节缓冲区，你可以往里面追加内容(byte, string, 其他buffer), 内部维护一个[]byte
        实现了io.Writer接口, 可以传给Fprintf()
    buf.WriteByte()
        方法
        作用: buf实例, 调用WriteByte方法, 往buf实例后面追加一个ASCII字符
        形参: 1个byte类型的ASCII码字符
        返回值: err 
    buf.WriteString()
        方法
        作用: buf实例, 调用WriteString方法, 往buf后面追加一个字符串
        形参: 往buf实例后面追加的字符串
        返回值1: 写入字符串的长度
        返回值2: err 
    buf.String()
        方法
        作用: buf实例, 调用String方法, 把实例整个缓冲区返回成string, string是固定不变的不能修改里面的内容
        无形参
        返回值: string类型
    buf.Bytes()
        方法
        作用: buf实例, 调用Bytes方法, 把实例整个缓冲区返回成[]byte, 这样可以操作这个字符串
    strconv.Itoa()
        函数
        作用: 将int转换成string
        strconv.Itoa(123)  // "123"
    strconv.FormatInt()
        函数
        作用: 不同进制的转换
        例子: strconv.FormatInt(int64(123), 2)   // "11111011"
    strconv.Atoi()
        函数
        例子: strconv.Atoi("123")  // 123
    strconv.ParseInt()
        例子: strconv.ParseInt("123", 10, 64)
    strconv.Atoi()和strconv.ParseInt()的区别
        Atoi: 单纯的将String转成Int
        ParseInt: 灵活, 将string, 根据10,16进制转换, 转换后装进int64的类型中
    strings.SplitN()
        函数
        作用: 对指定的分隔符, 最多分成N个子串.
        形参1: 字符串string
        形参2: 分隔符
        形参3: 分成几个子串. n=0: 不分, n<0: 不限分割次数
        返回值: []string, 分出来的子串装到[]string中
# ch4
    sha256.Sum256()
        函数
        标准库, crypto/sha256 包提供
        作用: 将字符串计算对应的sha256哈希值
        形参: 类型必须是[]byte, 所以字符串要做强转[]byte(s)
        返回值: [32]byte, 返回值是256bit字串, 256/8=32, 返回值是一个数组, 里面有32个字节的字节值
        2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
    flag.String()
        函数, 标准库 flag 包提供
        作用: 设置一个string类型的命令行选项
        形参1: 命令行的参数
        形参2: 如果没有参数, 用这个默认值
        形参3: 帮助信息, 用户 -h 会提示
        返回值: 返回*string, 是一个指针 
    flag.Parse()
        函数
        作用: 调用之后才会真正解析命令行参数并把它们填入上面定义的变量里
        会将命令行中的参数分为: 
        (1) flag参数: 以 - 开头的部分, 如-algo sha512
        (2) 非flag参数: 解析完 flag 后剩下的那些参数
            ```bash
            go run main.go -algo sha512 abc
            ```
    flag.Arg()
        函数
        作用: 返回第 i 个非 flag 参数（从 flag.Parse() 剩下的参数中取）
        形参: int, 第几个参数
        返回值: string
    flag.Args()
        函数
        作用: 返回所有剩下的非 flag 参数
        形参: 无
        返回值: []string
    utf8.DecodeRune()
        函数
        作用: 输入[]byte字节流, 从一个 UTF-8 编码的 []byte 字节流中，解析出第一个合法的字符（rune 类型），并返回
        形参: []byte字节流
        返回值1: []byte字节流里面的第一个rune文字字符
        返回值2: rune文字字符的长度
    utf8.DecodeLastRune()
        函数
        作用: 输入[]byte字节流, 从一个 UTF-8 编码的 []byte 字节流中，解析出最后一个合法的字符（rune 类型），并返回
        形参: []byte字节流
        返回值1: []byte字节流里面的最后一个rune文字字符
        返回值2: rune文字字符的长度    
    utf8.EncodeRune()
        函数
        作用: 将unicode编码的, rune文字字符编码成utf8编码, 写入[]byte字节流
        形参1: utf8编码的 []byte 字节流
        形参2: unicode编码的 rune 文字字符
        返回值: 成功写入几个字节
    unicode.IsSpace()
        函数
        作用: 判断输入的rune文字字符是否属于这些列出的空白字符: '\t', '\n', '\v', '\f', '\r', ' '
        形参: rune文字字符
        返回值: true | false
    unicode.IsLetter()
        函数
        作用: 判断输入的rune文字字符是否属于字母类字符, 英文字母, 汉字, 日语等语言文字
        形参: rune文字字符
        返回值: true | false
    unicode.IsPunct()
        函数
        作用: 判断输入的rune文字字符是否属于标点符号, 有英文标点和中文标点
        形参: rune文字字符
        返回值: true | false
    unicode.IsDigit()
        函数
        作用: 判断输入的rune文字字符是否属于数字0-9
        形参: rune文字字符
        返回值: true | false
    unicode.IsControl()
        函数
        作用: 判断输入的rune文字字符是否属于控制符号, '\n', '\r', '\t', '\b'
        形参: rune文字字符
        返回值: true | false
    unicode.Is()
        函数
        作用: 用来判断一个 rune 文字字符是否属于 unicode 类别或字符范围
        形参1: Unicode 范围表，表示一个字符类别（比如字母、数字、标点等）
        形参2: 一个文字字符 rune
        返回值: true | false
        例如: unicode.Is(unicode.Latin,r)
    bufio.NewReader()
        函数
        作用: 创建一个带缓冲区的 Reader，用来高效读取标准输入(键盘输入)
        形参1: 接收一个实现了 io.Reader 接口的对象(比如: os.Stdin, 文件, 网络连接等)
        返回值: 返回Reader读句柄, 内部维护缓冲区, 可以按行、按字节、按段落 等方式读取输入，提高效率、简化读取操作
    ReadString('\n')
        方法
        作用: 按行读. 直到遇到分隔符, 返回一个string
    ReadByte('\n')
        方法
        作用: 返回[]byte, 不能解析中文, 因为一个汉字占3个字节(utf8编码), 所以ReadByte()只会拿到一部分字节, 可能乱码, 但ReadRune()会完整识别出一个文字字符.
    ReadRune()
        方法
        作用: 从Reader中, 按 utf8 解码, 读取成一个完整的 unicode 文字字符 rune.
        无形参
        返回值1: 返回 unicode 文字字符, rune 类型
        返回值2: 这个 rune 文字字符在utf8下占了多少字节
        返回值3: err
    unicode.ReplacementChar
        常量
        = '\uFFFD' 代表 invalid 码点
    utf8.UTFMax
        常量
        = 4        代表 utf8 编码最大的字节数
    io.EOF
        是标准库io包中预定义的错误值, 用于表示输入流的结束(EOF=End Of File)
    scanner.Split()
        方法
```go
func (s *Scanner) Split(split SplitFunc)
```
        作用: bufio.Scanner 类型的一个方法, 它接收一个参数：SplitFunc 类型的函数，表示“如何切分输入流”。
```go
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
```
        作用: 它作用是：如何从字节流中“切出”下一个 token（标记）。比如按行、按单词、按字符。
        用法: scanner.Split() 期望传一个函数（SplitFunc）进去，告诉它怎么切分. bufio.ScanWords 是标准库提供的一个“按单词分词”的函数变量, 告诉它我就要用单词来切分.
    bufio.NewScanner 和 bufio.NewReader 区别
    (1) NewScanner 默认按"行"读取, 也可以设置分词器按"单词"拆分, 适合逐行, 逐词处理文本. 不如 Reader 灵活
    (2) NewReader 可以逐字符, 逐行, 逐分隔符, 可以读取任意格式的输入内容, 包括字节流, utf8字符等逐个读取, 处理 unicode 字符或检查无效字符
    json.Marshal()
        函数
        作用: 序列化json字符串
        形参1: 结构体实例
        返回值1: []byte字节流
        返回值2: err
    json.MarshalIndent()
        函数
        作用: 有缩进的输出
        形参1: 结构体实例
        形参2: 前缀
        形参3: 缩进几个空格
        返回值1: []byte字节流
        返回值2: err
    json.Unmarshal()
        函数
        作用: 接收一个完整的 JSON 字节切片, 一次性解析所有内容, 装到结构体实例的指针类型变量中.
        形参1: json序列化字节流 []byte
        形参2: 装到结构体实例
        返回值: err
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
    http.StatusOK
        http 包定义的常量
        StatusOK = 200
    resp.StatusCode
        成员属性
        Response 结构体里面的 StatusCode 成员属性
    resp.Body
        Response 结构体里面的 Body 成员属性
    resp.Body.Close()
        方法
        作用: 关闭释放响应资源
        (1) 如果不调用 Close，可能会导致资源泄漏（如文件描述符未释放），尤其是在高并发场景下，可能耗尽系统资源
        (2) 在高负载程序中，积累的未关闭连接可能导致系统资源耗尽（例如 "too many open files" 错误）
        (3) HTTP/1.1 默认启用 keep-alive，允许重用 TCP 连接以提高性能. 不调用 Close 可能阻止连接被放回连接池，降低效率
    resp.Status
        成员属性
        Response 结构体里面的 Status 成员属性, 用来描述 http 响应状态
    fmt.Errorf()
        函数
        作用: 于根据指定的格式化字符串 format 和参数 a 创建一个新的 error 类型值. 返回值是一个实现了 error 接口的类型
        形参1: 格式化的字符串
        形参2: 传递要格式化的值，对应格式化字符串中的占位符
        返回值: 一个 error 接口类型的值，包含格式化后的错误消息
    json.NewDecoder()
        函数
        作用: 初始化创建 json 解码器. 设置这个解码器, 准备从 io.Reader 缓冲区里面, 流式读取 JSON 数据
        形参1: io.Reader, 可以是 resp.Body, 即实现了 io.Reader，提供 HTTP 响应的正文数据流 
        返回值: json 解码器
    jsonDecoder.Decode()
        方法
        作用: 使用之前设置的 json 解码器, 执行解码操作, 从数据流中读取 json 数据, 将解析后的 json 数据填充到结构体指针变量中, 根据结构体字段的 json 标签（如 json:"total_count"）映射字段
        形参1: 解码后, 放入的结构体实例的内存地址
        返回值: error接口
    json.Decode()和json.Unmarshal() 区别
        (1) Unmarshal() 需要一个 []byte 包含完整的 JSON 数据; Decode() 从 io.Reader（如 resp.Body）流式读取 JSON 数据
        (2) Unmarshal() 一次性读取所有数据到内存; Decoder() 流式处理, 适合大 json 数据, 适合网络流
    io.ReadAll()
        函数
        作用: 读取整个响应正文到[]byte
        形参: resp.Body
        返回值1: []byte
        返回值2: error
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
    os.ReadFile()
        函数
        作用: 读取指定路径的文件的全部内容，返回一个字节切片（[]byte）和可能的错误. 适合读取小到中型文件（例如配置文件、JSON 数据、文本文件等）的全部内容
        形参: 要读取的文件的路径（相对或绝对路径）
        返回值1: 文件内容的字节切片。如果文件为空，返回空切片（[]byte{}
        返回值2: error
        示例:
```go
data, err := os.ReadFile("example.txt")
if os.IsNotExist(err) {
    fmt.Println("File does not exist")
} else if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
} else {
    fmt.Println("File content:", string(data))
}
```
    bufio.Reader 和 os.ReadFile 区别
    (1) os.ReadFile 适合小文件的读取
    (2) bufio.Reader 适合大文件的读取, 流式读取
```go
file, err := os.Open("large.txt")
defer file.Close()
reader := bufio.NewReader(file)
```
    os.MkdirAll()
        函数
        作用: 创建指定路径的目录（包括所有必要的父目录），并设置指定的权限模式。如果目录已存在，不会返回错误。幂等性：如果目标目录或其父目录已存在，os.MkdirAll 不会报错，直接返回 nil.
        形参1: path string：要创建的目录路径（相对或绝对路径）。可以是多级路径（如 xkcd_data/comics）
        形参2: 目录的权限模式（八进制表示，例如 0755）。指定新创建目录的权限
        返回值: error
        示例:
```go
os.MkdirAll("data/comics/2025", 0755)
```
    filepath.Join()
        函数
        作用: 将多个路径片段（elem）连接成一个路径字符串，使用操作系统特定的路径分隔符（/ 在 Unix，\ 在 Windows）
        形参: 可变参数，多个路径片段（例如，目录名、文件名）
        返回值1: 连接后的路径字符串，符合当前操作系统的路径格式
        返回值2: error
        示例:
```go
path := filepath.Join("xkcd_data", "2025", "07", "571.json")
fmt.Println(path) // Unix: xkcd_data/2025/07/571.json, Windows: xkcd_data\2025\07\571.json
```
    os.Stat()
        函数
        作用: 获取指定路径的文件或目录的元信息，返回 fs.FileInfo 接口，包含文件的名称、大小、修改时间、权限等。
        形参: 文件或目录的路径（相对或绝对路径）
        返回值1: 文件或目录的元信息, Name() string, Size() int64, Mode() fs.FileMode, ModTime() time.Time, IsDir() bool, Sys() any
        返回值2: error, 操作系统失败时候的错误(文件不存在, 权限不足); 成功时返回nil 
        示例:
```go
info, err := os.Stat("example.txt")
if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    os.Exit(1)
}
fmt.Printf("Name: %s\n", info.Name())
fmt.Printf("Size: %d bytes\n", info.Size())
fmt.Printf("IsDir: %v\n", info.IsDir())
fmt.Printf("ModTime: %v\n", info.ModTime())
fmt.Printf("Mode: %v\n", info.Mode())
```
    os.WriteFile()
        函数
        作用: 将字节切片 data 写入文件 name，并设置文件权限 perm。如果文件不存在，创建新文件；如果文件已存在，覆盖原有内容。适合小到中型数据（几 KB 到 MB）
        形参1: string, 目标文件的路径（相对或绝对路径）
        型参2: []byte, 要写入的字节数据
        形参3: 0644, 文件的权限模式
        返回值: error
        示例:
```go
text := []byte("Hello, World!\n")
if err := os.WriteFile("output.txt", text, 0644); err != nil { }
```
    fmt.Errorf()
        函数
        作用: 将原始错误包装到新的错误中
```go
    err := errors.New("原始错误")
    wrapped := fmt.Errorf("出错啦: %w", err)
    fmt.Println(wrapped)  // 打印: 出错啦: 原始错误

    // 判断或提取原始错误
    if errors.Is(wrapped, err) {
        fmt.Println("匹配上原始错误了！")
    }
    unwrapped := errors.Unwrap(wrapped)
    fmt.Println("原始错误是：", unwrapped)
```
    flag.NewFlagSet()
        函数
        作用: 创建一个新的标志集, 然后为这个标志集创建多个标志. 允许为每个子命令定义独立的标志，例如 searchCmd 定义了 --terms 标志.
        形参1: string, 标志集的名称，通常是子命令的名称
        形参2: errorHandling flag.ErrorHandling：错误处理方式，有三种选项
            * flag.ContinueOnError：遇到错误继续解析，返回错误。
            * flag.ExitOnError：遇到错误调用 os.Exit(2) 退出程序。
            * flag.PanicOnError：遇到错误触发 panic。
        返回值: *flag.FlagSet：返回新创建的标志集，允许定义和解析特定子命令的标志
```go
downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)
```
    NewFlagSet.String()
        方法
        作用: 为指定的 FlagSet 定义一个字符串类型的命令行标志，返回一个指向字符串值的指针（*string）
        形参1: 标志的名称，用于命令行输入（例如 --terms 或 -terms）
        形参2: 标志的默认值，如果用户未提供该标志，则使用此值
        形参3: 标志的帮助信息，描述标志的用途，显示在 --help 或错误提示中
        返回值: *string, 指向标志值的指针, 指向的是用户输入的字符串
```go
searchTerms := searchCmd.String("terms", "", "Search terms (comma-separated)")
./xkcd-tool search --terms "sleep,insomnia"
*searchTerms 的值为"sleep, insomnia"
```
    NewFlagSet.Parse()
        方法
        作用: 解析给定的命令行参数，根据 FlagSet 中定义的标志提取值，并存储到对应的标志变量中。如果解析失败，根据 FlagSet 的错误处理模式（例如，flag.ExitOnError）处理错误
        形参: 要解析的命令行参数切片, 通常是 os.Args[1:], os.Args[2:]
        返回值: error
        示例
```go
./xkcd-tool search --terms "sleep,insomnia" extra
os.Args 是 ["./xkcd-tool", "search", "--terms", "sleep,insomnia"]
os.Args[2:] 是 ["--terms", "sleep,insomnia", "extra"]
searchCmd.Parse 识别 --terms 标志, 将 "sleep,insomnia" 赋值给 *searchTerms. 忽略非标志参数 extra, 保留在 searchCmd.Args()
```
    os.Getenv()
        函数
        作用: 获取环境变量的值。如果环境变量不存在，返回空字符串
        形参: 要查询的环境变量的值
        示例:
```go
apiKey := os.Getenv("API_KEY")
export API_KEY=abc123
```
    NewFlagSet.Usage()
        方法
        作用: 调用 flagSet.Usage() 打印子命令的用法，通常包括标志名称、默认值和描述。默认输出到 os.Stderr. 在参数缺失或无效时调用 Usage，提供清晰提示
        示例:
```go
searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
searchTerms := searchCmd.String("terms", "", "Search terms (comma-separated)")
searchCmd.Usage = func() {
    fmt.Fprintf(searchCmd.Output(), "Usage: %s search --terms <terms>\n", os.Args[0])
    fmt.Fprintf(searchCmd.Output(), "Example: %s search --terms \"sleep,insomnia\"\n", os.Args[0])
    fmt.Fprintf(searchCmd.Output(), "Search XKCD comics by terms\n")
    searchCmd.PrintDefaults()
}
switch os.Args[1] {
case "search":
    searchCmd.Parse(os.Args[2:])
    if *searchTerms == "" {
        searchCmd.Usage()
        os.Exit(1)
    }
```
    NewFlagSet.PrintDefaults()
        方法
        作用: 将 FlagSet 中定义的所有标志的名称、默认值和使用说明打印到 flagSet.Output()（默认 os.Stderr）
    filepath.Ext
        函数
        作用: 用于提取文件路径的扩展名, 包括点. 如果没有扩展名，返回空字符串
        形参: string, 要提取扩展名的文件路径或文件名
        返回值: 文件扩展名（包括点，如 .jpg
        示例:
```go
filepath.Ext("image.jpg")  // 输出: .jpg
```
    strings.ReplaceAll()
        函数
        作用: 字符替换
        形参1: string, 要处理的字符串
        形参2: old string
        形参3: new string
        返回值: 替换后的新字符串, 原字符串不变. 如果 old 不存在于 s 中，返回 s 的副本
    os.Create()
        函数
        作用: 创建新文件, 如果存在会被覆盖
        形参: 文件路径
        返回值1: 文件句柄
        返回值2: error
    io.Copy
        函数
        作用: 拷贝数据
        形参1: dst io.Writer
        形参2: src io.Reader
        返回值1: 写了多少字节
        返回值2: error


