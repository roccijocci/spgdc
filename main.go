package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Discipleship struct {
	Date            time.Time `json: "date"`
	Topic           string    `json: "topic"`
	Text            string    `json: "text"`
	MemoryVerse     string    `json: "memoryverse"`
	Introduction    string    `json: "introduction"`
	LessonOutline   string    `json: "lessonoutline"`
	Questions       string    `json: "questions"`
	LifeApplication string    `json: "lifeapplication"`
}

func main() {
	router := gin.Default()
	router.Run()
}
