package entity

type Employee struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Experience   int    `json:"experience"`
	Gender       string `json:"gender"`
	PrevEmployer string `json:"prevEmployer"`
}
