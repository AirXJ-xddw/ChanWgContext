package context

import (
	"context"
	"fmt"
)

type Ctxkey string

func ValueContext(){
	ctx := context.WithValue(context.Background(),Ctxkey("key"),"val")

	get := func(ctx context.Context, k Ctxkey){
		if v,ok := ctx.Value(k).(string); ok{
			fmt.Println(v)
		}else{
			fmt.Println("not found")
		}
	}

	get(ctx,"key")
	get(ctx,"kye")
}