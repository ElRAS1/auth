package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	RefreshToken string `json:"refresh_token"`
}

type GetRefreshTokenRequest struct {
	OldRefreshToken string `json:"old_refresh_token"`
}

type GetRefreshTokenResponse struct {
	RefreshToken string `json:"refresh_token"`
}

type GetAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}
