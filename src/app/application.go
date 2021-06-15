package app

import (
	"github.com/abhilashdk2016/bookstore_items_api/clients/elasticsearch"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr: "localhost:8082",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
