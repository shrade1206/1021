package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// type Pic struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// 	Date []byte `json:"date"`
// }

// func main() {
// 	var (
// 		db  *gorm.DB
// 		err error
// 		pic []Pic
// 	)
// 	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/camdb?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		log.Fatal("Failed To Connect : ", err.Error())
// 	}
// 	log.Println("連線成功 ")

// 	server := gin.Default()
// 	server.GET("/cam", func(c *gin.Context) {
// 		err := db.Find(&pic)
// 		if err != nil {
// 			c.AbortWithStatus(404)
// 			log.Printf("查詢失敗 : %s", err.Error)
// 		} else {
// 			c.JSON(http.StatusOK, &pic)
// 		}
// 	})
// 	// server.GET("/cam/:id", selectOne)

// 	err = server.Run(":8080")
// 	if err != nil {
// 		log.Fatal("8080 err : ", err.Error())
// 	}

// }

// func mySql() {

// }

// func selectAll(c *gin.Context) {

// }

// func selectOne(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	err := db.First(&pic, id).Error
// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		log.Printf("查詢失敗 : %s", err.Error())
// 	} else {
// 		c.JSON(http.StatusOK, &pic)
// 	}
// }
