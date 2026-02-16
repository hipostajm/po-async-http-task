package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hipostajm/po-async-http-task/long-polling/model"
	"github.com/hipostajm/po-async-http-task/long-polling/service"
	"io"
)

type TaskHandler struct{
	service service.TaskService	
}

func getIDFromPathParam(r *http.Request) (*uuid.UUID, error){
	id := r.PathValue("id")
	parsedID, err := uuid.Parse(id)
	
	if (err != nil){
		return nil, errors.New("decode error")
	}
	return &parsedID, nil 
}

func writeBadRequest(w http.ResponseWriter, message string){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.NewOutput(message, false))
}

func logBody(r *http.Request){
	body, err := io.ReadAll(r.Body)
	if err != nil {
			return
	}
	defer r.Body.Close()

	bodyString := string(body)
	log.Println(bodyString)
}

func (h *TaskHandler) HandleTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	switch r.Method{
		case http.MethodPost:
			h.handleTaskPost(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) handleTaskPost(w http.ResponseWriter, r *http.Request){
	var taskInput model.TaskInput

	err := json.NewDecoder(r.Body).Decode(&taskInput)
	if (err != nil){
		log.Println(err)
		writeBadRequest(w, err.Error())
		return
	}
	
	task := model.NewTask(taskInput.Count, taskInput.Email)
	h.service.AddTask(task)

	w.Header().Set("Location", "/task/"+task.TaskID.String())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.NewOutput("", true))
}

func (h *TaskHandler) HandleTaskID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	switch r.Method{
		case http.MethodGet:
			h.handleTaskIDGet(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h  *TaskHandler) handleTaskIDGet(w http.ResponseWriter, r *http.Request){
	id, err := getIDFromPathParam(r)

	if (err != nil){
		log.Println(err)
		writeBadRequest(w, err.Error())
		return
	}

	task, err := h.service.GetTaskByID(*id)	

	if (err != nil){
		log.Println(err)
		writeBadRequest(w, err.Error())
		return
	}

	if (task.Status == model.Done){
		w.Header().Set("Location", "/task_result/"+task.TaskID.String())
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.NewGetTaskByIDOutput("", true, task.Email, task.Status))
}

func (h *TaskHandler) HandleTaskResultID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	switch r.Method{
		case http.MethodGet:
			h.handleTaskResultIDGet(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) handleTaskResultIDGet(w http.ResponseWriter, r *http.Request){
	id, err := getIDFromPathParam(r)

	if (err != nil){
		log.Println(err)
		writeBadRequest(w, err.Error())
		return	
	}

	task, err := h.service.GetTaskByID(*id)
	
	if (err != nil){
		log.Println(err)
		writeBadRequest(w, err.Error())
		return
	}

	if (task.Status != model.Done){
		message := "Task is not done"
		log.Println(message)
		writeBadRequest(w, message)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.NewGetTaskResultByIDOutput("", true, task.Email, task.Data))
}

func NewTaskHanlder(service service.TaskService) TaskHandler{
	return TaskHandler{service: service}
}


