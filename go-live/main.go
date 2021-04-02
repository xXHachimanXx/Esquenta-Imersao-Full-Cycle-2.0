package main

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xXHachimanXx/Esquenta-Imersao-Full-Cycle-2.0/http"
	"github.com/xXHachimanXx/Esquenta-Imersao-Full-Cycle-2.0/model"
)

func main() {
	p1 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Carrinho",
	}

	p2 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Boneca",
	}

	products := model.Products{}
	products.Add(p1)
	products.Add(p2)

	server := http.NewWebServer()
	server.Products = &products
	server.Serve()
}

// type Pessoa struct {
// 	Nome  string
// 	Idade int
// }

// func (p *Pessoa) setNome(nome string) {
// 	p.Nome = nome
// }

// // Funcao do struct Pessoa
// func (p Pessoa) andou() (string, error) {
// 	if p.Nome != "Felipe" {
// 		return "", errors.New("Nome tem que ser Felipe")
// 	}

// 	return p.Nome + " andou", nil
// }

// // Funcao publica
// func andou(pessoa Pessoa) (string, error) {
// 	if pessoa.Nome != "Felipe" {
// 		return "", errors.New("Nome tem que ser Felipe")
// 	}

// 	return pessoa.Nome + " andou", nil
// }

// func main() {
// 	//var nome string
// 	// var nome Nome

// 	// nome = append(nome, "Felipe")
// 	// nome = append(nome, "Andrade")

// 	// for _, v := range nome { // _ = Blank Idetifier
// 	// 	fmt.Println(v)
// 	// }

// 	p := Pessoa{
// 		Nome:  "Felipe",
// 		Idade: 21,
// 	}

// 	var p2 *Pessoa
// 	p2 = &p

// 	res, err := (*p2).andou()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	fmt.Println(res)
// }
