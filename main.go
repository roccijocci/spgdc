package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DiscipleshipClass struct {
	Id              string    `json: "_id"`
	Date            string    `json: "date"`
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
var ctx context.Context
var err error
var client *mongo.Client

func init() {
	discipleshipClasses = make([]DiscipleshipClass, 0)
	file, _ := ioutil.ReadFile("discipleshipclass.json")
	_ = json.Unmarshal([]byte(file), &discipleshipClasses)

	ctx = context.Background()
	client, err = mongo.Connect(ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	var listOfDiscipleshipClasses []interface{}
	for _, dsc := range discipleshipClasses {
		listOfDiscipleshipClasses = append(listOfDiscipleshipClasses, dsc)
	}
	collection := client.Database(os.Getenv(
		"MONGO_DATABASE")).Collection("discipleshipclasses")
	InsertManyResult, err := collection.InsertMany(ctx, listOfDiscipleshipClasses)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	log.Println("Inserted DiscipleshipClass:",
		len(InsertManyResult.InsertedIDs))
}

func NewDCHandler(c *gin.Context) {
	var discipleshipClass DiscipleshipClass

	if err := c.ShouldBindJSON(&discipleshipClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	discipleshipClass.Id = xid.New().String()
	//discipleshipClass.Date = time.Date(2020, time.August, 10, 0, 0, 0, 0, time.Now().Local().Location())
	discipleshipClass.PublishedAt = time.Now()
	discipleshipClasses = append(discipleshipClasses, discipleshipClass)
	c.JSON(http.StatusOK, discipleshipClass)

}

//handler endpoint lists
func ListDCHandler(c *gin.Context) {
	c.JSON(http.StatusOK, discipleshipClasses)
}

//updates
func UpdateDCHandler(c *gin.Context) {
	id := c.Param("id")
	var discipleshipClass DiscipleshipClass

	if err := c.ShouldBindJSON(&discipleshipClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	index := -1
	for i := 0; i < len(discipleshipClasses); i++ {
		if discipleshipClasses[i].Id == id {
			index = i
		}
	}
	if index == 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Discipleship class Memo not found"})
		return
	}
	discipleshipClasses[index] = discipleshipClass
	c.JSON(http.StatusOK, discipleshipClass)
}

func main() {
	router := gin.Default()
	//router.POST("/dc", NewDCHandler)
	router.GET("/dc", ListDCHandler)
	router.PUT("/dc/:id", UpdateDCHandler)
	router.Run()
}
