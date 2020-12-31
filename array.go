package main

import (
	"fmt"
)

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3}
var arr2 = [...]int{1, 2, 3}
var arr3 = [5]string{3: "hello", 4: "world"}

func main() {
	fmt.Print("hello world")
	fmt.Println(arr0)
	fmt.Println(arr1)
	// arr2.push(33)
	fmt.Println(arr2)
	fmt.Println(arr3)

	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}

	fmt.Println(d)

	//多维数组

}
