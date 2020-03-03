package contracts

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"
	"errors"

	"app/bootstrap"
	"app/models"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type (
	AppContext struct {
		echo.Context
	}
)

// Auth get logedin profile
func (x *AppContext) Auth() (*models.User, error) {
	auth := x.Get("user")
	if auth == nil {
		return nil, errors.New("Unauthorized")
	}
	user := auth.(*models.User)
	return user, nil
}

// URL setup url
func (x *AppContext) URL(u string) string {
	return fmt.Sprintf("%s/%s", viper.GetString("web_url"), u)
}

// With set flash
func (x *AppContext) With(k string, v interface{}) *AppContext {
	x.SetFlash(k, v)
	return x
}

// WithInputs set flash
func (x *AppContext) WithInputs() *AppContext {
	x.Request().ParseForm()
	x.SetFlash("inputs", x.Request().Form)
	return x
}

// Back back to parent page
func (x *AppContext) Back() error {
	return x.Redirect(http.StatusMovedPermanently, x.Request().Header.Get("Referer"))
}

// WithError set message flash
func (x *AppContext) WithError(s string) *AppContext {
	x.SetFlash("error", s)
	return x
}

// WithErrors set errors list flash
func (x *AppContext) WithErrors(errors []string) *AppContext {
	x.SetFlash("errors", errors)
	return x
}

// SetFlash set flash data
func (x *AppContext) SetFlash(name string, value interface{}) *AppContext {
	var j []byte
	j, err := json.Marshal(value)
	if err != nil {
		rt := reflect.TypeOf(value)
		switch rt.Kind() {
		case reflect.String:
			j = []byte(value.(string))
		default:
			panic("unknow value type")
		}
	}
	cookie := http.Cookie{
		Name:  name,
		Value: x.encode(j),
	}
	_ = cookie
	x.SetCookie(&cookie)
	return x
}

// HasFlash get flash data and delete it
func (x *AppContext) HasFlash(name string) bool {
	if cookie, err := x.Cookie(name); err == nil && cookie != nil {
		return true
	}
	return false
}

// GetFlash get flash data and delete it
func (x *AppContext) GetFlash(name string) (interface{}, error) {
	var cookie *http.Cookie
	var err error

	// x.Response().After(func() {
	// 	dc := http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0)}
	// 	go x.SetCookie(&dc)
	// })

	if cookie, err = x.Cookie(name); err == nil {
		dc := http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0)}
		go x.SetCookie(&dc)

		b, err := x.decode(cookie.Value)
		var value interface{}
		if err = json.Unmarshal(b, &value); err != nil {
			return b, nil
		}

		return value, nil
	}
	return nil, err
}

func (x *AppContext) encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func (x *AppContext) decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}

func (x *AppContext) RespError(code int, key string, i map[string]interface{}) error {
	i["type"] = key
	rd := bootstrap.RedisDB{}
	errorData, _ := rd.Redis(nil).Get("error_message").Result()
	msgs := make(map[string]models.ErrorMessage)
	json.Unmarshal([]byte(errorData), &msgs);
	if errMsg, ok := msgs[key]; ok {
		lang := x.Request().Header.Get("Accept-Language")
		if lang == "" {
			lang = "th"
		}
		var errorCode int64
		switch t := errMsg.Code.(type) {
		case string:
			errorCode, _ = strconv.ParseInt(t, 10, 64)
		default:
			errorCode = 500
		}
		i["code"] = errorCode
		if lang == "en" {
			i["title"] = errMsg.Message.EN
			i["detail"] = errMsg.Detail.EN
		} else {
			i["title"] = errMsg.Message.TH
			i["detail"] = errMsg.Detail.TH
		}
	} else {
		i["code"] = 500
		i["title"] = ""
		if v, ok := i["message"]; ok {
			i["message"] = v
		}
		i["detail"] = ""
	}

	return x.JSON(code, i)
}
