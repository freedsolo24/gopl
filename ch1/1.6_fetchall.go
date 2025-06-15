package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// 对输入的url, 并发获取多个url的body
// 观察:main函数运行的总时长, 不是所有协程运行时长的和, 而是不超过耗时最长时间的任务获取
// 说明了任务的并发
func fetchall() {

	// 得到当前的时间点
	startTime := time.Now()
	// 之前make一个map, make(map[string]key), 没有哈希桶
	// 现在make一个channel, 是一个同步通道, 里面元素的类型是string
	// 同步通道, 手递手的同步过程
	// 这是一个双向通道, chan是类型, 里面的元素是string, 类似[]string
	ch := make(chan string)
	// 把所有的参数放到[]string容器里面
	urls := os.Args[1:]

	// channel是引用类型, string也是引用类型, 所以传递的都是标头值
	// string的标头值[指针,len], channel本质是指针结构体,指向通道底层缓冲区
	// 把遍历的url和make的channel传给子协程, 一个url开启一个子协程
	// 如果urls切片里有3个元素, 就会创建3个协程, 每个协程都往ch通道里写数据, 一共写3次
	for _, url := range urls {
		go fetch(url, ch)
	}

	// 主协程在读取通道, 这是一个同步通道, 是一个"手递手的同步方式"
	// 遍历的次数是len(urls), 如果urls切片有3个元素, 就会遍历3次, 一共读通道3次
	for range urls {
		// 从通道里面读出来内容, 打印到Stdin
		fmt.Println(<-ch)
	}

	// 计算从startTime到现在经历了多久, 换算成秒
	stopTime := time.Since(startTime).Seconds()

	// 这个计算的是整体的时间
	fmt.Printf("%.2f elapsed\n", stopTime)
}

// 子协程运行这个函数, 使用同一个通道底层缓冲区
// ch chan<-string   ch是标识符  chan是通道类型  <- 通道的方向:只读  string 通道里的元素是string
func fetch(url string, ch chan<- string) {
	goroutineStartTime := time.Now()
	resp, err := http.Get(url)
	// 对Get()错误处理, 打印错误后, 退出程序
	if err != nil {
		// Sprint简单拼接形参, 成为字符串
		ch <- fmt.Sprint("Error:", err)

		// 函数正常结束消亡, 会执行defer
		// 如果向以前协程os.Exit(1), 整个程序立即终止, 所有goroutine也会被强制中止, 不会执行defer
		return
	}

	// 复制body到黑洞, 返回成功复制了多少个字节
	// 作用1: io.Discard你写给它的数据都会被默默“吃掉”，什么都不做, 把响应体读完但不保存内容，只为触发数据接收或统计字节数
	// 作用2: HTTP/1.1 默认支持 连接复用（keep-alive），前提是：必须把 resp.Body 的内容读完, 如果你不读完就 resp.Body.Close()，连接就不能复用，性能下降
	nbytes, err := io.Copy(io.Discard, resp.Body)

	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v\n", url, err)
		return
	}

	// 计算每个协程消耗的时间
	secs := time.Since(goroutineStartTime).Seconds()
	// %.2f: 小数位的精度是2
	// %7d:  右对齐, 打印宽度是7, 不足7位左边补空格
	// 往通道里写字符串, 用Sprintf格式化输出一个字符串
	ch <- fmt.Sprintf("%.2f|%7d|%s\n", secs, nbytes, url)

}
