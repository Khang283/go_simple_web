package types

type Student struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

type StudentDeleteRequest struct {
	Id string `json:"id"`
}

type GetStudentResponse struct {
	Status string
	Error  string
	Data   []Student
}

type CreateStudentRequest struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}
