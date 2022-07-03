
package main

import "fmt"

// 定数

const Pi = 3.14

const (
	Url = "http://"
	Sitename = "test"
)

func Plus(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(Pi)

	// 定数の上書きはできない
	// const Pi = 3

	fmt.Println(Url, Sitename)

	// 比較演算子
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)

	// 関数
	fmt.Println(Plus(2, 3))
}
