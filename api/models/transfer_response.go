package models

type Card struct {
	CardName   string `json:"card_name"`
	Percentage string `json:"percentage"`
}

type Cards struct {
	Cards []Card `json:"cards"`
}
