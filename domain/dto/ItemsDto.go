package dto

type Items struct {
	TotalItems int    `json:"total_items"`
	Items      []Item `json:"items"`
}

type Item struct {
	Weight int    `json:"weight"`
	Type   string `json:"type"`
}
