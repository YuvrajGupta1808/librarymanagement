package router

import (
	"yuvraj/project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	bookGroup := r.Group("/books")
	{
		bookGroup.POST("/", controllers.CreateBook)
		bookGroup.GET("/", controllers.GetBooks)
		bookGroup.GET("/:id", controllers.GetBookByID)
		bookGroup.PUT("/:id", controllers.UpdateBook)
		bookGroup.DELETE("/:id", controllers.DeleteBook)
	}
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/", controllers.GetAllUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}
	issueGroup := r.Group("/issues")
	{
		issueGroup.POST("/", controllers.IssueBook)
		issueGroup.DELETE("/returnbook", controllers.ReturnBook)
		issueGroup.PUT("/reissue", controllers.ReissueBook)
	}
	fineGroup := r.Group("/fine")
	{
		fineGroup.GET("/:userID/:bookID", controllers.CalculateFine)
		fineGroup.PUT("/:userID/:bookID", controllers.UpdateFine)
	}
	//r.GET("/return-book", controllers.ReturnBook)
	r.Run(":8080")

	return r
}
