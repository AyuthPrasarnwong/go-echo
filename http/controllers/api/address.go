package api

import (
	"net/http"

	"github.com/jinzhu/gorm"

	"app/exceptions"
	"app/http/requests/api/test"
	"app/models"

	"github.com/labstack/echo"
)

type (
	// AddressAPIController SMS api controller
	AddressAPIController struct {
		Controller
	}
)

func (ctl *AddressAPIController) ExValidate(c echo.Context) (err error) {
	request := new(test.TestStoreRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "validate success",
	})
}

// Provinces list province
func (ctl *AddressAPIController) Provinces(c echo.Context) (err error) {
	model := []models.Province{}
	ctl.DB(nil).Select([]string{"CH_ID", "CHANGWAT_E", "CHANGWAT_T"}).Group("CH_ID").Find(&model)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model,
	})
}

// ShowProvince get provinces detail
func (ctl *AddressAPIController) ShowProvince(c echo.Context) (err error) {
	model := models.Province{}
	if err := ctl.DB(nil).Select([]string{"CH_ID", "CHANGWAT_E", "CHANGWAT_T"}).
		Where("CH_ID = ?", c.Param("id")).
		Group("CH_ID").
		First(&model).Error; gorm.IsRecordNotFoundError(err) {
		return &exceptions.ErrorException{
			Message:  "Not found.",
			ErrorKey: "not-found",
			Code:     http.StatusNotFound,
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model,
	})
}

// Districts list district
func (ctl *AddressAPIController) Districts(c echo.Context) (err error) {
	model := []models.District{}
	ctl.DB(nil).Select([]string{"CH_ID", "AM_ID", "AMPHOE_E", "AMPHOE_T"}).
		Where("CH_ID = ?", c.Param("province_id")).
		Group("AM_ID").
		Find(&model)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model,
	})
}

// ShowDistrict list district
func (ctl *AddressAPIController) ShowDistrict(c echo.Context) (err error) {
	model := models.District{}
	if err := ctl.DB(nil).
		Select([]string{"CH_ID", "AM_ID", "AMPHOE_E", "AMPHOE_T"}).
		Where("CH_ID = ? AND AM_ID = ?", c.Param("province_id"), c.Param("id")).
		Group("AM_ID").
		First(&model).Error; gorm.IsRecordNotFoundError(err) {
		return &exceptions.ErrorException{
			Message:  "Not found.",
			ErrorKey: "not-found",
			Code:     http.StatusNotFound,
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model,
	})
}

// SubDistricts list district
func (ctl *AddressAPIController) SubDistricts(c echo.Context) (err error) {
	model := []models.SubDistrict{}
	ctl.DB(nil).
		Where("CH_ID = ? AND AM_ID = ?", c.Param("province_id"), c.Param("district_id")).
		Group("TA_ID").
		Find(&model)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model,
	})
}

// ShowSubDistrict list district
func (ctl *AddressAPIController) ShowSubDistrict(c echo.Context) (err error) {
	model := models.SubDistrict{}
	if err := ctl.DB(nil).
		Where("CH_ID = ? AND AM_ID = ? AND TA_ID = ?", c.Param("province_id"), c.Param("district_id"), c.Param("id")).
		Group("TA_ID").
		First(&model).Error; gorm.IsRecordNotFoundError(err) {
		return &exceptions.ErrorException{
			Message:  "Not found.",
			ErrorKey: "not-found",
			Code:     http.StatusNotFound,
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": model,
	})
}
