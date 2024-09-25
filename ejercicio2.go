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

func equivalencia_en_pies (altura float32) float32 {
	altura = altura * 3.28
	return altura
}

var altura float32 = 1.70

func ConversionEnPies (altura2 *float32) float32 {
	*altura2 = *altura2 - 0.10
	return *altura2
}

var altura2 float32 = 1.70

const Pi = 3.1416

func area (radio float64) float64 {
    return Pi * radio * radio
}

func circulo (radio float64) (area float64, perimetro float64) {
	area = Pi * radio * radio
	perimetro = 2 * Pi * radio
	return area, perimetro
}

func sumar1 (numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total = numero + total
	}
	return total
}

func ecodelamotaña (mensaje string, iteraciones uint) {
	if iteraciones > 1 {
		ecodelamotaña(mensaje, iteraciones-1)
	}
	fmt.Println(mensaje, iteraciones)
}

func circulo2 (radio float64) (area1 func() float64, perimetro1 func() float64) {
	area1 = func() float64 {
		return Pi * radio * radio
	}

	perimetro1 = func() float64 {
        return 2 * Pi * radio
    }

	return
}

func Objeto_Animal_Persona () {
	var juguete string
    fmt.Println("Elige persona, animal o cosa:")
    fmt.Scanln(&juguete)
    if juguete == "persona" {
        fmt.Println("El objeto es una persona")
    } else if juguete == "cosa" {
        fmt.Println("El objeto es una cosa")
    } else if juguete == "animal" {
        fmt.Println("El objeto es un animal")
    } else {
        fmt.Println("El objeto es otra categoria")
    }
}
func main() {
	// fmt.Println("texto de la funcion del ejecicio 2")
	// pensionados()
	// valor_cero()
	// Declaracion_corta_variables()
	// scoope()
	// puntero()
	// referencia_Copia()
	// fmt.Println("La altura es:", altura, "mts")
	// fmt.Println("La altura es:", equivalencia_en_pies(altura), " pies")
	// fmt.Println("La nueva altura es:", altura, "mts")
	// fmt.Println("La altura es:", altura2, "mts")
	// fmt.Println("Al envejecer:", ConversionEnPies(&altura), "mts")
	// fmt.Println("Despues de envejecer:", altura2, "mts")
	// fmt.Println("El area de un circulo cuyo radio es 3 es: ", area(3))
	// a, p := circulo(8)
	// fmt.Println("El area del circulo es ", a)
	// fmt.Println("El perimetro del circulo es ", p)
	// fmt.Println(sumar1(2))
	// fmt.Println(sumar1(2, 2))
	// fmt.Println(sumar1(2, 2, 2, 2, 2	, 2))	
	// ecodelamotaña("causa", 5)
	// area1, perimetro1 := circulo2(8)
	// fmt.Println("El area del circulo", area1())
	// fmt.Println("El perimetro del circulo", perimetro1())
	Objeto_Animal_Persona()

}

// 23,24, 25 ,26