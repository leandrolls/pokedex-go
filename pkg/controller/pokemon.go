package controller

import (
	"PokeDexAPI/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func ConsumeAPI() {
	resp, err := http.Get("http://pokeapi.co/api/v2/type/1")
	if err != nil {
		ReponseError()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ReponseError()
	}

	content := &model.Api{}

	err = json.Unmarshal(body, &content)
	if err != nil {
		ReponseError()
	}

	fmt.Println("Golpes:")
	for _, c := range content.Moves {
		fmt.Printf(c.Name)
		fmt.Println("")

	}
	fmt.Println("")

	// fmt.Println("Pokemons:")
	// implementar função pra buscar pokemons do tipo normal


}