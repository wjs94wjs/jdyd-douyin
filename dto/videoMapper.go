package dto

import (
	"github.com/gin-gonic/gin"
	"time"
)

type video struct {
	Id         int
	AuthorId   int
	PlayUrl    string
	CoverUrl   string
	PublicTime time.Time
	Title      string
}

func AddVideo(context *gin.Context) {

}

func RemoveVideo(context *gin.Context) {

}

func UpdateVideo(context *gin.Context) {

}

func ShowVideo(context *gin.Context) {

}
