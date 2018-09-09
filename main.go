package beaucerons

import (
	"net/http"

	"laurenSite/controller"
)

func init() {
	http.HandleFunc("/", controller.MainHandler)
	http.HandleFunc("/BreedStandard", controller.BreedStandardHandler)
	http.HandleFunc("/Litters/Upcoming", controller.UpcomingLittersHandler)
	http.HandleFunc("/Litters/Care", controller.LittersCareHandler)
	http.HandleFunc("/Parents/Jamais", controller.JamaisHandler)
	http.HandleFunc("/Parents/Mowgli", controller.MowgliHandler)
	http.HandleFunc("/Apply", controller.ApplyHandler)
	http.HandleFunc("/Apply/Send", controller.SendApplicationHandler)
	http.HandleFunc("/Apply/Success", controller.SuccessfulApplicationHandler)
	http.HandleFunc("/Contact", controller.ContactHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static"))))
}
