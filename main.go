package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Noofbiz/duvaillantfeu/controller"
)

func main() {
	http.HandleFunc("/", controller.MainHandler)
	http.HandleFunc("/BreedStandard", controller.BreedStandardHandler)
	http.HandleFunc("/Litters/Upcoming", controller.UpcomingLittersHandler)
	http.HandleFunc("/Litters/OLitter2018", controller.OLitter2018Handler)
	http.HandleFunc("/Litters/Care", controller.LittersCareHandler)
	http.HandleFunc("/Parents/Jamais", controller.JamaisHandler)
	http.HandleFunc("/Parents/Mowgli", controller.MowgliHandler)
	http.HandleFunc("/Parents/Oshi", controller.OshiHandler)
	http.HandleFunc("/Apply", controller.ApplyHandler)
	http.HandleFunc("/ActuallyApply", controller.ActualApplyHandler)
	http.HandleFunc("/Apply/Send", controller.SendApplicationHandler)
	http.HandleFunc("/Apply/Success", controller.SuccessfulApplicationHandler)
	http.HandleFunc("/Contact", controller.ContactHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
