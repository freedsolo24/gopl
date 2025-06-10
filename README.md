ch1
1.2 os.Args[1:] 返回值类型[]string
1.3 bufio.Newscanner() 形参1类型os.Stdin 返回值类型Scanner结构类 得到扫描器句柄
    scanner.Scan() 将系统缓冲区读到自己的临时缓冲区
    scanner.Text() 取出临时缓冲区的内容
    os.Open()      形参1是字符串，即文件路径 返回值1