package servers

import (
	"github.com/kataras/iris/v12/middleware/logger"
)

var requestLogger = logger.New(logger.Config{
	Status:             true,
	IP:                 true,
	Method:             true,
	Path:               true,
	Query:              true,
	MessageContextKeys: []string{"logger_message"},
	MessageHeaderKeys:  []string{"User-Agent"},
})
