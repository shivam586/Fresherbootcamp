//Routes/Routes.go
package Routes
import (
	"StudentRecord/Contrllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-record")
	{
		grp1.GET("student-data", Contrllers.GetStudentData)
		grp1.POST("student", Contrllers.CreateStudent)
		grp1.GET("student/:id", Contrllers.GetStudentByID)
		grp1.PUT("student/:id", Contrllers.UpdateStudent)
		grp1.DELETE("student/:id", Contrllers.DeleteStudent)
	}
	return r
}