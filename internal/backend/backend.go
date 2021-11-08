package backend

import (
	"github.com/milessh/polling-app/internal/handlers"
	"github.com/milessh/polling-app/pkg/server"
)

func Start() {
	// initialise database connection
	// err := db.InitDb()
	// if err != nil {
	// 	log.Fatalln("Failed to initialise, terminating")
	// }

	// initialise http server
	server.Init()

	// initialise http handlers
	handlers.SetupHandlers()

	// start http server
	server.Start()
}

func Terminate() {
	// stop accepting new connections and finish ongoing requests before shutting down http server
	server.Terminate()

	// close database connection
	// log.Println("Closing connection to database")
	// repo.Close()
}
