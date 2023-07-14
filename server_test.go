package main

import (
	"testing"
	"strings"
	"net/http"
	"net/http/httptest"
	"net/url"
    "github.com/labstack/echo/v4"
)

var (
	jData = `{"properties":{"name":"test","elements":{"a1":{"value":"a2"},"a2":{"value":"a3"},"a3":{"value":"a4"},"a4":{"value":"a5"}}}}`
)

func TestInsert (t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/insert")

	err := Insert(c)

	if ((rec.Body.String() != "INSERTED TO TABLE") || (err != nil)) {
		t.Fatalf("Insert():\n%v\nWant:\nINSERTED TO TABLE\nError:\n%v", rec.Body.String(), err.Error())
	}
}

func TestDisplay (t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("test")

	err := Display(c)

	if ((strings.Compare(strings.Replace(rec.Body.String(), "\n", "", -1), strings.Replace(string(jData), "\n", "", -1)) != 0) || (err != nil)) {
		t.Fatalf("Display():\n%v\nWant:\n%v\nError:%v", rec.Body.String(), string(jData), err.Error())
	}
}

func TestDelete (t *testing.T) {
	e := echo.New()

	f := make(url.Values)
	f.Set("name", "test")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	c.SetPath("/delete")

	err := Delete(c)

	if ((rec.Body.String() != "DELETED FROM TABLE") || (err != nil)) {
		t.Fatalf("Delete():\n%v\nWant:\nDELETED FROM TABLE\nError:\n%v", rec.Body.String(), err.Error())
	}
}