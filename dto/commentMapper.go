package dto

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

type Comment struct {
	Id          int64     `gorm:"column:id"`
	UserId      int64     `gorm:"column:user_id"`
	VideoId     int64     `gorm:"column:video_id"`
	CommentText string    `gorm:"column:create_text"`
	CreateDate  time.Time `gorm:"column:create_date"`
}

func init() {
	//配置MySQL连接参数
	username := "root"  //账号
	password := "123"   //密码
	host := "localhost" //数据库地址，可以是Ip或者域名
	port := 3306        //数据库端口
	Dbname := "douyin"  //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db

}

func AddComment(context *gin.Context) {
	comment := Comment{Id: context.GetInt64("id"),
		UserId:      context.GetInt64("userId"),
		VideoId:     context.GetInt64("video_id"),
		CommentText: context.GetString("create_text"),
		CreateDate:  context.GetTime("create_date")}

	DB.NewRecord(comment)
}

func RemoveComment(context *gin.Context) {

}

func UpdateComment(context *gin.Context) {

}

func ShowComment(context *gin.Context) {

}
