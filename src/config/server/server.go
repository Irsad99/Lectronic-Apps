package server

import (
	"log"
	"net/http"
	"os"

	"github.com/Irsad99/LectronicApp/src/routers"

	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	// mainRoute.Use(mainRoute)

	if mainRoute, err := routers.New(); err == nil {
		var addrs string = ""

		handler := c.Handler(mainRoute)

		if pr := os.Getenv("APP_PORT"); pr != "" {
			addrs = ":" + pr
		}

		log.Println("App running on 127.0.0.1" + addrs)

		if err := http.ListenAndServe(addrs, handler); err != nil {
			return err
		}

		return nil

	} else {
		return err
	}
}
