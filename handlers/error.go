package validators

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	Contracts "app/contracts"
	"app/exceptions"
	"github.com/fatih/camelcase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

// ErrorHandler handle error exception
func ErrorHandler(err error, c echo.Context) {
	cc := &Contracts.AppContext{c}
	var errorKey string
	isAPI := strings.HasPrefix(c.Request().URL.String(), "/api/")

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%s %d.html", time.Now().Format("2006-01-02"), code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)

	response := make(map[string]interface{})

	title := "Internal Server Error"
	var detail string
	if reflect.TypeOf(err).Elem().Name() == "ErrorException" {
		if exception := err.(*exceptions.ErrorException); exception != nil {
			code = int(exception.GetCode())
			title = exception.GetMessage()
			detail = exception.GetDetail()
			errorKey = exception.GetErrorKey()
			response["title"] = title
			response["detail"] = detail
		}
	} else if reflect.TypeOf(err).Name() == "ValidationErrors" {
		code = http.StatusUnprocessableEntity
		errorKey = "validation-error"

		validate := err
		if validate, ok := validate.(*validator.InvalidValidationError); ok {
			panic(validate)
		}

		//Validation errors occurred
		errors := make(map[string][]string)
		errorLists := make([]string, 0)
		//Use reflector to reverse engineer struct
		for _, validate := range validate.(validator.ValidationErrors) {
			//If json tag doesn't exist, use lower case of name
			namespace := strings.Split(validate.StructNamespace(), ".")
			namespace = namespace[1 : len(namespace)-1]
			nodePrefix := ""
			if len(namespace) > 0 {
				nodePrefix = strings.Join(namespace, ".") + "."
			}
			// fmt.Println(validate.StructField())
			name := validate.StructField()
			fieldName := fmt.Sprintf("%s%s", nodePrefix, strings.ToLower(strings.Join(camelcase.Split(name), "_")))
			keyName := fmt.Sprintf("%s%s", nodePrefix, strings.ToLower(strings.Join(camelcase.Split(name), " ")))
			value := validate.Param()

			switch validate.Tag() {
			case "required":
				errors[fieldName] = append(errors[fieldName], fmt.Sprintf("The %s is required", keyName))
				errorLists = append(errorLists, fmt.Sprintf("The %s is required", keyName))
				break
			case "required_if":
				param := strings.Split(value, `:`)
				paramField := param[0]
				paramValue := param[1]
				errors[fieldName] = append(errors[fieldName], fmt.Sprintf("The %v is required when %v is %v", keyName, paramField, paramValue))
				errorLists = append(errorLists, fmt.Sprintf("The %v is required when %v is %v", keyName, paramField, paramValue))
				break
			case "required_without":
				errors[fieldName] = append(errors[fieldName], fmt.Sprintf("The %v is required when %v is empty", keyName, value))
				errorLists = append(errorLists, fmt.Sprintf("The %v is required when %v is empty", keyName, value))
				break
			case "email":
				errors[fieldName] = append(errors[fieldName], fmt.Sprintf("The %s should be a valid email", keyName))
				errorLists = append(errorLists, fmt.Sprintf("The %s should be a valid email", keyName))
				break
			case "eq":
				errors[fieldName] = append(errors[fieldName], fmt.Sprintf("The %v should equal with %v", keyName, value))
				errorLists = append(errorLists, fmt.Sprintf("The %v should equal with %v", keyName, value))
			case "oneof":
				param := strings.Join(strings.Split(value, ` `), ", ")
				errors[fieldName] = append(errors[fieldName], fmt.Sprintf("The %v does not exist in %v", keyName, param))
				errorLists = append(errorLists, fmt.Sprintf("The %v does not exist in %v", keyName, param))
			case "eqfield":
				errors[fieldName] = append(errors[fieldName], fmt.Sprintf("The %s should be equal to the %s", keyName, validate.Param()))
				errorLists = append(errorLists, fmt.Sprintf("The %s should be equal to the %s", keyName, validate.Param()))
				break
			default:
				errors[fieldName] = append(errors[fieldName], fmt.Sprintf("The %s is invalid", keyName))
				errorLists = append(errorLists, fmt.Sprintf("The %s is invalid", keyName))
				break
			}
		}
		response["title"] = "The given data was invalid!"
		response["invalid-params"] = errors

		if !isAPI {
			cc.WithInputs().SetFlash("validations", errors).Redirect(http.StatusMovedPermanently, c.Request().Header.Get("Referer"))
			return
		}
	} else {
		if isAPI {
			(&echo.Echo{}).DefaultHTTPErrorHandler(err, c)
			return
		}

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			response["title"] = he.Message
			if he.Internal != nil {
				err = fmt.Errorf("%v, %v", err, he.Internal)
			}
			response["title"] = http.StatusText(code)
		} else {
			response["title"] = err.Error()
		}
	}

	if isAPI {
		cc.RespError(code, errorKey, response)
	} else {
		c.Render(code, "errors/error", &response)
	}
}
