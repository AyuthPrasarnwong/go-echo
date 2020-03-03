package models

type (
	SubDistrictTrans struct {
		EN string `json:"en"`
		TH string `json:"th"`
	}
	DistrictTrans struct {
		EN string `json:"en"`
		TH string `json:"th"`
	}
	ProvinceTrans struct {
		EN string `json:"en"`
		TH string `json:"th"`
	}
	// Geo Geo model
	Geo struct {
		ADLevel       int64            `gorm:"column:AD_LEVEL" json:"level"`
		SubDistrictID int64            `gorm:"column:TA_ID" json:"sub_district_id"`
		SubDistrict   SubDistrictTrans `json:"sub_district"`
		SubDistrictEN string           `gorm:"column:TAMBON_E" json:"-"`
		SubDistrictTH string           `gorm:"column:TAMBON_T" json:"-"`
		DistrictID    int64            `gorm:"column:AM_ID" json:"district_id"`
		District      DistrictTrans    `json:"district"`
		DistrictEN    string           `gorm:"column:AMPHOE_E" json:"-"`
		DistrictTH    string           `gorm:"column:AMPHOE_T" json:"-"`
		ProvinceID    int64            `gorm:"column:CH_ID" json:"province_id"`
		Province      ProvinceTrans    `json:"province"`
		ProvinceEN    string           `gorm:"column:CHANGWAT_E" json:"-"`
		ProvinceTH    string           `gorm:"column:CHANGWAT_T" json:"-"`
		Lat           float64          `gorm:"column:LAT" json:"latitude"`
		Long          float64          `gorm:"column:LONG" json:"longitude"`
	}
	// Province Provinces Model
	Province struct {
		ProvinceID int64         `gorm:"column:CH_ID" json:"id"`
		Province   ProvinceTrans `json:"name"`
		ProvinceEN string        `gorm:"column:CHANGWAT_E" json:"-"`
		ProvinceTH string        `gorm:"column:CHANGWAT_T" json:"-"`
	}
	// District District Model
	District struct {
		DistrictID int64         `gorm:"column:AM_ID" json:"id"`
		ProvinceID int64         `gorm:"column:CH_ID" json:"province_id"`
		District   DistrictTrans `json:"name"`
		DistrictEN string        `gorm:"column:AMPHOE_E" json:"-"`
		DistrictTH string        `gorm:"column:AMPHOE_T" json:"-"`
	}

	// SubDistrict Sub District Model
	SubDistrict struct {
		SubDistrictID int64            `gorm:"column:TA_ID" json:"id"`
		DistrictID    int64            `gorm:"column:AM_ID" json:"district_id"`
		ProvinceID    int64            `gorm:"column:CH_ID" json:"province_id"`
		SubDistrict   SubDistrictTrans `json:"name"`
		SubDistrictEN string           `gorm:"column:TAMBON_E" json:"-"`
		SubDistrictTH string           `gorm:"column:TAMBON_T" json:"-"`
	}
)

// TableName set table name for Geo
func (x *Geo) TableName() string {
	return "risk_tambon"
}

// AfterFind set language after find
func (x *Geo) AfterFind() (err error) {
	x.SubDistrict.EN = x.SubDistrictEN
	x.SubDistrict.TH = x.SubDistrictTH
	x.District.EN = x.DistrictEN
	x.District.TH = x.DistrictTH
	x.Province.EN = x.ProvinceEN
	x.Province.TH = x.ProvinceTH
	return
}

// TableName set table name for Geo
func (x *Province) TableName() string {
	return "risk_tambon"
}

// AfterFind set language after find
func (x *Province) AfterFind() (err error) {
	x.Province.EN = x.ProvinceEN
	x.Province.TH = x.ProvinceTH
	return
}

// TableName set table name for Geo
func (x *District) TableName() string {
	return "risk_tambon"
}

// AfterFind set language after find
func (x *District) AfterFind() (err error) {
	x.District.EN = x.DistrictEN
	x.District.TH = x.DistrictTH
	return
}

// TableName set table name for Geo
func (x *SubDistrict) TableName() string {
	return "risk_tambon"
}

// AfterFind set language after find
func (x *SubDistrict) AfterFind() (err error) {
	x.SubDistrict.EN = x.SubDistrictEN
	x.SubDistrict.TH = x.SubDistrictTH
	return
}
