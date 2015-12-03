package taskmanager

// work strucure
// all tasks need to write to Data
type WorkRequest struct {
	Id   int
	Data interface{}
}
