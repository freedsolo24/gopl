# time 包
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