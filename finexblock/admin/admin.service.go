package admin

import "context"

type AdminService struct {
	repo *AdminRepository
}

func (a *AdminService) Ctx() context.Context {
	return context.Background()
}

func (a *AdminService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}