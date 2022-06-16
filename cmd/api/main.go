package main

import (
	"boletia/cmd/api/bootstrap"
	_ "boletia/docs"
	_ "boletia/internal/plataform/server/handler/currency"
	log "github.com/sirupsen/logrus"
)

// @title           Currencies
// @version         1.0
// @description     This project query to currency api and change the values.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    BOLETIA
// @contact.email  team@boletia

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
