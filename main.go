package main

import (
	"../webFrameworkGin/db"
	helper "./helpers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Person struct {
	id        int    `json:"id"`
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
	username  string `json:"username"`
	password  string `json:"password"`
}

type Response struct {
	code    int     `json:"code"`
	status  string  `json:"status"`
	message string  `json:"message"`
	person  *Person `json:"person"`
}

func main() {
	router := gin.Default()

	router.GET("/", indexGet)
	router.GET("/users/:name", getUser)
	router.GET("/users/:name/*action", getUserAction)
	router.POST("/login", userAuthentication)

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

	c.JSON(http.StatusOK, gin.H{"id": id, "name": firstname, "lastname": lastname})

}

func userAuthentication(c *gin.Context) {
	log.Println("userAuthentication()")

	//response := Response{}
	var response Response

	connection, errorConnection := db.GetDatabaseConnection("root", "root", "127.0.0.1", "3307", "gotest")
	if errorConnection != nil {
		panic(errorConnection.Error())
	}

	defer connection.Close()

	errorConnection = connection.Ping()
	if errorConnection != nil {
		panic(errorConnection.Error())
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Println("Username: ", username)
	log.Println("password: ", password)

	if helper.IsEmpty(username) || helper.IsEmpty(password) {
		response := Response{http.StatusBadRequest, "error", "Incomplete parameter", nil}
		log.Println(response)
		c.JSON(response.code, gin.H{"status": "error", "code": http.StatusBadRequest, "message": "Incomplete parameter"})
	} else if username == "username" && password == "password" {

		c.JSON(response.code, gin.H{"status": "SUCCESS", "code": http.StatusOK, "message": "OK"})
		//selectDatabase, selectError := connection.Query("SELECT * FROM person WHERE username = ? AND password = ?", username, passsword)
		//if selectError != nil {
		//	panic(selectError.Error())
		//}
		//
		//person := new(Person)
		//
		//log.Println("Person ",person)
		//
		//for selectDatabase.Next() {
		//	selectError = selectDatabase.Scan(&person.id, &person.firstname, &person.lastname, &person.username, &person.password)
		//	if selectError != nil {
		//		panic(selectError.Error())
		//	}
		//
		//	log.Println(fmt.Println(person))
		//}
		//
		//response = Response{http.StatusOK, "success", "", person}
		//log.Println(response)
	} else {
		c.JSON(response.code, gin.H{"status": "FAIL", "code": http.StatusNoContent, "message": "login fail"})
	}
	//c.JSON(response.code, response)
}

func indexGet(c *gin.Context) {
	log.Println("indexGet()")

	c.JSON(http.StatusOK, gin.H{"message": "index"})
}
