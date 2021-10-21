package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/camdb")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	type Pic struct {
		Id   int
		Name string
		Date []byte
	}
	r := gin.New()
	r.GET("/person/:id", func(c *gin.Context) {
		var (
			pic    Pic
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id, name, date from pic where id = ?;", id)
		err = row.Scan(&pic.Id, &pic.Name, &pic.Date)
		if err != nil {
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": pic,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})
	err = r.Run(":8080")
	if err != nil {
		log.Printf("Server Error %s", err.Error())
	}
}
