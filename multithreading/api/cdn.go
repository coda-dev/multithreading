package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CdnCEP struct {
	Cep        string `json:"code"`
	Logradouro string `json:"address"`
	Bairro     string `json:"district"`
	Localidade string `json:"city"`
	Uf         string `json:"state"`
	StatusText string `json:"statusText"`
}

// cep 78897-066
func BuscaCdnCep(cep string) (*CdnCEP, error) {
	resp, error := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c CdnCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}
