package main_test

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"testing"
	"time"
)

func TestEcho(t *testing.T) {
	e := echo.New()

	g := e.Group("/g")
	g.GET("/skills", func(ctx echo.Context) error {
		return ctx.String(202, "static skills")
	})
	g.GET("/status", func(ctx echo.Context) error {
		return ctx.String(201, "static status")
	})
	g.GET("/:name", func(ctx echo.Context) error {
		return ctx.String(200, ctx.Param("name"))
	})

	go func() {
		t.Fatal(e.Start(":8007"))
	}()
	time.Sleep(time.Second)

	resp, err := http.Get("http://localhost:8007/g/s")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("ERROR, status code must be 200, but it is %v", resp.StatusCode)
	}
}
