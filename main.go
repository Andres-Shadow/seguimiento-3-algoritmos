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
		break
	case "3":
		break
	case "4":
		break
	case "5":
		break
	case "6":
		break
	case "7":
		break
	case "x":
		herramientas.GenearMatricesPorTamanio(tam, "datos1.dat", "datos2.dat")
		break
	}
	//rows, cols := 1024, 1024
	//matrix := generateRandomMatrix(rows, cols)
	//saveMatrixToFile(matrix, "datos.dat")
	//m := herramientas.GenerateRandomMatrix(1024, 2014)
	//herramientas.SaveMatrixToFile(m, "datos2.dat")
	//algoritmos.LlamarNaivStandart()
}
