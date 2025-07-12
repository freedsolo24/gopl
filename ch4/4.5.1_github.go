// http 请求: [repo:golang/go is:open json decoder]
// 需要拼接成: https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder
// api.github.com/search/issues 入口端点地址
// ?q=后面是查询的参数
// repo:USERNAME/REPOSITORY 映射的是 repo:golang/go
// is:open 是否公开
// json decoder 查询的关键字
// sort=created&order=asc 根据created date升序排列, 这个不能写在q查询参数中
// https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder&sort=created&order=asc

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue // 核心点是(1)内存效率, (2)共享底层数据, (3)空值处理
}

// (1) Items是一个切片, 切片里每一个元素指向Issue结构体的指针. Items切片只存储指向这些实例的指针, 当这个结构体被操作时, 只传递指针节省内存.
// (2) 指针允许多个地方引用同一个Issue实例, 如果多个地方修改 Issue 字段, 所有引用该指针的地方都会反应这个变化
// (3) []*Issue 允许 Issue 为 nil, 表示元素为初始化或无效. []Issue 每个元素都会初始化为 Issue 的零值, 无法表示"缺失状态"
// Items []Issue. Items是一个切片, 切片里每一个元素指向Issue结构体的副本.

// type IssuesSearchResult struct {
// 	TotalCount int `json:"total_count"`
// 	Issue      Issue    // 可以这么写, 但不符合API响应格式, 如果这么写, 只支持一个Issue
// }

// "items": [
//     {
//       "number": 123,
//       "title": "bug fix",
//       ...
//     },
//     {
//       "number": 124,
//       "title": "feature request",
//       ...
//     }
//   ]
// key是items, value是数组[ ], 里面有多个issue, 所以要匹配这个结构, 就得用Items: []*Issue

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User     // 不写成User User, 核心点是内存效率、共享语义和空值处理
	CreateAt time.Time `json:"created_at"`
	Body     string
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

const IssuesURL = "https://api.github.com/search/issues"

// 我输入的"go run . repo:golang/go is:open json decoder
// terms []string, 拿到的是 [repo:golang/go is:open json decoder]
func SearchIssues1(terms []string) (*IssuesSearchResult, error) {
	// 作用是: 参数 terms 拼接成一个带空格的搜索字符串, 再把它 url 编码, 生成可用于查询的 q=... 参数
	// QueryEscape把[repo:golang/go is:open json decoder], 转义成: "repo%3Agolang%2Fgo+is%3Aopen+json+decoder"
	q := url.QueryEscape(strings.Join(terms, " "))

	// 最后拼接成: https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder
	// 对这个地址发起http请求, 得到响应
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// 检查 HTTP 请求是否失败
	if resp.StatusCode != http.StatusOK {
		// 关闭 HTTP 响应正文的底层数据流（通常是网络连接或缓冲区）. 这会释放与该流相关的资源，例如网络 socket 或内存缓冲区
		// 如果不调用 Close，可能会导致资源泄漏（如文件描述符未释放），尤其是在高并发场景下，可能耗尽系统资源
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	// 声明且初始化一个 IssuesSearchResult 结构体的实例 result
	var result IssuesSearchResult
	// 初始化创建 json 解码器, 对 resp.Body 缓冲区中的数据流做解码, 解码后填充到到 &result 结构体实例中, 根据 json 标签做对应
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// 练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年
// 思路:
// (1) 遍历 IssuesSearchResult.Items 切片里的每一个元素
// (2) 拿今天的时间, 计算切片里的每个元素的 created_at 和现在的时间间隔
// (3) <=30 天, 属于一个月内; >30天&&<=1年, 属于一年内; >1年, 属于一年以上
// (4) 创建3个slice, 保存.
func SearchIssues2(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	url := fmt.Sprintf("%s?q=%s", IssuesURL, q)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("search query failed: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", err)
	}
	var result IssuesSearchResult
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed:%s", err)
	}
	resp.Body.Close()
	return &result, nil
}

func categories(result *IssuesSearchResult) {
	now := time.Now()
	var (
		monthIssues []*Issue
		yearIssues  []*Issue
		oldIssues   []*Issue
	)

	for _, item := range (*result).Items {
		duration := now.Sub((*item).CreateAt)

		switch {
		case duration < 30*24*time.Hour:
			monthIssues = append(monthIssues, item)
		case duration < 365*24*time.Hour:
			yearIssues = append(yearIssues, item)
		default:
			oldIssues = append(oldIssues, item)
		}
	}
	fmt.Println("🔹 Issues less than 1 month old:")
	for _, issue := range monthIssues {
		fmt.Printf("#%-5d, %9s, %.55s\n", (*issue).Number, (*issue).User.Login, (*issue).Title)
	}
	fmt.Println("🔹 Issues less than 1 year old:")
	for _, issue := range yearIssues {
		fmt.Printf("#%-5d, %9s, %.55s\n", (*issue).Number, (*issue).User.Login, (*issue).Title)
	}
	fmt.Println("🔹 Issues more than 1 year old:")
	for _, issue := range oldIssues {
		fmt.Printf("#%-5d, %9s, %.55s\n", (*issue).Number, (*issue).User.Login, (*issue).Title)
	}
}
