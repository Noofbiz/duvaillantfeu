package controller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
	"google.golang.org/appengine/urlfetch"
)

var tmpl = template.Must(template.ParseGlob("view/HTML/*"))

func MainHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BreedStandardHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "BreedStandard", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpcomingLittersHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "upcoming", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func OLitter2018Handler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "OLitter2018", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LittersCareHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "care", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func JamaisHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "jamais", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func MowgliHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "mowgli", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func OshiHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "oshi", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ApplyHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "notactuallyapply", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ActualApplyHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "apply", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SendApplicationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx := appengine.NewContext(r)

	recaptchaURL := "https://www.google.com/recaptcha/api/siteverify"
	recaptchaReq := url.Values{}
	recaptchaReq.Set("secret", os.Getenv("RECAPTCHA_SECRET_KEY"))
	recaptchaReq.Set("response", r.FormValue("g-recaptcha-response"))

	transport := urlfetch.Transport{Context: ctx, AllowInvalidServerCertificate: false}
	req, r_err := http.NewRequest("POST", recaptchaURL, bytes.NewBufferString(recaptchaReq.Encode()))
	recaptcha_resp, r_err := transport.RoundTrip(req)
	if r_err != nil {
		http.Error(w, r_err.Error(), http.StatusInternalServerError)
		return
	}
	defer recaptcha_resp.Body.Close()
	recaptcha_resp_map := make(map[string]interface{})
	recaptcha_resp_body, r_err := ioutil.ReadAll(recaptcha_resp.Body)
	if r_err != nil {
		http.Error(w, r_err.Error(), http.StatusInternalServerError)
	}
	json.Unmarshal(recaptcha_resp_body, &recaptcha_resp_map)

	if !(recaptcha_resp_map["success"].(bool)) {
		http.Error(w, "Failed recaptcha.", http.StatusInternalServerError)
		return
	}

	//to := "caligiure.ja@gmail.com"
	to := "vaillantfeubeaucerons@hotmail.com"

	emailBody := "Name: " + r.FormValue("Name") + "\r\n" +
		"Age: " + r.FormValue("Age") + "\r\n" +
		"Phone: " + r.FormValue("Phone") + "\r\n" +
		"Email: " + r.FormValue("Email") + "\r\n" +
		"Address: " + r.FormValue("Address") + "\r\n" +
		"Rent Or Own: " + r.FormValue("RentOrOwn") + "\r\n" +
		"How many members of your household are there, and how old are they: " + r.FormValue("Household") + "\r\n" +
		"Why do you want a beauceron and what is your experience with the breed: " + r.FormValue("WantExp") + "\r\n" +
		"What dogs/animals have you owned in the past? How long did you have them and what happened to them if they are no longer in your care?: " + r.FormValue("PastAnimals") + "\r\n" +
		"Vet References: " + r.FormValue("VetReference") + "\r\n" +
		"Trainer References:  " + r.FormValue("TrainerReference") + "\r\n" +
		"What is your dog training experience: " + r.FormValue("DogExp") + "\r\n" +
		"What titles and training have you done with your past/current dogs: " + r.FormValue("TitlesTraining") + "\r\n" +
		"What are you looking for in a puppy: " + r.FormValue("WhatLookingFor") + "\r\n" +
		"What is your activity level during an average week: " + r.FormValue("ActivityLevel") + "\r\n" +
		"How much time can you realistically devote to training your new beauceron puppy: " + r.FormValue("RealisticallyTraining") + "\r\n" +
		"What activities do you plan to use to stimulate your beauceron puppy mentally: " + r.FormValue("MentalStimulation") + "\r\n" +
		"What activities do you plan to use to stimulate your beauceron puppy physically: " + r.FormValue("PhysicalStimulation") + "\r\n" +
		"What sport or work do you want a beauceron for: " + r.FormValue("WhatSport") + "\r\n" +
		"Do you have a dedicated club to assist you: " + r.FormValue("Club") + "\r\n" +
		"What level are you hoping to achieve in your sport (if applicable): " + r.FormValue("SportLevel") + "\r\n" +
		"What are you looking for in terms of drives and energy level: " + r.FormValue("DriveEnergy") + "\r\n" +
		"What venues do you plan to show in (AKC/UKC/IABCA): " + r.FormValue("Venues") + "\r\n" +
		"Have you ever titled a dog to Championship or GCH before: " + r.FormValue("Titled") + "\r\n"

	msg := &mail.Message{
		Sender:  "New Applicant <applications@beaucerons-163616.appspotmail.com>",
		To:      []string{to},
		Subject: "New Puppy Application!",
		Body:    emailBody,
	}

	if err := mail.Send(ctx, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/Apply/Success", http.StatusSeeOther)
}

func SuccessfulApplicationHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "apply_success", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "contact", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
