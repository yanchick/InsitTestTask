package model

import "time"

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Task struct {
	Login       string    `json:"login"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Status      string    `json:"status"`
}
