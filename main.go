package main

import (
	"first-api/Config"
	"first-api/Models"
	"first-api/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

// @title GoAPI+Swag API
// @version 1.0
// @description GoAPI+Swag Learning

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /user-api
// @query.collection.format multi

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run(":8080")
}
