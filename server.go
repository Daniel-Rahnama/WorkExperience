package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Element struct {
	Value string `json:"value"`
}

type Properties struct {
	Name     string             `json:"name"`
	Elements map[string]Element `json:"elements"`
}

type DATA struct {
	Properties Properties `json:"properties"`
}

func Insert(c echo.Context) error {
	data := new(DATA)
	c.Bind(data)

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ArticlesDB")

	if (err != nil) {
		panic(err.Error)
	}

	defer db.Close()

	for Heading, Value := range data.Properties.Elements {
		query, err := db.Query(`INSERT INTO Articles VALUES ("` + data.Properties.Name + `", "` + Heading + `", "` + Value.Value + `")`)

		if (err != nil) {
			panic(err.Error)
		}
	
		defer query.Close()
	}

	return c.String(http.StatusOK, "INSERTED TO TABLE")
}

func Display(c echo.Context) error {
	Name := c.Param("name")

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ArticlesDB")

	if (err != nil) {
		panic(err.Error)
	}

	defer db.Close()

	query, err := db.Query(`SELECT Heading, Value FROM Articles WHERE Name = "` + Name + `"`)

	if (err != nil) {
		panic(err.Error)
	}

	defer query.Close()

	data := DATA{Properties:Properties{Name:Name, Elements:map[string]Element{}}}

	for (query.Next()) {
		var Heading string
		var Value string

		err := query.Scan(&Heading, &Value)

		if (err != nil) {
			panic(err.Error)
		}

		data.Properties.Elements[Heading] = Element{Value:Value}
	}

	return c.JSON(http.StatusOK, data)
}

func Delete(c echo.Context) error {
	Name := c.FormValue("name")

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ArticlesDB")

	if (err != nil) {
		panic(err.Error)
	}

	defer db.Close()

	query, err := db.Query(`DELETE FROM Articles WHERE Name = "` + Name + `"`)

	if (err != nil) {
		panic(err.Error)
	}

	defer query.Close()

	return c.String(http.StatusOK, "DELETED FROM TABLE")
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.POST("/insert", Insert)

	e.GET("/:name", Display)

	e.POST("/delete", Delete)

	e.Logger.Fatal(e.Start(":1323"))
}