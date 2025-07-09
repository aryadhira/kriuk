package models

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (e *Employee) TableName() string {
	return "employee"
}
