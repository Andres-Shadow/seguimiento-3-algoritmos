package main

import (
	"fmt"
	"os"
	"seg_3_algoritmos/algoritmos"
	"seg_3_algoritmos/herramientas"
	"strconv"
)

func main() {

	opcion := os.Args[1]
	var tam int

	if len(os.Args) < 3 {
		fmt.Println("1. primer tamaño")
		fmt.Println("2. segundo tamaño")
		fmt.Println("3. tercer tamaño")
		fmt.Print("Seleccione el tamaño para el arreglo: ")
		fmt.Scan(&tam)
	} else {
		tam, _ = strconv.Atoi(os.Args[2])
	}

	switch opcion {
	case "1":
		algoritmos.LlamarNaivStandart(tam)
		break
	case "2":
		algoritmos.LlamarNaivOnArray(tam)
		break
	case "3":
		algoritmos.LlamarNaivKahan(tam)
		break
	case "4":
		algoritmos.LlamarNaivLoopUnrollingTwo(tam)
		break
	case "5":
		algoritmos.LlamarNaivLoopUnrollingThree(tam)
		break
	case "6":
		algoritmos.LlamarNaivLoopUnrollingFour(tam)
		break
	case "7":
		algoritmos.LlamarStrassenNaiv(tam)
		break
	case "x":
		herramientas.GenearMatricesPorTamanio(tam)
		break
	}
}
