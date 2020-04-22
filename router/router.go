package router

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/xsean2020/go-web-server/controller"
)

type router struct {
	method  string
	path    string
	handler reflect.Method
}

var controllers = map[reflect.Type]controller.Controller{}

func Register(ctr controller.Controller) {
	ttype := reflect.TypeOf(ctr)
	if _, ok := controllers[ttype]; ok {
		panic(fmt.Sprintf("repeate controller [%s]", ttype.Name()))
	}
	controllers[ttype] = ctr
}

func parseRouters(ttype reflect.Type, ctr controller.Controller) []router {
	var routers []router
	alais := ctr.Alias()
	for k, v := range alais {
		alais[strings.ToLower(k)] = strings.ToLower(v)
	}
	num := ttype.NumMethod()
	for i := 0; i < num; i++ {
		fn := ttype.Method(i)
		if fn.Type.NumIn() != 2 {
			continue
		}
		// 反射获取方法注入路由
		var method = http.MethodGet
		var path, prefix string

		funcName := strings.ToLower(fn.Name)
		if m, ok := validHttpMethod[funcName]; ok {
			method = m
			prefix = funcName
		} else {

			if i := strings.Index(funcName, "_"); i > 0 {
				prefix = funcName[:i]
			}

			if m, ok := validHttpMethod[prefix]; ok {
				method = m
			}
		}

		if o, ok := alais[funcName]; ok {
			path = o
		} else {
			// 处理路由注册
			path = strings.ReplaceAll(strings.Replace(funcName, prefix, "", 1), "__", "/:")
			path = strings.ReplaceAll(path, `_`, `/`)
		}
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		routers = append(routers, router{path: path, handler: fn, method: method})
	}
	return routers
}
