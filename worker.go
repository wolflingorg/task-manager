package taskmanager

// Handler
// You must create function like this in your program
// Your task will be transfered to this function
type DoHandler func(work WorkRequest)

// Worker structure
type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
	Do          DoHandler
}

// Start worker
// registers itself in the queue and waits for a new task
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				w.Do(work)

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
func NewWorker(id int, workerQueue chan chan WorkRequest, Do DoHandler) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
		Do:          Do}

	return worker
}
