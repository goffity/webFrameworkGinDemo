package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", indexGet)
	router.GET("/users/:name", getUser)
	router.GET("/users/:name/*action", getUserAction)

	router.Run(":8080")
}
func getUserAction(c *gin.Context) {
	log.Print("getUserAction")

	name := c.Param("name")
	action := c.Param("action")
	c.String(http.StatusOK, "Hello %s %s", name, action)
}
func getUser(c *gin.Context) {
	log.Print("getUser")

	connection, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/cafe_roamer")
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	err = connection.Ping()
	if err != nil {
		panic(err.Error())
	}

	preparement, err := connection.Prepare("SELECT id FROM member WHERE member.name LIKE ?")

	name := c.Param("name")
	log.Print("name: ", name)
	err = preparement.QueryRow(name).Scan(&name)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Print("Member id is ", name)

	c.String(http.StatusOK, "Your ID is %s", name)

}

func indexGet(c *gin.Context) {
	log.Println("indexGet()")

	c.JSON(http.StatusOK, gin.H{"message": "index"})
}
