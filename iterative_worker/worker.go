package iterative_worker

import (
	"log"
	"time"
)

type Worker struct {
	ticker       *time.Ticker
	Done         chan bool
	TimeInterval time.Duration
	JobFunction  func()
}

func (worker *Worker) StartWorking() {
	// Ensure at least one execution of the job
	worker.JobFunction()

	worker.ticker = time.NewTicker(worker.TimeInterval)
	for {
		select {
		case <-worker.Done:
			return
		case t := <-worker.ticker.C:
			go func() {
				log.Println("Starting next iteration of job at ", t)
				worker.JobFunction()
				log.Println("Job iteration started at ", t, " completed successfully at ", time.Now())
			}()
		}
	}
}

func (worker *Worker) StopWorking() {
	worker.Done <- true
}
