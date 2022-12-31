package main

import (
	"fmt"
)

func main() {

	ejercicio1()
	ejercicio2()
	//ejercicio3()
	//ejercicio4()

}
func ejercicio1() {
	var nombre string = "Geral"
	fmt.Println(nombre)

	var direccion string = "casa"
	fmt.Println(direccion)

}

func ejercicio2() {
	var temperatura float32 = 21.0
	fmt.Println("temperatura de tulua:", temperatura)

	var humedad float32 = 0.88
	fmt.Println("Humedad de tulua:", humedad)

	var presion float32 = 1022.0
	fmt.Println("presion de tulua:", presion)
}

/*func ejercicio3(){

	//var 1nombre string    				//inicia con numero
	var apellido string 					//Correcta
	var int edad 							//int al final
	//1apellido := 6 						//inicia con numero
	var licencia_de_conducir = true			//corre pero no sige standar camaelCase
	//var estatura de la persona int 			// variable debe ser unsa sola Palabra
	cantidadDeHijos := 2					//correcta , solo local

}

func ejercicio4(){


	var apellido string = "gomez"			//Correcta
	var edad int = 35
	//var edad int = "35"						//no coincide tipo con valor asignado =>

	boolean := false
	//boolean := "false";

	var sueldo string = "45847.90"
	var sueldo float32= 45847.90
	//var sueldo string = 45847.90

	var nombre string = "Julian"
}*/
