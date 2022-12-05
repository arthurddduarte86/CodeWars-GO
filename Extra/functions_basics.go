package main

func main() {
	result := soma(11, 20)
	println("Abaixo é a função soma")
	println(result)
	println("Abaixo a função somaTudo")
	println(somaTudo(1, 2, 3, 4, 5))
	//
	//
	resultadO := func(x ...int) func() int {
		res := 0
		for _, val := range x {
			res += val
		}
		return func() int { return res * res }
	}
	println(resultadO(1, 2, 3, 4, 5)(), " é o resultadO")

}

func soma(a, b int) (resultado int) { resultado = a + b; return }

func somaTudo(X ...int) int {
	resultado := 0
	for _, v := range X {
		resultado += v
	}
	return resultado
}
