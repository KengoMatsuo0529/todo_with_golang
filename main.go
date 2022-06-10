// パッケージの宣言
package main

// パッケージのインポート
import (
	"fmt"
	"time"
)

// 関数外で暗黙的な定義をしても関数内で呼び出せない
// i5 := 500
// 明示的だと関数外でも定義可能
var i5 = 500

// 別の関数を定義
func outer() {
	// main関数内でouterを呼び出すが、その場合は明示的でも暗黙的でも定義してOK
	// var s5 string = "outer"
	s5 := "outer"
	fmt.Println(s5)
}

// メイン関数
func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())

	// 明示的な変数の定義
	var i int = 100
	fmt.Println(i)

	var s string = "Hello GO"
	fmt.Println(s)

	var t,f bool = true, false
	fmt.Println(t,f)

	var (
		i2 int = 200
		s2 string = "Hello GO2"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)
	// 数値を定義しないと、型の初期値が入る

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 暗黙的な宣言（varがいらず:=で型推論してくれる）
	i4 := 400
	fmt.Println(i4)

	// 値の更新は可能
	i4 = 450
	fmt.Println(i4)

	// i4 := 500
	// 値の再定義はできない

	// i4 = "Hello"
	// 46行目でintを入れているので、stringに更新はできない

	fmt.Println(i5)

	outer()
}