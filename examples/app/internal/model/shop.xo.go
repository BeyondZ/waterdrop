// Package model contains the types for schema 'swoyii'.
package model

// Code generated by xo. DO NOT EDIT.
import (
	"strings"
)

// Shop represents a row from 'swoyii.shop'.
type Shop struct {
	ID uint64 `json:"id"`

	BrandID uint64 `json:"brand_id"`

	ShopName string `json:"shop_name"`

	Address string `json:"address"`

	ShopStatus int8 `json:"shop_status"`

	OfficialStartTime uint32 `json:"official_start_time"`

	HasOfficialRun int8 `json:"has_official_run"`

	IsValid int8 `json:"is_valid"`

	FrozenStartTime uint32 `json:"frozen_start_time"`

	FrozenEndTime uint32 `json:"frozen_end_time"`

	IsFrozen int8 `json:"is_frozen"`

	HolidayStartTime uint32 `json:"holiday_start_time"`

	HolidayEndTime uint32 `json:"holiday_end_time"`

	IsHoliday int8 `json:"is_holiday"`

	ExpireTime uint32 `json:"expire_time"`

	Description string `json:"description"`

	IsAllday int8 `json:"is_allday"`

	ProvinceID uint32 `json:"province_id"`

	CityID uint32 `json:"city_id"`

	DistrictID uint32 `json:"district_id"`

	Lat float64 `json:"lat"`

	Lng float64 `json:"lng"`

	AlbumID uint64 `json:"album_id"`

	IsDel int8 `json:"is_del"`

	CreatedTime uint32 `json:"created_time"`

	UpdatedTime uint32 `json:"updated_time"`

	ProvinceName string `json:"province_name"`

	CityName string `json:"city_name"`

	DistrictName string `json:"district_name"`

	OperatorID uint64 `json:"operator_id"`

	OperatorType int8 `json:"operator_type"`
}

// Get table name
func GetShopTableName() string {
	return "shop"
}

// Get field string
func GetShopFieldString() string {
	return `id, brand_id, shop_name, address, shop_status, official_start_time, has_official_run, is_valid, frozen_start_time, frozen_end_time, is_frozen, holiday_start_time, holiday_end_time, is_holiday, expire_time, description, is_allday, province_id, city_id, district_id, lat, lng, album_id, is_del, created_time, updated_time, province_name, city_name, district_name, operator_id, operator_type`
}

// Get field string slice
func GetShopFieldStringSlice() []string {
	return strings.Split(`id, brand_id, shop_name, address, shop_status, official_start_time, has_official_run, is_valid, frozen_start_time, frozen_end_time, is_frozen, holiday_start_time, holiday_end_time, is_holiday, expire_time, description, is_allday, province_id, city_id, district_id, lat, lng, album_id, is_del, created_time, updated_time, province_name, city_name, district_name, operator_id, operator_type`, ",")
}

// Get add field
func GetShopAddField(s *Shop) []interface{} {
	return []interface{}{
		s.ID,
		s.BrandID,
		s.ShopName,
		s.Address,
		s.ShopStatus,
		s.OfficialStartTime,
		s.HasOfficialRun,
		s.IsValid,
		s.FrozenStartTime,
		s.FrozenEndTime,
		s.IsFrozen,
		s.HolidayStartTime,
		s.HolidayEndTime,
		s.IsHoliday,
		s.ExpireTime,
		s.Description,
		s.IsAllday,
		s.ProvinceID,
		s.CityID,
		s.DistrictID,
		s.Lat,
		s.Lng,
		s.AlbumID,
		s.IsDel,
		s.CreatedTime,
		s.UpdatedTime,
		s.ProvinceName,
		s.CityName,
		s.DistrictName,
		s.OperatorID,
		s.OperatorType,
	}
}

// Get scan field
func GetShopScanField(s *Shop) []interface{} {
	return []interface{}{
		&s.ID,
		&s.BrandID,
		&s.ShopName,
		&s.Address,
		&s.ShopStatus,
		&s.OfficialStartTime,
		&s.HasOfficialRun,
		&s.IsValid,
		&s.FrozenStartTime,
		&s.FrozenEndTime,
		&s.IsFrozen,
		&s.HolidayStartTime,
		&s.HolidayEndTime,
		&s.IsHoliday,
		&s.ExpireTime,
		&s.Description,
		&s.IsAllday,
		&s.ProvinceID,
		&s.CityID,
		&s.DistrictID,
		&s.Lat,
		&s.Lng,
		&s.AlbumID,
		&s.IsDel,
		&s.CreatedTime,
		&s.UpdatedTime,
		&s.ProvinceName,
		&s.CityName,
		&s.DistrictName,
		&s.OperatorID,
		&s.OperatorType,
	}
}
