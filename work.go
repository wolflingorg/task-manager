package taskmanager

// work strucure
// all tasks need to write to Data
type WorkRequest struct {
	Id   interface{}
	Data interface{}
}

// append new work
func NewWork(work WorkRequest) bool {
	if checkTaskExists(work.Id) == false {
		tasks_mutex.Lock()
		tasks = append(tasks, work.Id)
		tasks_mutex.Unlock()
		
		WorkQueue <- work

		return true
	}

	return false
}
