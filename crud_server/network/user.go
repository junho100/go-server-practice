package network

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once
	userROuterInstance *userRouter
)

type userRouter struct {
	router *Network
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userROuterInstance = &userRouter{
			router: router,
		}

		router.registerPOST("/", userROuterInstance.create)
		router.registerGET("/", userROuterInstance.get)
		router.registerUPDATE("/", userROuterInstance.update)
		router.registerDELETE("/", userROuterInstance.delete)
	})

	return userROuterInstance
}

func (u *userRouter) create(c *gin.Context) {
}

func (u *userRouter) get(c *gin.Context) {
}

func (u *userRouter) update(c *gin.Context) {
}

func (u *userRouter) delete(c *gin.Context) {
}
