package errors

import "github.com/kataras/iris/v12"

func NotFound(ctx iris.Context)  {
	ctx.WriteString("Oups Not Found")
}

func InternalServerError(ctx iris.Context)  {
	ctx.WriteString("Oups something went wrong, try again")
}