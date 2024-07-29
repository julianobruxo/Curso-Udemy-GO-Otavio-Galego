package main

import (
	"fmt"
	"linha-de-comando/app" //nome do go.mod criado/nome da aplicação
	"log"
	"os"
	"time"
)

func boasVindas() string {
	fmt.Println("Seja bem-vindo!\n Digite o site que deseja pesquisar sem wwww:")
	time.Sleep(time.Second * 2)
	var site string
	fmt.Scanln(&site)
	return site
}

func main() {

	for {
		site := boasVindas()
		fmt.Println("Listando os IPs do site especificado:")

		os.Args = append(os.Args, "ip", "--host", site)

		aplicacao := app.Gerar() // a var aplicação recebe a Aplicação

		erro := aplicacao.Run(os.Args) //Run(os.Args) é o parâmetro padrão usado para que os comandos do SO
		// sejam reconhecidos pela aplicação de linha de comando
		// Run sempre retorna um erro, então é necessário tratar o erro:
		if erro != nil {
			log.Fatal(erro) //Fatal is equivalent to [Print] followed by a call to os.Exit(1) que encerra o programa
		}

		os.Args = append(os.Args[:len(os.Args)-2], "servidores", "--host", site)
		erro = aplicacao.Run(os.Args)
		if erro != nil {
			log.Fatal(erro)
		}

		time.Sleep(time.Second * 1)
		fmt.Print("\nDeseja buscar por outro site? (sim/não): ")
		var resposta string
		fmt.Scanln(&resposta)
		if resposta != "sim" && resposta != "Sim" {
			fmt.Println("Obrigado por usar nosso programa!")
			time.Sleep(2 * time.Second) // Adiciona um atraso de 2 segundos antes de reiniciar a busca
			break
		}
	}
}
