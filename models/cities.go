package models

import (
	"github.com/Viva-con-Agua/echo-pool/pool"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

type (
	CityCreate struct {
		Name   string `json:"name" validate:"required"`
		Pillar string `json:"pillar" validate:"required"`
	}
	City struct {
		Uuid   string `json:"uuid" validate:"required"`
		Name   string `json:"name" validate:"required"`
		Pillar string `json:"pillar" validate:"required"`
	}
	Citys struct {
		City []City
	}
)

func (Citys *Citys) AddCity(City City) []City {
	Citys.City = append(Citys.City, City)
	return Citys.City
}

func (r *Citys) Permission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["Citys"] == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, pool.Unauthorized())
		}
		return next(c)
	}
}
