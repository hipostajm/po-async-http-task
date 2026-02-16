package model

type Output struct{
	Error string
	Sucess bool
}

func NewOutput(error string, sucess bool) Output{
	return Output{Error: error, Sucess: sucess}
}

type GetTaskByIDOutPut struct{
	Output
	Email string
	Status EmailStatus
}

func NewGetTaskByIDOutput(error string, sucess bool, email string, status EmailStatus) GetTaskByIDOutPut{
	return GetTaskByIDOutPut{Output: NewOutput(error, sucess), Email: email, Status: status}
}

type GetTaskResultByIDOutput struct{
	Output
	Email string
	Data string
}

func NewGetTaskResultByIDOutput(error string ,sucess bool, email string, data string) GetTaskResultByIDOutput{
	return GetTaskResultByIDOutput{Output: NewOutput(error, sucess), Email: email, Data: data}
}
