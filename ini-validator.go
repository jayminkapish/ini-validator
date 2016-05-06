package main

import (
	"log"

	"gopkg.in/gcfg.v1"
)

func main() {
	var cfg Config

	err := gcfg.ReadStringInto(&cfg, `# test config ini
	[global]
	    A = Test
	    B = #empty string
	    C = false
`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(cfg.Global.A, cfg.Global.B, cfg.Global.C)
	}
}

type Config struct {
	Global struct {
		A string
		B string
		C bool
	}
}
