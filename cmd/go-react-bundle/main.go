package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

func main() {
	// Create new instance of echo
	e := echo.New()

	// Setup our assetHandler and point it to our static build location
	assetHandler := http.FileServer(rice.MustFindBox("../../web/go-react-bundle/build").HTTPBox())

	// Setup a new echo route to load the build as our base path
	e.GET("/", echo.WrapHandler(assetHandler))

	// Serve our static assists under the /static/ endpoint
	e.GET("/static/*", echo.WrapHandler(assetHandler))

	// Test API route
	e.GET("/api/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start echo instance on 1323 port
	e.Logger.Fatal(e.Start(":1323"))
}
