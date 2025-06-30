/*
函数的作用:

	    拿到go run 之后跟的参数, 把参数解析成float64类型的数值
		把数值转换成华氏度, 摄氏度
		打印华氏度, 并且打印这个华氏度 等价于 多少的摄氏度
		打印摄氏度, 并且打印这个摄氏度 等价于 多少的华氏度
		相同的摄氏度,等价于,多少的华氏度
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	// 说明命令后面没有参数
	if len(os.Args[1:]) == 0 {
		var stds []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stds = append(stds, scanner.Text())
		}
		conv(stds)
	} else {
		// 遍历参数, 拿到一个字符串类型的参数, 把参数转换成float64
		conv(args)
	}
}

func conv(ss []string) {
	for _, s := range ss {
		fl, err := strconv.ParseFloat(s, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
		}
		c := Celsius(fl)
		f := Fahrenheit(fl)
		m := Meter(fl)
		mi := Mile(fl)
		p := Pound(fl)
		k := Kilogram(fl)

		// 我为自定义的类型: Celsius, Fahrenheit 自定义了String方法
		// 两个自定义的类型: 自动实现了Stringer接口
		// 我用%s打印f和c, 类型是Celsius, Fahrenheit, %s会自动调用我写的String()方法
		// fmt.Printf("%s",f) 等价于 fmt.Printf("%s",f.String)
		// Stringer接口声明要String()方法的返回值是string, 所以main函数中, fmt.Printf打印自定义类型一定要用%s或%v占位
		fmt.Printf("input %s\n", s)
		ctof := fmt.Sprintf("摄氏转华氏:")
		ftoc := fmt.Sprintf("华氏转摄氏:")
		mtomi := fmt.Sprintf("米转英里:")
		mitom := fmt.Sprintf("英里转米:")
		ptokg := fmt.Sprintf("磅转千克:")
		kgtop := fmt.Sprintf("千克转磅:")
		fmt.Printf("%-6s %s=%s,%6s %s=%s\n", ctof, c, CtoF(c),
			ftoc, f, FtoC(f))
		fmt.Printf("%-7s %s=%s,%7s %s=%s\n", mtomi, m, MtoMI(m),
			mitom, mi, MItoM(mi))
		fmt.Printf("%-7s %s=%s,%6s %s=%s\n", ptokg, p, PoundtoKilogram(p),
			kgtop, k, KilogramtoPound(k))
		fmt.Println("================")
	}
	a := uint64(0x0f0f0f0f0f0f0f0f)
	fmt.Println(PopCount(a))

	b := uint64(0x0f0f0f0f0f0f0f0f)
	fmt.Println(PopCountLoop(b))

}
