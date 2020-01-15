package models

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type (
	Role struct {
		Value  string
		Pillar string
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
			return echo.NewHTTPError(http.StatusUnauthorized, Unauthorized())
		}
		return next(c)
	}
}