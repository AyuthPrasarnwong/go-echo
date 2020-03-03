package routes

import (
	"net/http"

	"app/http/controllers/api"

	"github.com/labstack/echo"
)

//ConfigureHandlerHTTPAPI configures the handler functions
func ConfigureHandlerHTTPAPI(e *echo.Group) {
	e.GET("", func(c echo.Context) error {
		response := map[string]interface{}{
			"message": "Welcome to The World Of API(s)",
		}
		return c.JSON(http.StatusOK, response)
	})
	addressHandler := new(api.AddressAPIController)
	e.GET("/geo/provinces", addressHandler.Provinces)
	e.GET("/geo/provinces/:id", addressHandler.ShowProvince)
	e.GET("/geo/provinces/:province_id/districts", addressHandler.Districts)
	e.GET("/geo/provinces/:province_id/districts/:id", addressHandler.ShowDistrict)
	e.GET("/geo/provinces/:province_id/districts/:district_id/sub-districts", addressHandler.SubDistricts)
	e.GET("/geo/provinces/:province_id/districts/:district_id/sub-districts/:id", addressHandler.ShowSubDistrict)

	e.POST("/example-validate", addressHandler.ExValidate)
	//e.GET("/example ", addressHandler.ShowSubDistrict, middlewarefunction here)

}
