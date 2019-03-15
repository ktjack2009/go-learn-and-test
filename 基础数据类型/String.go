package main

import (
	"bytes"
	"fmt"
)

func main() {
	// StringDemo1()
	// StringDemo2()
	// StringDemo3()
	StringDemo4(1, 2, 3, 4)
}

func StringDemo4(values ...int) {
	var buf bytes.Buffer // 用于字节silce的缓存
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	a := "hello世界"
	buf.WriteByte('[')
	for i, v := range []rune(a) {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(v)
	}
	buf.WriteByte(']')
	s := buf.String()
	fmt.Println(s)
}

func StringDemo3() {
	var a interface{}
	a = "hello 世界"
	b := []rune(a.(string))
	c := string(b[:10])
	fmt.Println(c)

}

func StringDemo2() {
	// Unicode码点对应rune整数类型【rune是int32等价类型】
	s := "hello世界"
	fmt.Println(len(s)) // 11
	a := []rune(s)
	fmt.Println(len(a)) // 7
	for v := range a {
		// fmt.Printf("%c\n", a[v])
		fmt.Println(string(a[v]))
	}
}

func StringDemo1() {
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7]) // 104 111【字节值】
	fmt.Println(s[0:1])     // h
	fmt.Println(s[:5])      // 返回字符串
}
