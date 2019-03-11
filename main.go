// Copyright 2019 Felix Leber
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"html/template"
	"time"
)

var templates *template.Template
var deGreetings *[]Greeting
var enGreetings *[]Greeting

func main() {
	deGreetings = initDeResponses()
	enGreetings = initEnResponses()
	http.HandleFunc("/", handleEn)
	http.HandleFunc("/en", handleEn)
	http.HandleFunc("/de", handleDe)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleDe(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseFiles("templ.gohtml"))

	greetings := selectGreeting(deGreetings)

	rand.Seed(time.Now().UnixNano())

	resp := greetings[rand.Intn(len(greetings))]

	data := map[string]interface{}{
		"greets" : resp.Text,
	}
	templates.ExecuteTemplate(w, "index", data)
}

func handleEn(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseFiles("templ.gohtml"))

	greetings := selectGreeting(enGreetings)

	rand.Seed(time.Now().UnixNano())

	resp := greetings[rand.Intn(len(greetings))]

	data := map[string]interface{}{
		"greets" : resp.Text,
	}
	templates.ExecuteTemplate(w, "index", data)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}


