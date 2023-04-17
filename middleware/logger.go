package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s]  %s  %s  %d   %s \n",
			param.ClientIP,                      //request
			param.TimeStamp.Format(time.RFC822), //request
			param.Method,                        //request
			param.Path,                          //request
			param.StatusCode,                    //response
			param.Latency)                       //response
	})
}
