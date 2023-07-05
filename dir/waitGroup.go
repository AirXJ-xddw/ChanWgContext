package dir

import (
	"fmt"
	"sync"
)

func processW(wg *sync.WaitGroup, x int){
	fmt.Printf("processW %d is doing\n", x)
	wg.Done()
}

func WaitGroup()  {
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++ {
		go processW(&wg, i)
	}
	wg.Wait()
}
