package model

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
)

type EmailStatus string

const (
	Pending EmailStatus = "Pending"
	Done EmailStatus = "Done"
)

type Task struct{
	TaskID uuid.UUID
	Count uint
	Status EmailStatus
	Data string
	Email string
}

func NewTask(count uint, email string) *Task{
	return &Task{TaskID: uuid.New(), Count: count, Status: Pending, Data: fmt.Sprintf("%x", sha256.Sum256([]byte(uuid.NewString()))), Email: email}
}

