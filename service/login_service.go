package service

import (
	"context"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
)

type LoginService interface {
	DoLogin(ctx context.Context, loginDtoReq dto.LoginDtoReq) map[string]any
	DoRefreshToken(ctx context.Context, refreshDtoReq dto.RefreshDtoReq) map[string]any
}
