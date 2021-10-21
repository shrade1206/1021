package main

import (
	"bytes"
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
		log.Printf("SQL Error : %s", err.Error())
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Printf("Ping Error : %s", err.Error())
		return
	}

	r := gin.New()
	//Select One
	r.GET("/get/:id", func(c *gin.Context) {
		var (
			pic    Pic
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("SELECT id, name, picture FROM pic WHERE id = ?;", id)
		err = row.Scan(&pic.Id, &pic.Name, &pic.Picture)
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
	//Select All
	r.GET("/get", func(c *gin.Context) {
		var (
			pic  Pic
			pics []Pic
		)
		rows, err := db.Query("SELECT id, name, picture FROM pic;")
		if err != nil {
			log.Printf("Select Error : %s", err.Error())
			return
		}
		for rows.Next() {
			err = rows.Scan(&pic.Id, &pic.Name, &pic.Picture)
			pics = append(pics, pic)
			if err != nil {
				log.Printf("Select Error :%s", err.Error())
				return
			}
			defer rows.Close()
			c.JSON(http.StatusOK, gin.H{
				"result": pics,
				"count":  len(pics),
			})
		}
	})
	//Insert
	r.POST("/pic", func(c *gin.Context) {
		var buffer bytes.Buffer
		name := c.PostForm("name")
		picture := c.PostForm("picture")
		stmt, err := db.Prepare("INSERT into pic(name,picture)VALUE(?,?);")
		if err != nil {
			log.Printf("Insert Error : %s", err.Error())
			return
		}
		buffer.WriteString(name)
		buffer.WriteString(" ")
		buffer.WriteString(picture)
		defer stmt.Close()
		insertName := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s Insert success", insertName),
		})
	})
	//Post
	r.PUT("/person/put", func(c *gin.Context) {
		var buffer bytes.Buffer
		// id := c.Query("id")
		picture := c.PostForm("picture")
		stmt, err := db.Prepare("UPDATE pic set picture= ? where id= ?;")
		if err != nil {
			log.Printf("Update fail : %s", err.Error())
		}
		_, err = stmt.Exec(picture)
		if err != nil {
			log.Printf("Exec fail : %s", err.Error())

		}
		buffer.WriteString(picture)
		defer stmt.Close()
		PutName := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", PutName),
		})
	})
	//DEL
	r.DELETE("/del", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("DELETE FROM pic where id=?;")
		if err != nil {
			log.Printf("Del fail : %s", err.Error())
			return
		}
		_, err = stmt.Exec(id)
		if err != nil {
			log.Printf("")
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Success Del id : %s", id),
		})
	})

	err = r.Run(":8080")
	if err != nil {
		log.Printf("Server Error %s", err.Error())
	}
}

type Pic struct {
	Id      int
	Name    string
	Picture string
}

// var db *sql.DB

//Select One
// func (p Pic) get() (pic Pic, err error) {
// 	row := db.QueryRow("SELECT id, name, date FROM pic WHERE id =?", p.Id)
// 	err = row.Scan(&p.Id, &p.Name, &p.Date)
// 	if err != nil {
// 		log.Printf("Select Error : %s", err.Error())
// 		return
// 	}
// 	return
// }

//Select All
// func (p Pic) getAll() (pics []Pic, err error) {
// 	rows, err := db.Query("SELECT id, name, date FROM pic")
// 	if err != nil {
// 		return
// 	}
// 	for rows.Next() {
// 		var pic Pic
// 		rows.Scan(&pic.Id, &pic.Name, &pic.Date)
// 		pics = append(pics, pic)
// 	}
// 	defer rows.Close()
// 	return
// }
