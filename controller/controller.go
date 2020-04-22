package controller

import echo "github.com/labstack/echo/v4"

type Controller interface {
	Group() string
	Middlewares() []echo.MiddlewareFunc
	Alias() map[string]string // 路由别名  key: method_name  value:path
}
