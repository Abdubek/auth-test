// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Logged struct {
	Viewer *Viewer `json:"viewer"`
	Token  *Token  `json:"token"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Sites struct {
	ID   string  `json:"id"`
	Host *string `json:"host"`
}

type Token struct {
	AccessToken  *string `json:"accessToken"`
	RefreshToken *string `json:"refreshToken"`
}
