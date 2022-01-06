package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
	"unsafe"
)

type Celsius float64    //摄氏度
type Fahrenheit float64 //华氏度

const (
	AbsoluteZeroC Celsius = -273.15 //绝对零度
	FreezingC     Celsius = 0       //结冰点温度
	BoilingC      Celsius = 100     //沸水温度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
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
	fmt.Printf("b 占用 %d 个字节数\n", unsafe.Sizeof(b))

	message := make(chan string, 2)
	message <- "Hello, world!"
	message <- "Hi, go "

	fmt.Println(<-message)
	fmt.Println(<-message)

	// done := make(chan bool)
	// go worker(done)

	// <- done

	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func(){
	// 	time.Sleep(time.Second)
	// 	c1 <- "one"
	// }()

	// go func(){
	// 	time.Sleep(time.Second *2)
	// 	c2 <- "two"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select{
	// 	case msg1 := <- c1:
	// 		fmt.Println("received",msg1)
	// 	case msg2 := <- c2:
	// 		fmt.Println("received",msg2)
	// 	}

	// }

	//切片是引用类型
	nums := []int{1, 2, 3, 4}
	fmt.Printf("\"修改前\": %v\n", "修改前")
	fmt.Println(nums)

	nums2 := nums
	nums2[0] = 123
	fmt.Println(nums)
	fmt.Println(nums2)

	fmt.Println("我是补全的代码")

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f == 0)

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 123

	fmt.Println(b1)
	fmt.Println(b2)

	var s1 []int
	fmt.Println(s1 == nil)

	s1 = []int{1, 2, 3}
	fmt.Println("cap:", cap(s1))
	fmt.Println("len:", len(s1))

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v",err)
	}
	log.Printf("working directory = %s",cwd)

}

func worker(done chan<- bool) {
	fmt.Print("working....")
	time.Sleep(2 * time.Second)
	fmt.Println("done")

	done <- true
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main1() {
	// http://localhost:7070/?cycles=10
	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.Atoi(r.FormValue("cycles"))
		if err != nil {
			cycles = 2
		}
		lissajous(w, cycles)
	}
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":7070", nil)
	fmt.Println(err)
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
