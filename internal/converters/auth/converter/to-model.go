package converter

import (
	"github.com/ELRAS1/auth/internal/models/auth/model"
	"github.com/ELRAS1/auth/pkg/auth"
)

func LoginToModel(req *auth.LoginRequest) *model.LoginRequest {
	return &model.LoginRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
}

func GetRefreshTokenToModel(req *auth.GetRefreshTokenRequest) *model.GetRefreshTokenRequest {
	return &model.GetRefreshTokenRequest{
		OldRefreshToken: req.GetOldRefreshToken(),
	}
}

func GetAccessTokenToModel(req *auth.GetAccessTokenRequest) *model.GetAccessTokenRequest {
	return &model.GetAccessTokenRequest{
		RefreshToken: req.GetRefreshToken(),
	}
}
