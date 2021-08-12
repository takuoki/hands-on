package main

import (
	"database/sql"
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=root password=root dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := &handler{db: db}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/employees", h.listEmployees)

	e.Logger.Fatal(e.Start(":1323"))
}

type employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Position   string  `json:"position"`
	Department string  `json:"department"`
	Salary     float32 `json:"salary"`
}

type handler struct {
	db *sql.DB
}

func (h *handler) listEmployees(c echo.Context) error {

	dep := c.QueryParam("department")

	rows, err := h.db.Query(
		"SELECT id, name, position, department, salary " +
			"FROM employees " +
			"WHERE secret = FALSE AND department = '" + dep + "'")

	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	defer rows.Close()

	employees := []employee{}
	for rows.Next() {
		var id int
		var name string
		var position string
		var department string
		var salary float32
		if err := rows.Scan(&id, &name, &position, &department, &salary); err != nil {
			log.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}
		employees = append(employees, employee{
			ID:         id,
			Name:       name,
			Position:   position,
			Department: department,
			Salary:     salary,
		})
	}

	if err := rows.Close(); err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, employees)
}
