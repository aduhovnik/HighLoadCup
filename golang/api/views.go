package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang/models"
	"github.com/jinzhu/gorm"
	"fmt"
	"strconv"
	"time"
	"io/ioutil"
	"strings"
	"encoding/json"
)


func createEntity(c *gin.Context, val interface{}){
	_db:= c.MustGet("db").(*gorm.DB)
	_db.Create(val)
}

func CheckEntityExistence(c *gin.Context, id int, table string) bool{
	_db:= c.MustGet("db").(*gorm.DB)

	type IdStruct struct {
		Id int
	}
	var IdS IdStruct

	sql := fmt.Sprintf("select id from %v where id = %v", table, id)
	_db.Raw(sql).Scan(&IdS)

	if IdS.Id == 0 {
		return false
	}
	return true
}

func CreateUser(c *gin.Context){
	var newData models.User
	x, _ := ioutil.ReadAll(c.Request.Body)

	a := string(x)

	json.Unmarshal([]byte(a), &newData)

	if strings.Contains(a, " null") || a == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(200, gin.H{})

	go createEntity(c, newData)
}

func CreateVisit(c *gin.Context){
	var newData models.Visit
	x, _ := ioutil.ReadAll(c.Request.Body)

	a := string(x)

	json.Unmarshal([]byte(a), &newData)

	if strings.Contains(a, " null") || a == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(200, gin.H{})

	go createEntity(c, newData)
}

func CreateLocation(c *gin.Context){
	var newData models.Location

	x, _ := ioutil.ReadAll(c.Request.Body)

	a := string(x)

	json.Unmarshal([]byte(a), &newData)

	if strings.Contains(a, " null") || a == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(200, gin.H{})

	go createEntity(c, newData)
}

func UpdateUser(c *gin.Context, userId int){
	_db:= c.MustGet("db").(*gorm.DB)
	var user, newData models.User

	_db.Find(&user, userId)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	x, _ := ioutil.ReadAll(c.Request.Body)

	a := string(x)

	json.Unmarshal([]byte(a), &newData)

	if strings.Contains(a, " null") || a == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(200, gin.H{})
	go _db.Model(&user).UpdateColumns(newData)
}

func UpdateVisit(c *gin.Context, vId int){
	_db:= c.MustGet("db").(*gorm.DB)
	var v, newData models.Visit

	_db.Find(&v, vId)

	if v.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	x, _ := ioutil.ReadAll(c.Request.Body)

	a := string(x)

	json.Unmarshal([]byte(a), &newData)

	if strings.Contains(a, " null" ) || a == ""{
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(200, gin.H{})

	go _db.Model(&v).UpdateColumns(newData)
}

func UpdateLocation(c *gin.Context, lId int){
	_db:= c.MustGet("db").(*gorm.DB)
	var l, newData models.Location

	_db.Find(&l, lId)

	if l.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	x, _ := ioutil.ReadAll(c.Request.Body)

	a := string(x)

	json.Unmarshal([]byte(a), &newData)

	if strings.Contains(a, " null") || a == ""{
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(200, gin.H{})

	go _db.Model(&l).UpdateColumns(newData)
}

func GetUser(c *gin.Context, userId int){
	_db:= c.MustGet("db").(*gorm.DB)
	var user models.User
	_db.Find(&user, userId)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, user)
}


func GetUserVisits(c *gin.Context, userId int){

	if ! CheckEntityExistence(c, userId, "users") {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	country := c.DefaultQuery("country", "")
	fromDate := c.DefaultQuery("fromDate", "")
	toDate := c.DefaultQuery("toDate", "")
	toDistance := c.DefaultQuery("toDistance", "")

	sql := fmt.Sprintf("select mark, visited_at, place from visits " +
		"join locations on user = %v and visits.location = locations.id where user = %v ", userId, userId)

	ok := true

	if country != "" {
		add_sql := fmt.Sprintf(" AND locations.country = '%v'", country)
		sql += add_sql
	}

	if fromDate != "" {
		fDate, err := strconv.Atoi(fromDate)
		if err != nil {
			ok = false
		}
		add_sql := fmt.Sprintf(" AND visits.visited_at > %v", fDate)
		sql += add_sql
	}

	if toDate != "" {
		tDate, err := strconv.Atoi(toDate)
		if err != nil {
			ok = false
		}
		add_sql := fmt.Sprintf(" AND visits.visited_at < %v", tDate)
		sql += add_sql
	}

	if toDistance != "" {
		tDist, err := strconv.Atoi(toDistance)
		if err != nil {
			ok = false
		}
		add_sql := fmt.Sprintf(" AND locations.distance < %v", tDist)
		sql += add_sql
	}

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	type UsersVisits struct {
		Visits []struct {
			Mark      int    `json:"mark"`
			VisitedAt int    `json:"visited_at"`
			Place     string `json:"place"`
		} `json:"visits"`
	}

	_db:= c.MustGet("db").(*gorm.DB)

	uVisits := UsersVisits{}
	_db.Raw(sql+" ORDER BY visited_at").Scan(&uVisits.Visits)

	c.Set("Transfer-Encoding", "identity")
	c.JSON(http.StatusOK, gin.H{"visits":uVisits.Visits})
}


func GetVisit(c *gin.Context, vId int){
	_db:= c.MustGet("db").(*gorm.DB)
	var v models.Visit
	_db.Find(&v, vId)

	if v.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, v)
}

func GetLocation(c *gin.Context, lId int){
	_db:= c.MustGet("db").(*gorm.DB)
	var l models.Location
	_db.Find(&l, lId)

	if l.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, l)
}

func GetLocationAvg(c *gin.Context, lId int){

	if ! CheckEntityExistence(c, lId, "locations") {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	fromDate := c.DefaultQuery("fromDate", "")
	toDate := c.DefaultQuery("toDate", "")
	fromAge := c.DefaultQuery("fromAge", "")
	toAge := c.DefaultQuery("toAge", "")
	gender := c.DefaultQuery("gender", "")

	sql := fmt.Sprintf("select round(avg(v.mark), 5) as avg from locations l " +
		" left join visits v on v.location = %v and l.id = %v" +
		" left join users u on v.user = u.id where l.Id = %v ", lId, lId, lId)

	ok := true

	if fromDate != "" {
		fDate, err := strconv.Atoi(fromDate)
		if err != nil {
			ok = false
		}
		add_sql := fmt.Sprintf(" AND v.visited_at > %v", fDate)
		sql += add_sql
	}

	if toDate != "" {
		tDate, err := strconv.Atoi(toDate)
		if err != nil {
			ok = false
		}
		add_sql := fmt.Sprintf(" AND v.visited_at < %v", tDate)
		sql += add_sql
	}

	now := time.Now()

	if fromAge != "" {
		fAge, err := strconv.Atoi(fromAge)
		if err != nil {
			ok = false
		}
		t := now.AddDate(-fAge, 0, 0)
		bTime := t.UnixNano() / int64(time.Second)
		add_sql := fmt.Sprintf(" AND u.birth_date < %v", bTime)
		sql += add_sql
	}

	if toAge != "" {
		tAge, err := strconv.Atoi(toAge)
		if err != nil {
			ok = false
		}
		t := now.AddDate(-tAge, 0, 0)
		tTime := t.UnixNano()/int64(time.Second)
		add_sql := fmt.Sprintf(" AND u.birth_date > %v", tTime)
		sql += add_sql
	}

	if gender != "" {
		if gender != "m" && gender != "f"{
			ok = false
		}
		add_sql := fmt.Sprintf(" AND u.gender = '%v'", gender)
		sql += add_sql
	}

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	_db:= c.MustGet("db").(*gorm.DB)

	type avg_mark struct {
		Avg float64 `json:"avg"`
	}

	var ans avg_mark

	_db.Raw(sql).Scan(&ans)

	c.JSON(http.StatusOK, ans)
}

