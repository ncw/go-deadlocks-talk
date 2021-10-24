package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Job struct {
	// stuff
}

func NewJob() *Job {
	fmt.Println("New Job")
	return &Job{}
}

// Does something (eg list a directory)
// May produce more Jobs (eg more directories to list)
func (j *Job) Process() []*Job {
	var newJobs []*Job
	for i := rand.Intn(3); i > 0; i-- {
		newJobs = append(newJobs, NewJob())
	}
	fmt.Println("Done Job")
	return newJobs
}

// Process all the jobs from the channel keeping track of the number
// of active jobs in the activeJobs WaitGroup
func ProcessJobs(activeJobs *sync.WaitGroup, jobs chan *Job) {
	for job := range jobs {
		newJobs := job.Process()
		for _, newJob := range newJobs {
			activeJobs.Add(1)
			jobs <- newJob
		}
		activeJobs.Done()
	}
}

func main() {
	var activeJobs sync.WaitGroup
	jobs := make(chan *Job, 1)
	activeJobs.Add(1)
	jobs <- NewJob()

	go ProcessJobs(&activeJobs, jobs)
	go ProcessJobs(&activeJobs, jobs)
	go ProcessJobs(&activeJobs, jobs)
	activeJobs.Wait()
	close(jobs)
}
