package main

import "fmt"

func main() {
	//map é outra estrutuda de dados na qual podemos gauraddr dados. Ela é chave valor assim como o struct porem ela é mais rigid já que a chave e o valor tem tipos padronizados. 

	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"]) //jeito de acessar o valor

	usuario["nome"] = "Teste"

	fmt.Println(usuario["nome"]) //Teste

	//map de map

	infos := map[string]map[string]string{
		"nomes": {
			"nome":  "pedro",
			"nome2": "luiz",
		},
		"cidades": {
			"cidade1": "Diamantina",
		},
		
	}

	fmt.Println(infos)
	fmt.Println(infos["cidades"]["cidade1"])

	delete(infos, "nomes")

	fmt.Println(infos)
}