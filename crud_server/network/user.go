package network

import (
	"crud-server/service"
	"crud-server/types"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once
	userROuterInstance *userRouter
)

type userRouter struct {
	router *Network

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userROuterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerPOST("/", userROuterInstance.create)
		router.registerGET("/", userROuterInstance.get)
		router.registerUPDATE("/", userROuterInstance.update)
		router.registerDELETE("/", userROuterInstance.delete)
	})

	return userROuterInstance
}

func (u *userRouter) create(c *gin.Context) {
	u.userService.Create(nil)

	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("Success", 1),
	})
}

func (u *userRouter) get(c *gin.Context) {

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("Success", 1),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	u.userService.Update(nil, nil)

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("Success", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	u.userService.Delete(nil)

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("Success", 1),
	})
}
