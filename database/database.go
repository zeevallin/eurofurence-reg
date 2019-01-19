package database

type Database struct {
	Keyword     string       `json:"keyword"`
	Registrants []Registrant `json:"registrants"`
}

type Registrant struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	Arrival     string `json:"arrival"`
	Departure   string `json:"departure"`
	Comment     string `json:"comment"`
	RoomType    string `json:"room_type"`
	Occupancy   int    `json:"occupancy"`
	Occupants   []Occupants
}

type Occupants struct {
	BadgeID     string `json:"badge_id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Town        string `json:"Town"`
	PostCode    string `json:"post_code"`
	Country     string `json:"country"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
