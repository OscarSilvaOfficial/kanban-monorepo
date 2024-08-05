package core

type Task struct {
	name, description string
	rate int
}

func (task *Task) create(name, description string, rate int) {
	task.name = name
	task.description = description
	task.rate = rate
}

func CreateTask(name, description string, rate int) *Task {
	task := new(Task)
	task.create(name, description, rate)
	return task
}
