package models

import (
	"github.com/Viva-con-Agua/echo-pool/pool"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

type (
	RoleCreate struct {
		Name   string `json:"name" validate:"required"`
		Pillar string `json:"pillar" validate:"required"`
	}
	Role struct {
		Uuid   string `json:"uuid" validate:"required"`
		Name   string `json:"name" validate:"required"`
		Pillar string `json:"pillar" validate:"required"`
	}
	Roles struct {
		Role []Role
	}
)

func (roles *Roles) AddRole(role Role) []Role {
	roles.Role = append(roles.Role, role)
	return roles.Role
}

func (r *Roles) Permission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["roles"] == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, pool.Unauthorized())
		}
		return next(c)
	}
}
