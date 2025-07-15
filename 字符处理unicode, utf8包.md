# unicode 包
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

unicode.ReplacementChar
    常量
    = '\uFFFD' 代表 invalid 码点

# utf8 包
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

utf8.UTFMax
    常量
    = 4        代表 utf8 编码最大的字节数