package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint  `json:"idade"`
}

func main(){
	//json.Marshal() // converte map pra json ou struct pra json
	//json.Unmarshal()//transforma json em sruct ou map dependendo da necessidadade

	c := cachorro{"Rex", "Dalmata", 3}
	cJson, error := json.Marshal(c)

	if(error != nil){
		log.Fatal(error)
	}

	fmt.Println(cJson) // vai logar um slice de byte
	fmt.Println(bytes.NewBuffer(cJson)) //com o uso da lib bytes consigo vizualizar o json {"nome":"Rex","raca":"Dalmata","idade":3}
}