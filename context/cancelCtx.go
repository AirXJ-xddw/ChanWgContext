package context

import (
	"context"
	"fmt"
	"time"
)

func CancelContext(){
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	go func(ctx context.Context) {
		for range time.Tick(time.Second){
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				fmt.Println("monitor working")
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
}
