package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rupakveerla/go-booking/pkg/config"
	"github.com/rupakveerla/go-booking/pkg/render"
)

const portNumber = ":4000"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// change this to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Unable to create template cache!")
	}

	app.TemplateCache = tc
	app.UseCache = true

	render.NewTemplates(&app)

	fmt.Printf("Appplication started on port %v\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
