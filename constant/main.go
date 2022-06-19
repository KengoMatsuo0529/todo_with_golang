
package main

import "fmt"

// 定数

const Pi = 3.14

const (
	Url = "http://"
	Sitename = "test"
)

func main() {
	fmt.Println(Pi)

	// 定数の上書きはできない
	// const Pi = 3

	fmt.Println(Url, Sitename)
}
