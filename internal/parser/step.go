package parser

type Step struct {
	Name           string          `json:"name"`
	Key            string          `json:"key"`
	Description    string          `json:"description"`
	Rect           string          `json:"rec"`
	StepParameters []StepParameter `json:"step_parameters"`
}
