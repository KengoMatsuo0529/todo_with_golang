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

	// バイト型
	c := []byte{72, 73}
	fmt.Println(c)

	// 配列
	// 配列の型のみを宣言。この場合配列の値には初期値が入る
	var arry1 [3]int
	fmt.Println(arry1)

	// 配列の値を指定して宣言。intの初期値は0でstringの初期値は""(空文字)
	var arry2 [3]string = [3]string{"a", "b"}
	fmt.Println(arry2)

	arry3 := [2]string{"tes", "t"}
	fmt.Println(arry3)

	// [...]とすると個数を指定しなくても値の中身を分析して個数推測してくれる
	arry4 := [...]int{1, 2}
	fmt.Println(arry4)

	// 配列のインデックスを指定し値を代入
	arry2[2] = "c"
	fmt.Println(arry2)

	// 配列の長さを確認
	fmt.Println(len(arry3))

	// インターフェイス型
	var x interface {}
	// 初期値はnil
	fmt.Println(x)

	// interfaceはintやstring、配列など様々な型を代入できる
	x = 1
	fmt.Println(x)

	x = 2.34
	fmt.Println(x)

	x = "test"
	fmt.Println(x)

	x = [2]int{1, 2}
	fmt.Println(x)
}