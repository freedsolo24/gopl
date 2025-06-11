# ch1
* 1.2 示例
    os.Args[1:] 
        返回值类型[]string
    strings.Join()
        作用: 把切片里面的N个string元素, 用分隔符连接
        形参1: []string
        形参2: string类型, 分隔符
        返回值: string
* 1.3 示例
    os.Stdin
        类型是os.*File结构类. 实现了Read方法, 可以赋值给Reader接口类
        也是一个打开的文件句柄, fd=0
    bufio.Newscanner() 
        形参1类型io.Reader接口类 
        返回值Scanner结构类 得到扫描器句柄
    scanner.Scan() 
        将系统缓冲区读到自己的临时缓冲区
        返回值: bool
    scanner.Text() 
        取出临时缓冲区的内容
        返回值: string
    os.Open()      
        形参是字符串，即文件路径
        返回值1: *os.File结构类,代表一个打开的文件句柄或者说对文件的抽象. 相当于返回一个文件描述符,表示打开一个文件句柄, 通过句柄操作文件.
        返回值2: err
    os.ReadFile()
        作用: 一次性把整个文件读取到大内存缓冲区
        形参是字符串string
        返回值1 是[]byte
        返回值2 err
    strings.Split()
        作用: 分割, 把一个string, 按照分割符分割
        形参1: string, 要分割的string
        形参2: string, 分隔符
        返回值: []string
    strings.Split和strings.Join互为逆操作
    