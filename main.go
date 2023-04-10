package main

import "fmt"

func main() {
	//Declaración de constantes
	const pi float64 = 3.1416 //se agrega tipo de variable
	const pi2 = 3.1416        // no se le agrega
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración variables enteras
	//Si no se ha usado la variable antes se le puede colocar := los cuales indican que la variable no ha sido declarada anteriormente, por lo cual creara la variable desde cero. Se debe colocar el valor ¿

	base := 12
	//otra forma de crearla es colocando var
	//no se puede compilar si hay variables que no se han usado
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//Zero values es un valor que tendrá una variable que no se ha inicializado.
	var a int     //0
	var b float64 //0
	var c string  //vacio
	var d bool    //false

	fmt.Println(a, b, c, d)

	//Area del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del Cuadrado: ", areaCuadrado)

	fmt.Println("Hello world")

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplicación

	result = x * y
	fmt.Println("Multiplicación: ", result)

	//División

	result = y / x
	fmt.Println("División: ", result)

	//Modulo o residuo de la División. útil para saber si es par o no

	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental

	x++
	fmt.Println("Incremental de X: ", x)

	//Decremental

	x--
	fmt.Println("Decremental de X: ", x)

	//Áreas

	//Rectángulos

	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El área del rectangulo es: ", areaRectangulo)

	//Círculo

	const pinew float64 = 3.1416
	var radioCirculo float64 = 10

	areaCirculo := pinew * radioCirculo * radioCirculo

	fmt.Println("El área del círculo  es: ", areaCirculo)

	//Trapecio

	var baseUno float64 = 20
	var baseDos float64 = 18
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno * baseDos) * alturaTrapecio) / 2

	fmt.Println("El área del trapecío es: ", areaTrapecio)

}
