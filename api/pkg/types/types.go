package types

import "time"

type JobData struct {
	Module string            `json:"module"`
	Inputs map[string]string `json:"inputs"`
}

type Job struct {
	ID        string    `json:"id"`
	Created   time.Time `json:"created"`
	Owner     string    `json:"owner"`
	OwnerType string    `json:"owner_type"`
	State     string    `json:"state"`
	Status    string    `json:"status"`
	Data      JobData   `json:"data"`
}
