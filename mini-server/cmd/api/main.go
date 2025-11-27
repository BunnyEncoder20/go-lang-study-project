package main

import (
	"fmt"
	"net/http"

	"github.com/BunnyEncoder20/go-lang-study-project/mini-server/internal/handlers"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true) // enable the logger

	// this returns a pointer to a chi.Mux type (just a struct which we use to setup our API)
	var r *chi.Mux = chi.NewRouter()

	// passing the router to the handlers:
	handlers.Handler(r)

	// starting up the server using the http package
	fmt.Println("Starting the Go API Mini Server...")
	fmt.Println(`
  ________           _________                                
 /  _____/  ____    /   _____/ ______________  __ ___________ 
/   \  ___ /  _ \   \_____  \_/ __ \_  __ \  \/ // __ \_  __ \
\    \_\  (  <_> )  /        \  ___/|  | \/\   /\  ___/|  | \/
 \______  /\____/  /_______  /\___  >__|    \_/  \___  >__|   
        \/                 \/     \/                 \/       
	`)

	// Listening and servering at port 8000
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
