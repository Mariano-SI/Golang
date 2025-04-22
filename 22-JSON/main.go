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
	

	cachorro1 := cachorro{"Rex", "Dalmata", 3}
	cachorroJson, error := json.Marshal(cachorro1)

	if(error != nil){
		log.Fatal(error)
	}

	fmt.Println(cachorroJson) // vai logar um slice de byte
	fmt.Println(bytes.NewBuffer(cachorroJson)) //com o uso da lib bytes consigo vizualizar o json {"nome":"Rex","raca":"Dalmata","idade":3}


	//json.Unmarshal()//transforma json em sruct ou map dependendo da necessidadade

	var cachorro2 cachorro //declara a struct em branco
	if erro := json.Unmarshal(cachorroJson, &cachorro2); erro != nil{ //recebe um slice de byte que seria o json e um endereço de memoria onde vai colocar a conversao
		log.Fatal(erro)
	} 

	fmt.Println(cachorro2)

}