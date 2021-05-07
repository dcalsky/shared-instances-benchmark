package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CustomContext struct {
	echo.Context
	Content string
}

type CustomRoute struct {
	Server  *echo.Echo
	Content string
}

func (r *CustomRoute) Hello(c echo.Context) error {
	return c.String(http.StatusOK, r.Content)
}

var customContext = &CustomContext{
	Content: "Custom",
}

// Passing extended custom context
func BenchmarkContext(b *testing.B) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			customContext.Context = c
			return next(customContext)
		}
	})
	e.GET("/", func(c echo.Context) error {
		cc := c.(*CustomContext)
		return c.String(http.StatusOK, cc.Content)
	})
	benchmarkRoutes(b, e)
}

// Manual dependency injection in the routes
func BenchmarkSingleton(b *testing.B) {
	e := echo.New()
	customRoute := CustomRoute{
		Server:  e,
		Content: "Custom",
	}
	e.GET("/", customRoute.Hello)
	benchmarkRoutes(b, e)
}

func benchmarkRoutes(b *testing.B, router http.Handler) {
	b.Helper()
	b.ReportAllocs()
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(w, r)
	}
}
