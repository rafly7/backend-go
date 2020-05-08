package errors

import "github.com/kataras/iris/v12"

type Errors interface {
	NotFound() func(iris.Context)
	InternalServerError() func(iris.Handler)
}
