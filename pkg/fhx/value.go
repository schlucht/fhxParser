package fhx

type Value struct {
	CV     string
	Origin bool
	Date   string
}

type History struct {
	OriginValue   Value
	HistoryValues []Value
}
