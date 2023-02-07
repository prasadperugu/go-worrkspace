package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct{
	ID    string `json:"id"`
	Item  string `json:"title"`
	Completed bool `json:"completed"`

}

var todos = [] todo{
	{ID:"1",Item:"From Mars",Completed:false},
	{ID:"2",Item:"From Saturn",Completed:false},
	{ID:"3",Item:"From Prasad's Dream",Completed:false},
}
	func getTodos(context *gin.Context){
		context.IndentedJSON(http.StatusOK, todos)

	}

	func main() {
		router  := gin.Default() //creating a server
		router.GET("/todos",getTodos)
		router.Run("localhost:8080")
	}