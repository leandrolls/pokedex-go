package model

type Api struct {
	Moves []move       	   `json:"moves"`
}

type move struct {
	Name string `json:"name"`
}
