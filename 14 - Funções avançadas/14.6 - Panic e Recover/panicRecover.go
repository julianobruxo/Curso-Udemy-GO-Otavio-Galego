package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado com sucesso")
		// recover salvou a execução do programa pois ao calcular a média 6 e ativar o panic,
		// ele buscou a função com defer, tratou o erro e devolveu zero, ou seja, 6 <
		// ao ser menor que seis, retornou false e continuou a função alunoAprovado
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6. Saindo do Programa")
	// o que a função panic faz? Ela interrompe o fluxo normal do programa e para tudo
} // ao ser executado, ele chama todas as funções com defer. Na ausência delas, o panic mata o programa

func main() {
	fmt.Println("Panic & Recover")
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execução")
}
