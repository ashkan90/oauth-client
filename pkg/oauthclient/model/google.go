package model

type GoogleResource struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
}

func (g GoogleResource) ToGeneric() *Generic {
	return &Generic{
		Email:         g.Email,
		VerifiedEmail: g.VerifiedEmail,
	}
}
