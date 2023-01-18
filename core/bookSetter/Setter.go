package main

import (
	"context"
	"net/http"
	"log"
	"github.com/chaimakr/book_management_system/core/setter/config"
	"github.com/chaimakr/book_management_system/core/setter/database"
	"github.com/chaimakr/book_management_system/core/setter/handlers"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {

	err := profiler.Start(
		profiler.WithService("book-setter-service"),
		profiler.WithEnv("dev"),
		profiler.WithVersion("0.1.0"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

   tracer.Start()
   defer tracer.Stop()

	conf := config.GetConfig()
	ctx := context.TODO()
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)
	client := &database.BookClient{
		Col: collection,
		Ctx: ctx,
	}

	r := gin.Default()

	// Use the tracer middleware with your desired service name.
	r.Use(gintrace.Middleware("book-setter"))
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/command/metrics")

	// set middleware for gin
	m.Use(r)
	req := http.Request{}

	todos := r.Group("/command")
	{
		todos.POST("/add", handlers.InsertBook(client, &req))
		todos.PATCH("/:id", handlers.UpdateBook(client, &req))
		todos.DELETE("/:id", handlers.DeleteBook(client, &req))
	}

	r.Run(":8080")
}
