package main

import (
	"net/http"

	"github.com/h8gi/blgo/controllers"
	"github.com/h8gi/blgo/models"
	"github.com/h8gi/blgo/views"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("postgres",
		"host=localhost user=yagihiroki dbname=gomi sslmode=disable password=mypassword")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Fragment{})
	controllers.SetDB(db)

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
	e.GET("/", controllers.ShowIndex)

	api := e.Group("/api")
	api.Use(middleware.CORS())

	api.GET("/fragments", controllers.GetFragmentsList)
	api.POST("/fragments", controllers.PostFragment)

	api.GET("/fragments/:id", controllers.GetFragment)
	api.PUT("/fragments/:id", controllers.UpdateFragment)
	api.DELETE("/fragments/:id", controllers.DeleteFragment)

	e.Logger.Fatal(e.Start(":1323"))
}
