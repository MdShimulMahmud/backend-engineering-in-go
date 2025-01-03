package models

type Todo struct {
	ID        int    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string `json:"title"`
	Completed bool   `json:"completed"`
}