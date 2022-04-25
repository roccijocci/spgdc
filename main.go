package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type DiscipleshipClass struct {
	Id              string    `json: "id"`
	Date            time.Time `json: "date"`
	Topic           string    `json: "topic"`
	Text            []string  `json: "text"`
	MemoryVerse     []string  `json: "memoryVerse"`
	Introduction    []string  `json: "introduction"`
	LessonOutline   []string  `json: "lessonOutline"`
	Questions       []string  `json: "questions"`
	LifeApplication []string  `json: "lifeApplication"`
	PublishedAt     time.Time `json: "publishedAt`
}

var discipleshipClasses []DiscipleshipClass

func init() {
	discipleshipClasses = make([]DiscipleshipClass, 0)
}

func NewDCHandler(c *gin.Context) {
	var discipleshipClass DiscipleshipClass

	if err := c.ShouldBindJSON(&discipleshipClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	discipleshipClass.Id = xid.New().String()
	discipleshipClass.Date = time.Date(2020, time.August, 10, 0, 0, 0, 0, time.Now().Local().Location())
	discipleshipClass.PublishedAt = time.Now()
	discipleshipClasses = append(discipleshipClasses, discipleshipClass)
	c.JSON(http.StatusOK, discipleshipClass)

}

func main() {
	router := gin.Default()
	router.POST("/dc", NewDCHandler)
	router.Run()
}
