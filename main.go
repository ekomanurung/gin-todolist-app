package main

import (
	"fmt"

	"gin-todolist/configuration"
	"gin-todolist/docs"
	"github.com/gin-gonic/gin"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Todolist API"
	docs.SwaggerInfo.Description = "Go todolist service using swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	// Initialize Configurations
	configuration.ConfigureLogLevel()
	configuration.ConfigureDependencies(r)

	fmt.Printf("WANT TO CHECK this")

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Panic when starting the web server caused by: %v\n", err.Error()))
	}
}
