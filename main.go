package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	//Serves static files frorm this directory
	app.OnBeforeServe().Add(
		func(e *core.ServeEvent) error {
			e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
			return nil
		})

	app.OnBeforeServe().Add(
		func(e *core.ServeEvent) error {
			// add new "GET /hello" route to the app router (echo)
			e.Router.AddRoute(
				echo.Route{
					Method: http.MethodGet,
					Path:   "/hello",
					Handler: func(c echo.Context) error {
						return c.String(200, "Hello world!")
					},
					Middlewares: []echo.MiddlewareFunc{
						apis.ActivityLogger(app),
					},
				})

			return nil
		})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
