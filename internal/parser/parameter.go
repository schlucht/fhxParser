package parser

type Parameter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       Value  `json:"value"`
	UnitID      int    `json:"unit_id"`
}

type StepParameter struct {
	Name    string `json:"name"`
	Origin  string `json:"origin"`
	DeferTo string `json:"deferto"`
	Group   string `json:"group"`
}
