package main

import (
	"github.com/gin-gonic/gin"
	//"fmt"
	"golang/api"
	"github.com/jinzhu/gorm"
	"golang/db"
	"strconv"
	"golang/commands"
)

func ApiMiddleware(db* gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {
	commands.Load_data()

	//r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	_db := db.Database()

	r.Use(ApiMiddleware(_db))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users/:user", GetUsersHandler)
	r.POST("/users/:user", PostUsersHandler)
	r.GET("/users/:user/:path", GetUsersVisitsHandler)

	r.GET("/visits/:visit", GetVisitsHandler)
	r.POST("/visits/:visit", PostVisitsHandler)

	r.GET("/locations/:location", GetLocationsHandler)
	r.GET("/locations/:location/avg", GetLocationsAvgHandler)
	r.POST("/locations/:location", PostLocationsHandler)

	r.Run(":80")
}


func PostUsersHandler(c *gin.Context) {
	path1 := c.Param("user")
	if path1 == "new" {
		api.CreateUser(c)
	} else {
		userId, _ := strconv.Atoi(path1)
		api.UpdateUser(c, userId)
	}
}

func PostVisitsHandler(c *gin.Context) {
	path1 := c.Param("visit")

	if path1 == "new" {
		api.CreateVisit(c)
	} else {
		userId, _ := strconv.Atoi(path1)
		api.UpdateVisit(c, userId)
	}
}

func PostLocationsHandler(c *gin.Context) {
	path1 := c.Param("location")

	if path1 == "new" {
		api.CreateLocation(c)
	} else {
		userId, _ := strconv.Atoi(path1)
		api.UpdateLocation(c, userId)
	}
}

func GetUsersHandler(c *gin.Context) {
	id := c.Param("user")
	userId, _ := strconv.Atoi(id)
	api.GetUser(c, userId)
}

func GetUsersVisitsHandler(c *gin.Context) {
	id := c.Param("user")
	userId, _ := strconv.Atoi(id)
	api.GetUserVisits(c, userId)
}

func GetVisitsHandler(c *gin.Context) {
	id := c.Param("visit")
	userId, _ := strconv.Atoi(id)
	api.GetVisit(c, userId)
}

func GetLocationsHandler(c *gin.Context) {
	id := c.Param("location")
	userId, _ := strconv.Atoi(id)
	api.GetLocation(c, userId)
}

func GetLocationsAvgHandler(c *gin.Context) {
	id := c.Param("location")
	userId, _ := strconv.Atoi(id)
	api.GetLocationAvg(c, userId)
}