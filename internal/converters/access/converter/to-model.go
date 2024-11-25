package converter

import (
	"github.com/ELRAS1/auth/internal/models/access/model"
	"github.com/ELRAS1/auth/pkg/access"
)

func CheckToModel(req *access.CheckRequest) *model.CheckRequest {
	return &model.CheckRequest{
		EndpointAddress: req.GetEndpointAddress(),
	}
}
