package dto

import (
	"github.com/gin-gonic/gin"
	"time"
)

type comment struct {
	Id          int
	UserId      int
	videoId     int
	commentText string
	createDate  time.Time
}

func AddComment(context *gin.Context) {

}

func RemoveComment(context *gin.Context) {

}

func UpdateComment(context *gin.Context) {

}

func ShowComment(context *gin.Context) {

}
