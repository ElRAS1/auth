package converter

import (
	"github.com/ELRAS1/auth/internal/models/auth/model"
	"github.com/ELRAS1/auth/pkg/auth"
)

func LoginToApi(req *model.LoginResponse) *auth.LoginResponse {
	return &auth.LoginResponse{
		RefreshToken: req.RefreshToken,
	}
}

func GetRefreshTokenToApi(req *model.GetRefreshTokenResponse) *auth.GetRefreshTokenResponse {
	return &auth.GetRefreshTokenResponse{
		RefreshToken: req.RefreshToken,
	}
}

func GetAccessTokenToApi(req *model.GetAccessTokenResponse) *auth.GetAccessTokenResponse {
	return &auth.GetAccessTokenResponse{
		AccessToken: req.AccessToken,
	}
}
