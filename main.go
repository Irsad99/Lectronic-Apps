package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Irsad99/LectronicApp/src/routers"
)

func main() {
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("aplikasi berjalan pada port 8080")

	if err := http.ListenAndServe(":8080", mainRoute); err != nil {
		log.Fatal("aplikasi gagal dijalankan")
	}
}
