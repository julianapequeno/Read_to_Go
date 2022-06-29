# Hello, Go! :)

<div align="center">
    <image src="./images/golang.png" width=400 heigth=400>


Seja bem vinda mais nova linguagem de programação deste github!! (Não apenas dele né).
</div>

Decidi começar a ver um pouco da linguagem de programação da google, a Golang. Assim, vou compartilhar aqui alguns códigos que estaria fazendo para aprender. 

Inicialmente, __conceitos básicos__ para não esquecer:
- Go é uma linguagem compilada
- A biblioteca "fmt" é a que contém o println, ela é a _biblioteca padrão_ da linguagem Go
- As funções _build_in_ da linguagem iniciam com letra maiúscula (que diferenteee).
- Ele deixa o ; ser __opcional__ (😱)
- As {} devem ficar ao lado da função, e não embaixo. Essa convenção é definida.
- Go é uma linguagem para simplificar, para ser rápida e direta.
- A linguagem possui várias convenções que ajudam o desenvolvedor a focar no problema a ser resolvido.

__E VAMOS DE CURIOSIDADE E MAIS CONVENÇÕES__
 Se você não colocar um valor nas variáveis do go, elas sempre serão vazias. Até na string, a qual ele considera como um caractere vazio.
 __Convenção__: não deixa você declarar uma variável e NÃO A UTILIZAR. Dá um _erro de compilação_👀

 - Não existe float! Apenas float32 e float64.

O Go consegue inferir o tipo das variáveis! Obs: Para float, é melhor deixar, para ter certeza de que será float32 ou float64.

__EXECUTANDO OS ALGORITMOS__
##### Exemplo feito com o algoritmo do hello.go
1. Compilando e executando de forma separada

```go
    go build hello.go
    ./hello 
```

2. Compilando e executando em um só comando
```go
    go run hello.go
```

