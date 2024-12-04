package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID       int
	Work     func() (int, error)
	ResultCh chan<- Result
}

type Result struct {
	JobID int
	Value int
	Err   error
}

func workerErrH(id int, jobs <-chan Job) {
	for j := range jobs {
		fmt.Printf("worker %d started job %d\n", id, j.ID)
		value, err := j.Work()
		j.ResultCh <- Result{JobID: j.ID, Value: value, Err: err}
		fmt.Printf("worker %d finished job %d\n", id, j.ID)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	for w := 1; w <= 3; w++ {
		go workerErrH(w, jobs)
	}

	for j := 1; j <= numJobs; j++ {
		workFunc := func() (int, error) {
			time.Sleep(time.Second)
			if j%2 == 0 {
				return 0, fmt.Errorf("error in job %d", j)
			}
			return j * 2, nil
		}
		jobs <- Job{ID: j, Work: workFunc, ResultCh: results}
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		result := <-results
		if result.Err != nil {
			fmt.Printf("Job %d failed: %v\n", result.JobID, result.Err)
		} else {
			fmt.Printf("Job %d result: %d\n", result.JobID, result.Value)
		}
	}
}
