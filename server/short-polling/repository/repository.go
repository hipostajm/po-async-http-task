package repository

import (
	"github.com/google/uuid"
	"github.com/hipostajm/po-async-http-task/long-polling/model"
)

type Repository interface{
	AddTask(datum *model.Task)
	GetTaskByID(id uuid.UUID) (*model.Task, error)
	waitToSetDone(id uuid.UUID)
} 
