package routing

import "github.com/labstack/echo/v4"

func RoutingAPI(echo *echo.Echo) {
	// Start Server
	echo.Logger.Fatal(echo.Start(":8001"))
}