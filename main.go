package main

import (
	"fmt"
	"unsafe"
)

func print(s []string) {
	for _, v := range s {
	 fmt.Print(v)
	}
}
func main() {
	print([]string{"你好, ", "脑子进了煎鱼"})

	fmt.Println()
	var a byte = 'A'
	var b rune = 'B'
	fmt.Printf("a 占用 %d 个字节数\n", unsafe.Sizeof(a))
	fmt.Printf("b 占用 %d 个字节数\n",unsafe.Sizeof(b))
}
   