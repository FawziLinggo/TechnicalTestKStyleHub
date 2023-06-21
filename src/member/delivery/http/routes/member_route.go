package routes

import (
	"TechnicalTestKStyleHub/src/member/delivery/http/handlers"
	"github.com/labstack/echo/v4"
)

type MemberRoute struct {
	app           *echo.Echo
	MemberHandler handlers.MemberHandler
}

func NewMemberRoute(app *echo.Echo, MemberHandler handlers.MemberHandler) MemberRoute {
	return MemberRoute{
		app:           app,
		MemberHandler: MemberHandler,
	}
}

func (route MemberRoute) RegisterRoutes() {

	// root
	memberRoute := route.app.Group("/")

	// route
	memberRoute.GET("member", route.MemberHandler.GetMember)
	memberRoute.POST("member", route.MemberHandler.CreateMember)
	memberRoute.PUT("member", route.MemberHandler.EditMemberByID)
	memberRoute.DELETE("member", route.MemberHandler.DeleteMemberByID)

}
