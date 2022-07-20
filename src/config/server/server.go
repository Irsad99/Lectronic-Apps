package server

import (
	"log"
	"net/http"
	"os"

	"github.com/Irsad99/LectronicApp/src/routers"

	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "server",
	Short: "start api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {

	// headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	// origins := handlers.AllowedOrigins([]string{"http://localhost:8080/"})

	// mainRoute.Use(mainRoute)

	if mainRoute, err := routers.New(); err == nil {
		var addrs string = ""

		if pr := os.Getenv("APP_PORT"); pr != "" {
			addrs = "127.0.0.1:" + pr
		}

		log.Println("App running on " + addrs)

		if err := http.ListenAndServe(addrs, mainRoute); err != nil {
			return err
		}

		return nil

	} else {
		return err
	}
}
