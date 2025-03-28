package main

import (
	"context"
	"fmt"
	"time"
)

// Суть: каждая горутина может быть отменена или прервана по таймауту.

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine is done", ctx.Err())
			return

		default:
			fmt.Println("Doing something...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go doSomething(ctx)
	time.Sleep(3 * time.Second)
	fmt.Println("Done")

}

//Вывод:
//
//Горутина автоматически завершится через 2 секунды, потому что сработает таймаут.
//
//Контекст помогает:
//
//Отменять операции.
//
//Передавать дедлайны.
//
//Прекращать цепочки вызовов по иерархии.
