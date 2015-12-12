package taskmanager

// Queue of free workers
var (
	WorkerQueue chan chan WorkRequest
	WorkQueue   chan WorkRequest
)

// Create workers and delegate work for them
// TODO add check the number of working processes
func StartDispatcher(nworkers int, handler WorkHandler) {
	WorkerQueue = make(chan chan WorkRequest, nworkers)
	WorkQueue = make(chan WorkRequest, nworkers)

	for i := 0; i < nworkers; i++ {
		worker := NewWorker(i+1, WorkerQueue, handler)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				go func() {
					worker := <-WorkerQueue
					worker <- work
				}()
			}
		}
	}()
}
