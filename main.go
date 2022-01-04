package main

import (
	"fmt"
	"unsafe"
	"time"
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


	message := make(chan string,2)
	message <- "Hello, world!"
	message <- "Hi, go "

	fmt.Println(<-message)
	fmt.Println(<-message)

	done := make(chan bool)
	go worker(done)

	<- done


	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func(){
		time.Sleep(time.Second *2)
		c2 <- "two"
	}()


	for i := 0; i < 2; i++ {
		select{
		case msg1 := <- c1:
			fmt.Println("received",msg1)
		case msg2 := <- c2:
			fmt.Println("received",msg2)
		}
	
	}

}

func worker(done chan<- bool){
	fmt.Print("working....")
	time.Sleep(3*time.Second)
	fmt.Println("done")

	done <- true
}