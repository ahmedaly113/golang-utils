package worker

import "github.com/ahmedaly113/golang-utils/sync"

// Request Base type of all work objects
type Request interface{}

// Worker for now
type Worker struct {
	Queue  chan Request
	Errs   chan error
	Op     func(Request) error
	Marker *sync.CountWG
}

// Start a worker
func (w *Worker) Start() {
	go w.run()
}

func (w *Worker) run() {
	for work := range w.Queue {
		if err := w.Op(work); err != nil {
			w.Errs <- err
		}
		w.Marker.Done()
	}
}
