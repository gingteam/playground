package main

import (
	"C"
	"fmt"
	"time"
)

// we needed to have empty main in package main :)
// because -buildmode=c-shared requires exactly one main package
func main() {

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

//export hello
func hello() {
	fmt.Println("Hello from Go")
}

//export run
func run() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println("result:", <-results)
	}
}
