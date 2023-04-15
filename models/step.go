package models

type Step struct {
	StepId      int         `json:"step_id"`
	Name        string      `json:"name"`
	Key         string      `json:"key"`
	Description string      `json:"description"`
	Rect        string      `json:"rec"`
	Parameters  []Parameter `json:"parameters"`
}
