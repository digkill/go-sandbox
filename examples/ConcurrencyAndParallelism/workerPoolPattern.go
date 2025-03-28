package main

import (
	"fmt"
	"time"
)

// Worker Pool Pattern — создание пула горутин для управления задачами.
// Суть: ограничиваем количество одновременно работающих горутин (workers), которые берут задачи из общего канала.

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // имитация работы
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// запускаем пул воркеров
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// отправляем задания
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// собираем результаты
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result:", <-results)
	}
}
