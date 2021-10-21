package main

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"gocv.io/x/gocv"
// )

// type Pic struct {
// 	Id   int    `json:"id" form:"id"`
// 	Name string `json:"name" form:"name"`
// 	Date []byte `json:"date" form:"date"`
// }

// var upgrader = &websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// func main() {
// 	fmt.Println("Go WebSocket")
// 	mySql()
// 	server := gin.New()
// 	server.GET("/ws", wsEndpoint)

// 	server.GET("/cam", selectAll)
// 	server.GET("/cam/:id", selectOne)
// 	// server.POST("/cam/insert", insert)
// 	// server.PUT("/cam/update/:id", update)
// 	// server.DELETE("/cam/:id", del)

// 	server.NoRoute(gin.WrapH(http.FileServer(http.Dir("./public"))), func(c *gin.Context) {
// 		path := c.Request.URL.Path
// 		method := c.Request.Method
// 		fmt.Println(path)
// 		fmt.Println(method)
// 		//檢查path的開頭使是否為"/"
// 		if strings.HasPrefix(path, "/") {
// 			fmt.Println("ok")
// 		}
// 	})
// 	err := server.Run(":8080")
// 	if err != nil {
// 		log.Fatal("8080 err : ", err.Error())
// 	}
// }

// var (
// 	db  *gorm.DB
// 	err error
// 	pic []Pic
// )

// func mySql() {
// 	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/camdb?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		log.Fatal("Failed To Connect : ", err.Error())
// 	}
// 	db.AutoMigrate(&pic)
// }
// func selectAll(c *gin.Context) {
// 	db.Find(&pic)
// 	c.JSON(http.StatusOK, &pic)
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

// func insert(c *gin.Context) {

// }
// func update(c *gin.Context) {

// }
// func del(c *gin.Context) {

// }

// func wsEndpoint(c *gin.Context) {
// 	// 透過http請求程序調用upgrader.Upgrade，來獲取*Conn (代表WebSocket連接)
// 	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println("使用者已連線")

// 	var newImg []byte

// 	for {
// 		messageType, p, err := ws.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		if string(p) == "run" {

// 			func() {
// 				//設定視訊鏡頭，0 = 預設鏡頭
// 				webcam, err := gocv.VideoCaptureDevice(0)
// 				if err != nil {
// 					log.Println(err)
// 				}

// 				time.Sleep(time.Second)
// 				img := gocv.NewMat()
// 				defer img.Close()

// 				webcam.Read(&img)
// 				defer webcam.Close()
// 				//設定副檔名、來源
// 				buf, err := gocv.IMEncode(".jpg", img)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				defer buf.Close() //nolint
// 				//設定變數取得暫存檔案
// 				newImg = buf.GetBytes()
// 				// d, _ := os.ReadFile(a)
// 				//轉換成base64的字串型別
// 				data := base64.StdEncoding.EncodeToString(newImg)
// 				//把轉換好的字串傳送到前端，前端接收在轉換回圖片
// 				if err := ws.WriteMessage(messageType, []byte(data)); err != nil {
// 					log.Println(err)
// 					return
// 				}
// 			}()

// 		}
// 		if string(p) == "save" {
// 			//用來生成新文件使用(檔名、來源、)
// 			err := os.WriteFile("demo.jpg", newImg, os.ModePerm)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		}
// 		log.Println("使用者訊息: " + string(p))
// 	}
// }
