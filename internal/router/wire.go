package router

import "github.com/google/wire"

var RouterSet = wire.NewSet(
	NewUserRouter,
	NewAuthRouter,
	ProvideRouters,
)

func ProvideRouters(userRouter *UserRouter, authRouter *AuthRouter) []Router {
	return []Router{userRouter, authRouter}
}
