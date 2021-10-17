package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var router = mux.NewRouter()

func StartApplication() {
	mapUrls()
	log.Println("about to start the application...")
	log.Fatal(http.ListenAndServe(":8072", router))
}
