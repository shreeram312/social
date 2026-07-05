package main

import (
	"log"
	"github.com/shreeram/social/internal/env"
)



func main(){
	if err := env.Load(); err != nil {
		log.Println("no .env file found, using defaults")
	}

	cfg:=config{
		addr: env.GetString("ADDR",":8080"),
	}


	app:= &application{
		config:cfg,
	}

	mux:=app.mount()
	log.Fatal(app.run(mux))



}