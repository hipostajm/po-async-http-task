package main

import (
	"net/http"
	"github.com/hipostajm/po-async-http-task/long-polling/handler"
	"github.com/hipostajm/po-async-http-task/long-polling/repository"
	"github.com/hipostajm/po-async-http-task/long-polling/service"
)

func main(){
	repository := repository.NewMemoryRepository()
	service := service.NewTaskService(repository)
	handler := handler.NewTaskHanlder(service)
		
	http.HandleFunc("/task", handler.HandleTask)
	http.HandleFunc("/task/{id}", handler.HandleTaskID)
	http.HandleFunc("/task_result/{id}", handler.HandleTaskResultID)

	http.ListenAndServe(":8080", nil)
}
