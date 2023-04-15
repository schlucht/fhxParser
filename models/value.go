package models

import "time"

type Value struct {
	ValueId     int       `json:"value_id"`
	StringValue string    `json:"stringvalue,omitempty"`
	Set         string    `json:"value_set,omitempty"`
	High        int       `json:"high,omitempty"`
	Low         int       `json:"low,omitempty"`
	Cv          int       `json:"cv,omitempty"`
	Unit        string    `json:"unit,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
