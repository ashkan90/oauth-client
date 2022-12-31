package model

type Generic struct {
	Token         string `json:"-"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
}
