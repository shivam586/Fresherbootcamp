//Models/User.go
package Models
import (
	"StudentRecord/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//GetAllStudent Fetch all student data
func GetAllStudent(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
//CreateStudent ... Insert New data
func CreateStudent(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
//GetStudentByID ... Fetch only one student data by Id
func GetStudentByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
//UpdateStudent ... Update student
func UpdateStudent(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}
//DeleteStudent ... Delete Student
func DeleteStudent(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
