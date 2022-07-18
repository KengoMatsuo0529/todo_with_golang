package main

import (
	"fmt"
	"time"
)

// init（初期化処理）

func init() {
	fmt.Println("init")
}

// goroutin（並行処理）

func sub() {
	for {
		fmt.Println("sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}


func main() {
	go sub()

	for {
		fmt.Println("main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}