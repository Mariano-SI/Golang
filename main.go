package main

func main() {
	// Declaração com var
	// Quando declaramos uma variável com var nao somos obrigados a inicializar ela, mas podemos fazer isso. Se não lhe dermos um valor durante a incialização, devemos declarar o tipo da variável.
	//Variaveis declaradas podem ser reatribuidas, ou seja, podemos mudar o valor dela depois de inicializada.

	//inicializando a variável com atribuição
	var a = 10;
	println(a);

	//inicializando a variável sem atribuição
	var b bool;
	b = true;
	println(b);

	//Declaração com const
	// Constantes são variáveis que não podem ser reatribuídas, ou seja, não podemos mudar o valor dela depois de inicializada. sempre que declaramos uma const ela precisa ser inicializada com um valor, e não podemos mudar o valor dela depois de inicializada.
	// Sobre tipagem: se não declaramos o tipo da variável, o Go vai inferir o tipo dela baseado no valor que estamos atribuindo a ela. Se declaramos o tipo da variável, o Go não vai inferir o tipo dela, e vamos ter que atribuir um valor do mesmo tipo que declaramos.

	//declaração sem tipagem
	const numero = 10;
	//numero = 20; // Isso vai dar erro, pois estamos tentando reatribuir uma constante.

	//declaração com tipagem
	const numero2 int = 10;


	//declaração com :=
	// O := é uma forma de declarar e inicializar uma variável ao mesmo tempo, e o Go vai inferir o tipo dela baseado no valor que estamos atribuindo a ela. Não podemos usar o := para declarar variáveis globais, apenas variáveis locais.
	//variaveis declaradas com := podem ser reatribuídas, ou seja, podemos mudar o valor dela depois de inicializada.

	numero3 := 10; // O Go vai inferir que o tipo da variável é int, pois estamos atribuindo um valor inteiro a ela.
	numero3 = 20; // Isso é permitido, pois estamos reatribuindo uma variável, e não uma constante.

	println(numero3); // Isso vai imprimir 20, pois estamos reatribuindo o valor da variável numero3.

}
