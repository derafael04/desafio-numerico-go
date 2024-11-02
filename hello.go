// Projeto conversão de Kelvin para Celsius
package main

import "fmt"

const tempEbulicaoAguaK = 373.0

func main() {

	var tempKelvin = tempEbulicaoAguaK
	var tempCelsius = tempKelvin - 273

	fmt.Printf("Temperatura de ebulição em Kelvin: %g e a temperatura de ebulição em Celsius: %g", tempKelvin, tempCelsius)

}
