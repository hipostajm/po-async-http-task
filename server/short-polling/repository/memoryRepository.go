package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hipostajm/po-async-http-task/long-polling/model"
)

type MemoryRepostiory struct{
	tasks map[uuid.UUID]*model.Task
}

func (r* MemoryRepostiory) AddTask(datum *model.Task){
	r.tasks[datum.TaskID] = datum
	go r.waitToSetDone(datum.TaskID)
}

func (r* MemoryRepostiory) GetTaskByID(id uuid.UUID) (*model.Task, error){
	task, ok := r.tasks[id]
	if (!ok){
		return nil, errors.New("ID not found")
	}
	return task,nil 
}

func (r* MemoryRepostiory) waitToSetDone(id uuid.UUID){
	task := r.tasks[id]
	time.Sleep(time.Duration(task.Count)*time.Second)
	task.Status = model.Done
}

func NewMemoryRepository() *MemoryRepostiory{
	return &MemoryRepostiory{tasks: make(map[uuid.UUID]*model.Task)}
}
