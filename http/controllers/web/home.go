package web

import (
	"net/http"

	"github.com/labstack/echo"

	"app/models"
)

type (
	// HomeController user web controller
	HomeController struct {
		Controller
	}
)

// Index show login page
func (ctl *HomeController) Index(c echo.Context) error {
	auth := c.Get("user")

	data := map[string]interface{}{}
	if auth != nil {
		data["user"] = auth.(*models.User)
	}

	data["message"] = c.QueryParam("message")

	return c.Render(http.StatusOK, "home", data)
}
