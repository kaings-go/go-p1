package main

import (
	"fmt"
	"log"
	"net/http"
)


func (app *App) HomePage (w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
}


func (app *App) serve() {
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	app.InfoLog.Println("Starting web server... ")
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
 