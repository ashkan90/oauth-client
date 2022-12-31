package oauth

type CallbackResponse struct {
	Provider string
	State    string
	Code     string
}
