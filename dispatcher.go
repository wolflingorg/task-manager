package taskmanager

// Queue of free workers
var WorkerQueue chan chan WorkRequest

// Create workers and delegate work for them
// TODO add check the number of working processes
func StartDispatcher(nworkers int, WorkQueue chan WorkRequest, Do DoHandler) {
	WorkerQueue = make(chan chan WorkRequest, nworkers)

	for i := 0; i < nworkers; i++ {
		worker := NewWorker(i+1, WorkerQueue, Do)
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
