package main

import (
	"net/http"

	"github.com/h8gi/blgo/api_controllers"
	"github.com/h8gi/blgo/models"
	"github.com/h8gi/blgo/views"
	"github.com/h8gi/blgo/web_controllers"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	// db, err := gorm.Open("postgres",
	// 	"host=localhost user=yagihiroki dbname=gomi sslmode=disable password=mypassword")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.AutoMigrate(&models.Fragment{}).Error; err != nil {
		panic(err)
	}
	api_controllers.SetDB(db)

	e := echo.New()

	// register templates
	t := views.NewTemplate("./views/*.html")
	e.Renderer = t

	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "assets")
	e.GET("/", web_controllers.ShowIndex)
	e.GET("/fragments/:name", web_controllers.FragmentPage)

	api := e.Group("/api")
	api.Use(middleware.CORS())

	api.GET("/fragments", api_controllers.GetFragmentsList)
	api.POST("/fragments", api_controllers.PostFragment)

	api.GET("/fragments/:name", api_controllers.GetFragment)
	api.PUT("/fragments/:name", api_controllers.UpdateFragment)
	api.DELETE("/fragments/:name", api_controllers.DeleteFragment)

	e.Logger.Fatal(e.Start(":1323"))
}
