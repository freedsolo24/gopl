# flag 包
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