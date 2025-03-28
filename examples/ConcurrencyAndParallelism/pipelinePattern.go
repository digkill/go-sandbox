package main

import (
	"fmt"
)

// Суть: данные обрабатываются поэтапно, каждый этап в своей горутине.

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	ch := gen(1, 2, 3, 4, 5, 6)
	out := square(ch)

	for result := range out {
		fmt.Println(result)
	}
}

// Цепочка: gen -> square -> output.
//
//Преимущества:
//
//Отделение логики по стадиям.
//
//Гибкость и масштабируемость.
