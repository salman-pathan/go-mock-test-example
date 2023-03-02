package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mockingtest/domain"
	"mockingtest/service"
)

type userRoute struct {
	rg          gin.IRoutes
	userService service.UserService
}

func NewUserRoute(rg gin.IRoutes, userService service.UserService) userRoute {
	return userRoute{
		rg:          rg,
		userService: userService,
	}
}

func (r *userRoute) HandleRoutes() {
	r.rg.POST("/add", r.addUserHandler)
}

func (r *userRoute) addUserHandler(ctx *gin.Context) {
	var request domain.AddUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var resp domain.AddUserResponse
	userID, err := r.userService.AddUser(request)
	if err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	resp.UserId = userID
	ctx.JSON(http.StatusOK, resp)
}
