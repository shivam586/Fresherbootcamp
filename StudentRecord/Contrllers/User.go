//Contrllers/User.go
package Contrllers
import (
	"StudentRecord/Models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
//Get Student
func GetStudentData(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllStudent(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//Create Student
func CreateStudent(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateStudent(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//Get Student by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetStudentByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//UpdateStudent... Update the user information
func UpdateStudent(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//DeleteUser ... Delete the user
func DeleteStudent(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}