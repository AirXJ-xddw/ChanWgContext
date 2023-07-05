package dir

import "fmt"

func processC(ch chan int, rang int){
	x, y := 0, 1
	for i:=0; i < rang; i++{
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func Channel(){
	channels := make(chan int , 10)
	go processC(channels, cap(channels))

	for i := range channels{
		fmt.Println(i)
	}
}