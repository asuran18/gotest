package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
 * 字符串示例
 */
func main() {
	fmt.Println("字符串长度获取")
	//len 获取字节长度或字符个数
	tip1 := "test"
	//4个字符
	fmt.Println(len(tip1))
	tip2 := "忍者"
	//6个字节 字符串以 UTF-8 格式存储，中文单个字符占3个字节
	fmt.Println(len(tip2))
	//utf8.RuneCountInString 获取Unicode 字符数量
	fmt.Println(utf8.RuneCountInString("忍者"))        //2
	fmt.Println(utf8.RuneCountInString("忍者,fight!")) //9

	fmt.Println("字符串遍历")
	for i := 0; i < len(tip1); i++ {
		fmt.Printf("ascii: %c \n", tip1[i])
	}
	//带汉字输出
	for _, s := range tip2 {
		fmt.Printf("Unicode: %c\n", s)
	}

	fmt.Println("字符串截取")
	comma := strings.Index(tip1, "e")
	// =>est
	fmt.Printf(tip1[comma:])

	fmt.Println("字符串修改")
	strBytes := []byte(tip1)
	for i := 0; i < len(tip1); i++ {
		strBytes[i] = tip1[i] + 1
	}
	fmt.Printf(string(strBytes))

	fmt.Println("字符串拼接，非+方式")
	var sb bytes.Buffer
	sb.WriteString("add")
	sb.WriteString(" string")
	fmt.Println(sb.String())
}
