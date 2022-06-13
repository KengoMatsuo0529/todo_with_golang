// 型について
package main

import "fmt"

func main() {
	// 整数の型
	var i int = 100
	fmt.Println(i)

	// 型を確認する方法
	fmt.Printf("%T\n", i)

	// 型変換
	fmt.Printf("%T\n", int32(i))

	// 浮動小数点型
	var flt64 float64 = 2.4
	fmt.Println(flt64)

	flt := 3.2
	fmt.Println(flt)

	fmt.Println(flt64 + flt)
	fmt.Printf("%T, %T\n", flt64, flt)

	// bool
	var f, t bool = false, true
	fmt.Println(f, t)

	// 文字列型
	var s string = "Hello Golang"
	fmt.Println(s)

	fmt.Println(`test
	test
		test`)
	fmt.Println(`"`)

	// Hello GolangのHを抜き出したい => 72
	fmt.Println(s[0])
	// バイト型を文字列に変換が必要 => H
	fmt.Println(string(s[0]))
}