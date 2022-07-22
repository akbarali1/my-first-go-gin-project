package Controllers

import (
	"awesomeProject/Models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func GetUser(c *gin.Context) {
	var UserModel []Models.UserModel
	_, err := dbmap.Select(&UserModel, "select * from user")

	if err == nil {
		c.JSON(200, UserModel)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func GetUserDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var UserModel Models.UserModel
	err := dbmap.SelectOne(&UserModel, "SELECT * FROM user WHERE id=? LIMIT 1", id)

	if err == nil {
		userId, _ := strconv.ParseInt(id, 0, 64)

		content := &Models.UserModel{
			Id:        userId,
			Username:  UserModel.Username,
			Password:  UserModel.Password,
			Firstname: UserModel.Firstname,
			Lastname:  UserModel.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func Login(c *gin.Context) {
	var UserModel Models.UserModel
	c.Bind(&UserModel)
	err := dbmap.SelectOne(&UserModel, "select * from user where Username=? LIMIT 1", UserModel.Username)

	if err == nil {
		userId := UserModel.Id

		content := &Models.UserModel{
			Id:        userId,
			Username:  UserModel.Username,
			Password:  UserModel.Password,
			Firstname: UserModel.Firstname,
			Lastname:  UserModel.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func PostUser(c *gin.Context) {
	var UserModel Models.UserModel
	c.Bind(&UserModel)

	log.Println(UserModel)

	if UserModel.Username != "" && UserModel.Password != "" && UserModel.Firstname != "" && UserModel.Lastname != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO user (Username, Password, Firstname, Lastname) VALUES (?, ?, ?, ?)`, UserModel.Username, UserModel.Password, UserModel.Firstname, UserModel.Lastname); insert != nil {
			userId, err := insert.LastInsertId()
			if err == nil {
				content := &Models.UserModel{
					Id:        userId,
					Username:  UserModel.Username,
					Password:  UserModel.Password,
					Firstname: UserModel.Firstname,
					Lastname:  UserModel.Lastname,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var UserModel Models.UserModel
	err := dbmap.SelectOne(&UserModel, "SELECT * FROM user WHERE id=?", id)

	if err == nil {
		var json Models.UserModel
		c.Bind(&json)

		userId, _ := strconv.ParseInt(id, 0, 64)

		user := Models.UserModel{
			Id:        userId,
			Username:  UserModel.Username,
			Password:  UserModel.Password,
			Firstname: json.Firstname,
			Lastname:  json.Lastname,
		}

		if user.Firstname != "" && user.Lastname != "" {
			_, err = dbmap.Update(&user)

			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
