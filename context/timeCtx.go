package context

import (
	"context"
	"fmt"
	"time"
)

func TimeContext(){
	ctx,cancel := context.WithTimeout(context.Background(), 3 * time.Second)

	defer cancel()
	go func(ctx context.Context) {
		for range time.Tick(time.Second){
			select {
			case <-time.After(4 * time.Second):
				fmt.Println("Slept over")
			case <-ctx.Done():
				fmt.Println(ctx.Err())
			}
		}
	}(ctx)

	time.Sleep(4 * time.Second)
}
