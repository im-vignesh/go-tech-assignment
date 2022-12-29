package types

type Employee struct {
	ID           int64  `json:"id"`
	Name         string `json:"employee_name"`
	Salary       int64  `json:"employee_salary"`
	Age          int64   `json:"age"`
	ProfileImage string `json:"profile_image"`
}

type GetEmployeeResponse struct {
	Status  string   `json:"status"`
	Data    Employee `json:"data"`
	Message string   `json:"message"`
}