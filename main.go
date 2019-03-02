package main

import (
	"../webFrameworkGin/db"
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

	connection, errorConnection := db.GetDatabaseConnection("root", "root", "127.0.0.1", "3307", "gotest")
	if errorConnection != nil {
		panic(errorConnection.Error())
	}
	defer connection.Close()

	errorConnection = connection.Ping()
	if errorConnection != nil {
		panic(errorConnection.Error())
	}

	name := c.Param("name")
	log.Print("name: ", name)

	selDB, selError := connection.Query("SELECT * FROM person WHERE first_name=?", name)
	if selError != nil {
		panic(selError.Error())
	}

	var id int
	var firstname, lastname string

	for selDB.Next() {

		selError = selDB.Scan(&id, &firstname, &lastname)
		if selError != nil {
			panic(selError.Error())
		}
		log.Print("id ", id)
		log.Print("name ", firstname)
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "name": name, "lastname": lastname})

}

func indexGet(c *gin.Context) {
	log.Println("indexGet()")

	c.JSON(http.StatusOK, gin.H{"message": "index"})
}
