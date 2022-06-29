package main /*package define o pacote, e o nome main, define que esse arquivo é um programa */

import (
	"fmt" //Biblioteca build_in, para o Println
)

func main() {
	name := "Juliana" //* declara a variável e ao mesmo tempo atribui-se um valor a ela
	var idade = 19
	var versao float32 = 1.1
	/*Perceba os três tipos de declaração acima*/
	fmt.Println("Hello, sra.", name, "! Sua idade é", idade)
	fmt.Println(">> Este programa está na versão", versao)

	//fmt.Println("O tipo da variável nome é: ", reflect.TypeOf((name)))

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando) //Outra opção, esta daqui, pensa da seguinte forma: "Se o tipo da variável é um int, então vou receber um int!!", e não precisa especificar isso com o %d
	fmt.Println("O endereço da minha variável comando é ", &comando)
	fmt.Println("O comando escolhido foi: ", comando)
}
