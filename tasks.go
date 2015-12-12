package taskmanager

import (
	"sync"
)

// array of work ids
var tasks []interface{}

// mutex for delete and add to tasks
var tasks_mutex = &sync.Mutex{}

// get active tasks count
func GetTasksCount() int {
	return len(tasks)
}

// get tasks
func GetTasksIds() []interface{} {
	return tasks
}

// check task with id exists
func checkTaskExists(id interface{}) bool {
	_, ok := getTaskPosition(id)
	
	return ok
}

// delete task by id
func deleteTask(id interface{}) {
	i, ok := getTaskPosition(id)
	if ok == true {
		tasks_mutex.Lock()
		tasks = append(tasks[:i], tasks[i+1:]...)
		tasks_mutex.Unlock()
	}
}

// get task position and bool if exists
func getTaskPosition(id interface{}) (int, bool) {
	for i, v := range tasks {
		if v == id {
			return i, true
		}
	}

	return 0, false
}
