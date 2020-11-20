package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	gomail "gopkg.in/gomail.v2"
)

var tmpl = template.Must(template.New("somesuch").Funcs(template.FuncMap{
	"IsEven": func(i int) bool {
		return i%2 == 0
	},
}).ParseGlob("view/HTML/*"))

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

func CakeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "cake", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NotActuallyApplyHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "notactuallyapply", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ApplyHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "apply", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SendApplicationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	recaptchaURL := "https://www.google.com/recaptcha/api/siteverify"
	recaptchaReq := url.Values{}
	recaptchaReq.Set("secret", os.Getenv("RECAPTCHA_SECRET_KEY"))
	recaptchaReq.Set("response", r.FormValue("g-recaptcha-response"))

	transport := http.DefaultTransport
	req, r_err := http.NewRequest("POST", recaptchaURL, bytes.NewBufferString(recaptchaReq.Encode()))
	if r_err != nil {
		http.Error(w, r_err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
		http.Error(w, fmt.Sprint(recaptcha_resp_map), http.StatusInternalServerError)
		return
	}

	//to := "caligiure.ja@gmail.com"
	to := "vaillantfeubeaucerons@hotmail.com"

	emailBody := "Name: " + r.FormValue("Name") + "<br>" +
		"Age: " + r.FormValue("Age") + "<br>" +
		"Phone: " + r.FormValue("Phone") + "<br>" +
		"Email: " + r.FormValue("Email") + "<br>" +
		"Address: " + r.FormValue("Address") + "<br>" +
		"Rent Or Own: " + r.FormValue("RentOrOwn") + "<br>" +
		"If renting, do you have your landlord's permission to own a puppy of this size?: " + r.FormValue("IfRenting") + "<br>" +
		"How many members of your household are there, and how old are they: " + r.FormValue("Household") + "<br>" +
		"Why do you want a beauceron and what is your experience with the breed: " + r.FormValue("WantExp") + "<br>" +
		"What dogs/animals have you owned in the past? How long did you have them and what happened to them if they are no longer in your care?: " + r.FormValue("PastAnimals") + "<br>" +
		"Vet References: " + r.FormValue("VetReference") + "<br>" +
		"Trainer References:  " + r.FormValue("TrainerReference") + "<br>" +
		"What is your dog training experience: " + r.FormValue("DogExp") + "<br>" +
		"What titles and training have you done with your past/current dogs: " + r.FormValue("TitlesTraining") + "<br>" +
		"What are you looking for in a puppy: " + r.FormValue("WhatLookingFor") + "<br>" +
		"What is your activity level during an average week: " + r.FormValue("ActivityLevel") + "<br>" +
		"How much time can you realistically devote to training your new beauceron puppy: " + r.FormValue("RealisticallyTraining") + "<br>" +
		"What activities do you plan to use to stimulate your beauceron puppy mentally: " + r.FormValue("MentalStimulation") + "<br>" +
		"What activities do you plan to use to stimulate your beauceron puppy physically: " + r.FormValue("PhysicalStimulation") + "<br>" +
		"What sport or work do you want a beauceron for: " + r.FormValue("WhatSport") + "<br>" +
		"Do you have a dedicated club to assist you: " + r.FormValue("Club") + "<br>" +
		"What level are you hoping to achieve in your sport (if applicable): " + r.FormValue("SportLevel") + "<br>" +
		"What are you looking for in terms of drives and energy level: " + r.FormValue("DriveEnergy") + "<br>" +
		"What venues do you plan to show in (AKC/UKC/IABCA): " + r.FormValue("Venues") + "<br>" +
		"Have you ever titled a dog to Championship or GCH before: " + r.FormValue("Titled") + "<br>" +
		"Did they say they don't abuse animals?: " + r.FormValue("NotAboose") + "<br>"

	msg := gomail.NewMessage()
	msg.SetHeader("From", "New Applicant <applications@beaucerons-163616.appspotmail.com>")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "New Puppy Application!")
	msg.SetBody("text/html", emailBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, "caligiure.ja", "vlvptiwwqgfkujbt")
	if err := d.DialAndSend(msg); err != nil {
		http.Error(w, r_err.Error(), http.StatusInternalServerError)
		return
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

func FAQHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "faq", [][]string{
		{
			"Important Links",
			"American Beauceron Club",
			"http://www.beauce.org/",
			"Les Amis du Beauceron (French Parent Club)",
			"http://www.amisdubeauceron.org/",
			"Beauceron Breed Standard",
			"http://www.beauce.org/docs/Beauceron%20Breed%20Standard.pdf",
			"So You Think You Want a Beauceron? (ABC Article)",
			"http://www.beauce.org/docs/So%20You%20Think%20You%20Want%202009.pdf",
			"Orthopedic Foundation for Animals (OFA)",
			"https://www.ofa.org",
			"Leerburg (Training Video Library, Products, Classes, and More)",
			"http://www.leerburg.com",
			"Puppy Culture",
			"https://shoppuppyculture.com/",
			"What Does It Mean to Be A Preservation Breeder? (Showsight Magazine)",
			"https://www.showsightmagazine.com/article/what-does-it-mean-to-be-a-preservation-breeder/124",
		},
		{
			"What is a purposeful preservation breeder? Why do you consider yourself one?",
			"A preservation breeder is an individual who is breeding two purebred animals together in order to preserve a certain type of animal. To go further, a purposeful preservation breeder definition according to PWC legend Bill Shelton: “A purposeful preservation breeder is someone who is going to do as little harm as possible and move the breed in the right direction”. For me personally, this means ensuring I critically evaluate my animals as they grow which includes not only seeking opinions from conformation experts in the show ring but also putting them under various pressures, different environments, and in novel situations to evaluate their responses. I work my animals because it is very difficult to truly know what you have if a breeder does not go beyond a show ring and a couch. I do not breed often in order to wait out my progeny and see what I have accomplished for each breeding, thus minimizing any harm I may do by overbreeding a problematic bitch/stud.",
		},
		{
			"Do you title your beaucerons?",
			"Yes. All of my beaucerons are titled in a discipline of work/sport (usually herding as it is the original purpose of the breed) as well as a conformation title or evaluation from a French Judge. I feel this proves a balanced animal.",
		},
		{
			"Do you health test your beaucerons?",
			"Yes. All of my breeding stock has public results listed on the Orthopedic Foundation for Animals (OFA) site regardless of abnormal or normal results; there are no hidden health secrets here at Vaillant Feu. I encourage all of my puppy buyers, even those with pets, to complete the American Beauceron Club’s health testing requirements as it adds to our knowledge of the breed’s health as a whole and helps add to my knowledge of what kind of health I am achieving in my breedings across the board, not just in the animals I keep.",
		},
		{
			"Do you sell breeding rights?",
			"I do not sell puppies for breeding, as you cannot promise a puppy will turn out especially in a slow maturing breed no matter how nice they seem at 8 weeks. Breeding rights can be released later at no extra cost after verification of clear health testing results and titling has been achieved for those with full registration.",
		},
		{
			"Do you sell protection dogs?",
			"If you are seeking a beauceron to sit on a world podium with you or do civil work for the police, I do not have the dogs for you but I can point you in the right direction. If you are seeking a beauceron to play with in a protection sport because you love the breed first and the sport second, I may have the right dog for you or can point you in a solid direction of a better pairing at the moment. I will not sell puppies to “protect your spouse and kids while you’re away but will go on lots of walks”; unless you have a dedicated trainer/club lined up or are a protection trainer yourself that I can & will verify, you will have to seek another breeder to sell you one. ",
		},
		{
			"How do they compare to malinois?",
			"A beauceron will never be what a malinois is; please don’t go looking for a malinois in a beauceron body or you will be disappointed almost every time. Their energy levels can be similar (some beaucerons can be more even), but their drives are not equal. You will find more food driven beaucerons that enjoy a toy/ball to some extent rather than toy crazed monsters. Their grips are usually like a typewriter on the bite as well, rather chewy and lots of movement/regrips; so most that come from malinois will find the continuous grip building frustrating.",
		},
		{
			"I want a harlequin. Will you sell me one?",
			"I do not breed for color, I breed with the best pairing I can find and preserving my breed in mind first and foremost. I may not be breeding a harlequin pairing that year, for example. I do take color preferences into consideration but it is not a base factor in the selection process for which puppies go where.",
		},
		{
			"Do you pick the puppy for me or can I choose it? Can I send someone to evaluate the puppy myself?",
			"I choose the right puppy for you after interviewing you initially after applying and getting to know you over the course of however long until your puppy is born and ready to go at 8 weeks old. I raise my litters with Puppy Culture and at the 8 week old point, I know so much not only about the parents who made your puppy but that individual & the culmination of seeing it grow, react, explore, and develop a personality. I have objective 3rd parties come evaluate the puppies as well but these are individuals I know and trust in their field of work. With the environment we have going on with the AR movement in my area, I will not allow any evaluators I do not know personally into my home.",
		},
		{
			"Can I meet the parents of my puppy?",
			"Absolutely, provided that I own both the male and female. In order to do what is best for the breed, I will choose to go overseas on occasion for a sire and that dog will not be available to be met by interested parties obviously. I usually invite interested parties to meet my dogs at events or training in a public area at first, as there are a large number of Animal Rights activists in this area and it is safer for everyone to meet in public first.",
		},
		{
			"How do I get pick of the litter or on the waiting list?",
			"There is no “pick of the litter” with me, everyone will receive the best puppy for your home and purpose. Contacting me and expressing interest in receiving a puppy from one of my pairings is the best way to get your foot in the door as I do not take waiting lists or names until I have a concrete breeding plan in place.",
		},
		{
			"I need a service dog. Do you sell service dog prospects?",
			"I do not sell service dog prospects because I cannot guarantee that type of unique temperament that is required of a beauceron as a service dog. Beaucerons are an incredibly physical breed that could very easily without malicious intent hurt a person with disabilities and their slow maturity rate means you may wait 3+ years before your dog is considered fully trained. It is a hard breed to turn into a service animal and most who utilize beaucerons as service animals are professional trainers themselves.",
		},
		{
			"Do you give discounts?",
			"No, I do not.",
		},
		{
			"Do you take deposits?",
			"I take deposits from interested approved parties after the birth of the puppies, occasionally when pregnancy is confirmed. A deposit shows your serious commitment to receiving a puppy from that pairing.",
		},
		{
			"Do you do payment plans?",
			"No, I do not do any form of payment plans at the moment.",
		},
		{
			"What are your expectations of a puppy owner?",
			"At the base level, I expect a puppy owner to love & care for their puppy for the duration of its life. Care is including but not limited to lifelong socialization & training, daily mental & physical stimulation, and yearly veterinary visits. Beyond that, expectations will vary depending on the purpose for which you bought your puppy. For example: if you asked me for a dog to title in agility, I will expect to see you training in agility with your puppy. I will never pressure or force you to pursue an avenue of training, I am only here to support and encourage the goals you explained to me when you first inquired about purchasing the puppy.",
		},
		{
			"What can I expect from you as a breeder?",
			"Lifelong support and encouragement, for starters. Through the bad and good, I will be there to support in any way that I can. I will celebrate your successes and give encouragement through the failures & missteps. You will get the expertise I have gained over my years in the dog world training and trialing, as well as the backup expertise of my breed mentors (if I do not know the answer myself, I can find it for you). I am an Army veteran, therefore I value honest and direct communication always. I will be as involved as you want me to be with your puppy and your journey together; some people talk to their breeders every week, some prefer to check in every month or so. I do actively train and enter trials with my dog as well as enter the conformation ring occasionally so as a breeder, I have a lot of guidance and help I can give you as a buyer who may be starting out for the first time or the first time with this breed.",
		},
		{
			"What happens if my puppy does not work out?",
			"I will always take any dog I have produced back for any reason, at any point in its life. In my contract it states the dog cannot be surrendered to any rescue or shelter nor rehomed or given to any individual who is not the person on the original contract.",
		},
		{
			"Do you guarantee health?",
			"Yes, my puppies are guaranteed for dysplasia and other genetic health issues that severely impact the quality of the animal’s life.",
		},
		{
			"Do you guarantee drive/work/success in a certain job/sport?",
			"Absolutely not. Puppies can be “liars” and a “crapshoot”, which means that a puppy who seems wild & in charge at 3 months old may become a reserved less intense adult after maturity. While I try to breed and place appropriately, I can only tell you what I see in that moment as far as potential and the rest is up to not only you as the handler but how the dog ultimately develops. I disclose everything my early testing projects, but drive can evolve due to many factors.",
		},
		{
			"Do your dogs live in your home?",
			"All of the Vaillant Feu dogs live inside of my home with me and are members of my family. While we have an outdoor kennel run, none of the dogs live in it.",
		},
		{
			"Do you do any ENS or other puppy rearing programs?",
			"Yes, I do Early Neurological Stimulation as well as the Puppy Culture program. I provide my buyers with photos and videos of this program throughout their growth.",
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TrainingHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "training", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RLitter2020Handler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "RLitter2020", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
