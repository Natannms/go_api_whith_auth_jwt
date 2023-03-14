//go:build !release
// +build !release

package docs

import (
	"github.com/Natannms/go_api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Handler retorna a documentação Swagger para a API
var Handler = httpSwagger.WrapHandler

// Init gera a documentação Swagger para a API
func Init() {
	docs.SwaggerInfo.Title = "Go BASIC API"
	docs.SwaggerInfo.Description = "API to manage users and authenticate them using JWT tokens."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/"

}
