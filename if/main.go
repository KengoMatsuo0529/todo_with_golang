package main

import (
	"fmt"
	"strconv"
)

// エラー文

// 関数が終了したタイミングで実行する処理を登録できる
// 複数のdeferを登録すると、後に登録されたものから実行される
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func main() {
	var s string = "100"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

	// for文

	a := 0
	for {
		a++
		if a == 3 {
			break
		}
		fmt.Println("Hello")
	}

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for b := 0; b < 10; b++ {
		fmt.Println(b)
	}

	array := [3]int {1,2,3}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// switch文

	// var x := interface{} = 3

	// v := x.(int)
	// switch v.(type) {
	// case int:
	// 	fmt.Println("int")
	// case string:
	// 	fmt.Println("string")
	// case float32:
	// 	fmt.Println("default")
	// }

	// defer
	TestDefer()
}


