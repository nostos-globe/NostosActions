package models

type Trip struct {
	TripID      uint   `json:"TripID"`
	UserID      uint   `json:"user_id""`
	Name        string `json:"name"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
	StartDate   string `json:"start_date" `
	EndDate     string `json:"end_date"`
}

type TripMedia struct {
	Trip  Trip    `json:"trip"`
	Media []Media `json:"media"`
}
