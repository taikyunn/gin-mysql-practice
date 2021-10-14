package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/taisa831/sandbox-gin/controllers"
)

func Router(dbConn *gorm.DB) {
	todoHandler := controllers.todoHandler{
		Db: dbConn,
	}

	r := gin.Default()
	r.LoadHTMLGlob("templetes/*")
	r.GET("/todo", todoHandler.GetAll)                 // 一覧画面
	r.POST("/todo", todoHandler.CreateTask)            // 新規作成
	r.GET("/todo/:id", todoHandler.EditTask)           // 編集画面
	r.POST("/todo/delete/:id", todoHandler.DeleteTask) // 削除
	r.Run()
}
