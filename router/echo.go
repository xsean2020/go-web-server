package router

import (
	"reflect"

	echo "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func AddRoutersToEcho(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	ctx := reflect.TypeOf((*echo.Context)(nil)).Elem()
	for ttype, controller := range controllers {
		g := e.Group(controller.Group(), controller.Middlewares()...)
		// 基于反射包装包装echo注册
		routers := parseRouters(ttype, controller)
		for _, router := range routers {
			// 参数检查
			if router.handler.Type.In(1).Implements(ctx) {
				g.Add(router.method, router.path, func(ctx echo.Context) error {
					err, _ := router.handler.Func.Call([]reflect.Value{reflect.ValueOf(controller), reflect.ValueOf(ctx)})[0].Interface().(error)
					return err
				})
				log.WithField("method", router.method).WithField("path", controller.Group()+router.path).WithField("controller", ttype.String()).Info("register router")
			}
		}
	}
}
