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
