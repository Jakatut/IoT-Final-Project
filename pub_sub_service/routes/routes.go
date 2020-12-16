package routes

/*
 *
 * file: 		routes.go
 * project:		logging_service - NAD-A3
 * programmer: 	Conor Macpherson
 * description: Defines routes used in the logging service and initializes the logger, cors, and jwt token authentication.
 *
 */

import (
	"pub_sub_service/config"
	"pub_sub_service/handlers"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Setup configures the router and assigned routes their middleware & handlers.
//
// Parameters:
//	*gin.Engine				router		- gin router
//	*core.FileMutexPool		mutexPool	- contains Read/Write mutexes for each log type.
//	*core.LogTypeCounter	counters	- contains id counters for each log type
//
func Setup(router *gin.Engine) {

	// Add logger, cross origin restrictions.
	router.Use(
		gin.Logger(),
		cors.New(cors.Config{
			AllowMethods:     []string{"POST", "GET"},
			AllowHeaders:     []string{"Content-Type", "Origin", "Accept", "Authorization", "*"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return true
			},
		}),
	)

	router.GET("/readings", func(c *gin.Context) {
		handlers.ReadingHandler(c.Writer, c.Request)
	})

	port := strconv.Itoa(config.GetConfig().Server.Port)
	router.Run(":" + port)
}
