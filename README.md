# Task Manager Golang
Task Manager - позволяет распараллеливать выполнение задач.

# Установка
`go install github.com/wolflingorg/task-manager`

# Использование
- создаем функцию Handler, по образу `type WorkHandler func(work WorkRequest, worker_id int)`
- вызываем `StartDispatcher(nworkers int, handler WorkHandler)` в который передаем количество воркеров и созданную Вами функцию обработчик
- добавляем задачи, вызывая функцию `NewWork(work WorkRequest)`

# Пример использования
