package main

type Flags uint

// 枚举, 设置标志位
const (
	// 1 左移0位, 1*2^0=1     0000 0001
	FlagUp Flags = 1 << iota
	// 1 左移1位, 1*2^1=2     0000 0010
	FlagBroadcast
	// 1 左移2位, 1*2^2=4     0000 0100
	FlagLoopback
	// 1 左移3位, 1*2^3=8     0000 1000
	FlagPointToPoint
	// 1 左移4位, 1*2^4=16    0001 0000
	FlagMulticast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }
