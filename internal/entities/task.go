package entities

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Due         string `json:"due"`
	Status      string `json:"status"`
}
