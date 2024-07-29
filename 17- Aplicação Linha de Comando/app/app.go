package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App { // ponteiro que aponta para a aplicação CLI
	//cli.App é um método que vai retornar a função NewApp que precisamos
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"              // configura o nome da Aplicação
	app.Usage = "Busca Ips e nomes de Servidor na Internet" // configura o propósito da Aplicação

	flags := []cli.Flag{ // flags são sempre precedidas de --
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br", //nome padrãopassado caso não seja especificado valor para não retornar erro
		},
	}

	app.Commands = []cli.Command{
		{ //commando é um struct q lista os camandos que serão utilizados em um slice
			Name:   "ip", // lista o nome do comando
			Usage:  "Busca endereços de IPs na Net",
			Flags:  flags, //são os parâmentros que passamos para a CLI
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "buscar o nome dos servidores na net",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}
func buscarIps(c *cli.Context) {
	host := c.String("host") // var host recebeu o parametro cli.Context  o valor de --host(amazon.com)

	ips, erro := net.LookupIP(host) // vai retornar um slice de ips e um erro
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	//hora de chamar uma função do pacote net para retornar os servidores
	servidores, erro := net.LookupNS(host) //NS=name server
	if erro != nil {                       // se der erro, tratamos ele com log.Fatal
		log.Fatal(erro)
		//não havendo erro, iteramos pelo slice retornado e listamos eles
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
	fmt.Println("Obrigado por usar!")
}
