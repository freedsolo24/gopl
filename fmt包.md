# fmt 包
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

fmt.Errorf()
    函数
    作用: 将原始错误包装到新的错误中。根据指定的格式化字符串 format 和参数 a 创建一个新的 error 类型值. 返回值是一个实现了 error 接口的类型
    形参1: 格式化的字符串
    形参2: 传递要格式化的值，对应格式化字符串中的占位符
    返回值: 一个 error 接口类型的值，包含格式化后的错误消息
    示例:
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