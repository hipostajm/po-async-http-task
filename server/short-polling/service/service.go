package service

import (
	"github.com/google/uuid"
	"github.com/hipostajm/po-async-http-task/long-polling/model"
	"github.com/hipostajm/po-async-http-task/long-polling/repository"
)

type TaskService struct{
	repository repository.Repository;
}

func (s *TaskService) AddTask(task *model.Task){
	s.repository.AddTask(task)
}

func (s *TaskService) GetTaskByID(id uuid.UUID) (*model.Task, error){
	return s.repository.GetTaskByID(id)	
}

func NewTaskService(repository repository.Repository) TaskService{
	return TaskService{repository: repository}
}
