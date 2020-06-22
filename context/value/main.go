package main

import (
	"context"
	"fmt"
)

func main() {
	/*
		Build a parent context with multiple values
	*/
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "value: 1")
	ctx = context.WithValue(ctx, "key2", 2)
	key3 := 3
	ctx = context.WithValue(ctx, key3, make(chan int)) //any types really
	/*
		Extract values saved from the context
	*/
	fmt.Println(ctx.Value("key1"))
	fmt.Println(ctx.Value("key2"))
	fmt.Println(ctx.Value(key3))
}
