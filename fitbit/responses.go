package fitbit

type AuthResponse struct {
	AccessToken  string `json:"access_token,"`
	RefreshToken string `json:"refresh_token,"`
	UserID       string `json:"user_id,"`
}
