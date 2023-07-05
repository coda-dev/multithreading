package main

import (
	"fmt"
	"time"

	"github.com/coda-dev/multithreading/api"
)

func main() {

	viaCEP := make(chan *api.ViaCEP)
	cdnCEP := make(chan *api.CdnCEP)

	// ViaCEP
	go func() {
		for {
			cep, _ := api.BuscaViaCep("78897-066")
			viaCEP <- cep
		}
	}()

	// CdnCEP
	go func() {
		for {
			cep2, _ := api.BuscaCdnCep("90880-440")
			cdnCEP <- cep2
		}
	}()

	select {
	case msg := <-viaCEP:
		if msg.Logradouro != "" {
			fmt.Printf("ViaCEP: %s\n", msg.Logradouro+" "+msg.Cep)
		}
	case msg := <-cdnCEP:
		if msg.Logradouro != "" {
			fmt.Printf("CdnCEP: %s\n", msg.Logradouro+" "+msg.Cep)
		}

	case <-time.After(time.Second):
		println("timeout")
	}

}
