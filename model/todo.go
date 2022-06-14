package model

type Auth struct {
	Login string `json:"username"`
}

type Task struct {
	Description string `json:"description"`
	State       string `json:"state"`
}

type Response struct {
	Status string `json:"status"`
	Data   []Task `json:"data"`
}
