package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"app/bootstrap"
	contracts "app/contracts"
	handlers "app/handlers"
	"app/routes"

	validations "app/validators/validations"

	echotemplate "github.com/foolin/echo-template"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	gl "github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

var (
	uni       *ut.UniversalTranslator
	templates = template.New("")
)

type (

	// Validator struct
	Validator struct {
		validator *validator.Validate
	}

	// TemplateRenderer is a custom html/template renderer for Echo framework
	TemplateRenderer struct {
		templates *template.Template
	}

	// Host enable for using subdomain
	Host struct {
		Echo *echo.Echo
	}

	Req struct {
	}
)

// Validate validate request data
func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func init() {
	// set config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("Error reading config file, %s", err)
	}
}

// Main start server
func main() {
	// init log
	filename := fmt.Sprintf("storage/logs/debug-%s.log", time.Now().Format("2006-01-02"))
	lf, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	defer lf.Close()

	if err != nil {
		logger.Fatalf("cannot open '%s', (%s)", filename, err.Error())
		flag.Usage()
		os.Exit(-1)
	}

	log.SetOutput(lf)

	defer logger.Init("initLog", true, true, lf).Close()

	// define sub host if enabled
	hosts := map[string]*Host{}

	appPort := viper.GetString("app_port")
	// web app
	app := createServer(lf)

	bootstrap.InitialDatabases()

	// apply web routing
	routes.ConfigureHandlerHTTP(app)
	// apply api routing
	// define api subdomain (if enabled)
	if viper.GetBool("subDomainAPI.enabled") == false {
		routes.ConfigureHandlerHTTPAPI(app.Group("api"))
		app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", appPort)))
	} else {
		api := createServer(lf)

		routes.ConfigureHandlerHTTPAPI(api.Group(""))
		hosts[fmt.Sprintf("%s:%s", viper.GetString("subDomainAPI.subdomain"), appPort)] = &Host{api}
		hosts[fmt.Sprintf("%s:%s", viper.GetString("app_url"), appPort)] = &Host{app}
		e := echo.New()

		e.Any("/*", func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			host := hosts[req.Host]

			if host == nil {
				err = echo.ErrNotFound
			} else {
				host.Echo.ServeHTTP(res, req)
			}

			return
		})

		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appPort)))
	}
}

func createServer(f *os.File) *echo.Echo {
	debug := viper.GetBool("debug")
	// new server
	e := echo.New()
	// set static file
	e.Static("/", "assets")

	e.Debug = debug
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(gl.ERROR)

	var cc *contracts.AppContext

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc = &contracts.AppContext{c}
			return next(cc)
		}
	})

	const path = "resources/views/"
	const trimSuffix = ".tpl"

	e.Renderer = echotemplate.New(echotemplate.TemplateConfig{
		Root:      "resources/views",
		Extension: ".tpl",
		Master:    "layouts/main",
		Partials:  []string{},
		Funcs: template.FuncMap{
			"isNil": func(t interface{}) bool {
				return t == nil
			},
			"old": func(v string) interface{} {
				if flash, _ := cc.GetFlash("inputs"); flash != nil {
					inputs := flash.(map[string]interface{})
					if value, ok := inputs[v]; ok {
						inputVal := value.([]interface{})
						return inputVal[0]
					}
				}
				return nil
			},
			"has": func(v string) bool {
				return cc.HasFlash(v)
			},
			"flash": func(v string) interface{} {
				flash, _ := cc.GetFlash(v)
				return flash
			},
			"validations": func() interface{} {
				flash, _ := cc.GetFlash("validations")
				return flash
			},
			"errors": func() interface{} {
				flash, _ := cc.GetFlash("errors")
				return flash
			},
			"keyExists": func(key string, arr interface{}) bool {
				if arr != nil {
					array := arr.(map[string]interface{})
					if _, ok := array[key]; ok {
						return true
					}
				}
				return false
			},
		},
		DisableCache: true,
	})

	// setup validator
	validatorHandler := validator.New()
	// register custom validation
	validatorHandler.RegisterValidation("date", validations.DateValidation)
	validatorHandler.RegisterValidation("datetime", validations.DatetimeValidation)
	validatorHandler.RegisterValidation("date_range", validations.DateRangeValidation)
	validatorHandler.RegisterValidation("required_if", validations.RequiredIf)

	e.Validator = &Validator{validator: validatorHandler}

	e.HTTPErrorHandler = handlers.ErrorHandler

	return e
}
