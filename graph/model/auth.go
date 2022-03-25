package model

type UsersMutation struct {
	Login   *interface{} `json:"login"`
	Logout  *interface{} `json:"logout"`
	Refresh *interface{} `json:"refresh"`
}

type Viewer struct {
	ID    string         `json:"id"`
	Email *string        `json:"email"`
	Sites []*interface{} `json:"sites"`
}
