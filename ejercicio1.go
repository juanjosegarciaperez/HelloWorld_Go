package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func IMPRIMIR() {
	fmt.Println("texto desde la funcion imprimir")
}

func Hola_string(s string) string {
	return "Hola " + s
}

func sumar(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("a y b son iguales")
	} else {
		fmt.Println("a y b son diferentes")
	}
}

func tipos_datos() {
	var texto string = "Juan"
	var numero int = 1000
	var decimal float64 = 0.00001

	fmt.Println(reflect.TypeOf(texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))
}

func arreglo() {
	var aprendices [5]string

	aprendices[0] = "Leonardo"
	aprendices[1] = "Juan Sebastian"
	aprendices[2] = "Frank"
	aprendices[3] = "Juan Jose"
	aprendices[4] = "Jaider"

	fmt.Println(aprendices[4])
}

func conver_string_to_boolean() {
	var palabra string = "true"

	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Println("Este es el error", err)
}

func conver_bool_to_string() {
	var palabra_bool bool = true

	StrBool := strconv.FormatBool(palabra_bool)
	fmt.Println(StrBool, reflect.TypeOf(StrBool))
}
func main() {
	// IMPRIMIR()
	//fmt.Println("texto desde la funcion main")
	// fmt.Println(Hola_string("jaime"))
	// fmt.Println(sumar(3, 4))
	// comparar_bool()
	// arreglo()
	// tipos_datos()
	// conver_string_to_boolean()
	// conver_bool_to_string()
}

// ejercicios 15, 16, 17, d18, d19, d20, 22
