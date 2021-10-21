package main

// import (
// 	"database/sql"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/camdb")
// 	if err != nil {
// 		log.Printf("MySQL Error : %s", err.Error())
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		log.Printf("Ping Error : %s", err.Error())
// 	}

// 	//Table
// 	type Test struct {
// 		Id   int
// 		Name string
// 		// Date string
// 	}

// 	r := gin.Default()

// 	r.GET("/pic/:id", func(c *gin.Context) {
// 		var (
// 			test   Test
// 			result gin.H
// 		)
// 		id := c.Param("id")
// 		row := db.QueryRow("select id, name from person where id = ?;", id)
// 		err = row.Scan(&test.Id, &test.Name)
// 		if err != nil {
// 			result = gin.H{
// 				"result": nil,
// 				"count":  0,
// 			}
// 		} else {
// 			result = gin.H{
// 				"result": test,
// 				"count":  1,
// 			}
// 		}
// 		c.JSON(http.StatusOK, result)
// 	})
// 	err = r.Run(":8080")
// 	if err != nil {
// 		log.Printf("Server Error %s", err.Error())
// 	}
// }
