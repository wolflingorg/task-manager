# Task Manager Golang
Task Manager - позволяет распараллеливать выполнение задач.

# Установка
`go install github.com/wolflingorg/task-manager`

# Использование
- создаем функцию Handler, по образу `type WorkHandler func(work WorkRequest, worker_id int)`
- вызываем `StartDispatcher(nworkers int, handler WorkHandler)` в который передаем количество воркеров и созданную Вами функцию обработчик
- добавляем задачи, вызывая функцию `NewWork(work WorkRequest)`

# Пример использования
```go
package main

import (
	"fmt"
	tm "task-manager"
	"time"
)

// функция обработчик
func Handler(work tm.WorkRequest, worker_id int) {
	// тут пишем, что нужно сделать с задачей
	// саму задачу берем из поле Data
	fmt.Println(work.Data)
}

func main() {
	// запускаем воркеров
	tm.StartDispatcher(10, Handler)

	task_id := 0

	for {
		select {
		case <-time.After(time.Second):
			for i := task_id; i < task_id+10; i++ {
				// создаем задание
				// у каждого задания должен быть уникальный id
				// данные задания могут быть любыми (число, строка, структура)
				work := tm.WorkRequest{Id: i, Data: i}
				// запускаем задание в работу
				tm.NewWork(work)
			}
		}
	}
}
```

Реальный пример можно посмотреть на проекте https://github.com/wolflingorg/rss-grabber
Здесь Task Manager используется для "обновления RSS каналов"
