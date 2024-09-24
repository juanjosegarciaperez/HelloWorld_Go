package main

import (
	"fmt"
)

func pensionados() {
	var (
		nombre     string = "jaime"
		edad       int    = 18
		pensionado bool   = false
	)

	fmt.Println("Nombre ", nombre)
	fmt.Println("Edad ", edad)
	fmt.Println("Pensionado ", pensionado)
}

func valor_cero () {
	var ( 
		nombre string
		edad int 
		peso float64
		estudiante bool 
	)
	fmt.Println("Nombre ", nombre)
	fmt.Println("Edad ", edad)
	fmt.Println("Peso ", peso)
	fmt.Println("Estudiante ", estudiante)
}
// no se puede asignar un valor directamente a las variables que estén dentro de un bloque var()
//no se puede usar := en variables globales
var materia = "español"
func Declaracion_corta_variables() {

	
	nombre := "Pablo"
	edad := 19
	peso := 70
	estudiante := false
	
	fmt.Println("nombre", nombre)
	fmt.Println("Edad", edad)
	fmt.Println("peso", peso)
	fmt.Println("Estudiante", estudiante)
	fmt.Println("Materia", materia)
}

var var1 = "este es el primer nivel"
func scoope () {
	var var2 = "este es el segundo nivel"
	{
		var var3 = "este es el tercer nivel"
		fmt.Println("var3:", var3)
	}
	fmt.Println("var1:", var1)
	fmt.Println("var2:", var2)
}

func puntero () {
	color := "rojo"
	fmt.Println(color, &color)
}

func referencia_Copia () {
	vehiculo1 := "rojo"
	fmt.Println("el vehiculo1 es ", vehiculo1)
	vehiculo2 := vehiculo1
	fmt.Println("el vehiculo2 es", vehiculo2)
	vehiculo3 := &vehiculo1
	fmt.Println("el vehiculo3 es", *vehiculo3)

	vehiculo1 = "gris"

	fmt.Println("el vehiculo1 es", vehiculo1, &vehiculo1)
	fmt.Println("el vehiculo2 es", vehiculo2, &vehiculo2)
	fmt.Println("el vehiculo3 es", *vehiculo3, &vehiculo3)

}

func main() {
	// fmt.Println("texto de la funcion del ejecicio 2")
	// pensionados()
	// valor_cero()
	// Declaracion_corta_variables()
	// scoope()
	// puntero()
	// referencia_Copia()
}
