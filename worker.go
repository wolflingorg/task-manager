package taskmanager

// Handler
// You must create function like this in your program
// Your task will be transfered to this function
type WorkHandler func(work WorkRequest, worker_id int)

// Worker structure
type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
	Handler     WorkHandler
}

// Start worker
// registers itself in the queue and waits for a new task
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				w.Handler(work, w.ID)
				deleteTask(work.Id)

			case <-w.QuitChan:
				return
			}
		}
	}()
}

// Stop worker
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// Create new worker
func NewWorker(id int, workerQueue chan chan WorkRequest, Handler WorkHandler) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
		Handler:     Handler}

	return worker
}
