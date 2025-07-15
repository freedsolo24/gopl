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

func init() { ... }
    init函数
    作用: 包中的init函数在包导入的时候, 最先执行. 在此函数中用来初始化一个数据表的变量pc[265]byte, 让pc这个数组变量执行后里面有值
    init函数不能被调用和引用