package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	pb "github.com/FranZk9/lab3Distribuidos/proto"
	"google.golang.org/grpc"
)

var diccionario_relojes = make(map[string][3]int)
var diccionario_valores = make(map[string]string)

func main() {
	patron, compileErr := regexp.Compile("^(GetSoldados)")
	if compileErr != nil {
		fmt.Println("Error de compilacion de regex: ", compileErr)
	}

	port := ":50051"
	connS, err := grpc.Dial("dist085"+port, grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	defer connS.Close()
	serviceCliente := pb.NewMessageServiceClient(connS)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Utilizar 'GetSoldados <nombre sector> <nombre base>' para realizar la consulta")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		match := patron.MatchString(text)
		if match {
			input := strings.Split(text, " ")
			nombre_sector := input[1]
			nombre_base := input[2]
			key_diccionario := nombre_sector + "-" + nombre_base
			clock_stored, exist := diccionario_relojes[key_diccionario]
			if exist {
				fmt.Println("Valor si esta en memoria, aplicando Monotonic Reads")
				res, err := serviceCliente.GetSoldados(context.Background(),
					&pb.CommandGet{
						NombreSector: nombre_sector,
						NombreBase:   nombre_base,
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				res_split := strings.Split(res.Body, "-")
				clock_split := strings.Split(res_split[1], "/")
				clock_split_1, err := strconv.Atoi(clock_split[0])
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				clock_split_2, err := strconv.Atoi(clock_split[1])
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				clock_split_3, err := strconv.Atoi(clock_split[2])
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				var vectortemp [3]int = [3]int{clock_split_1, clock_split_2, clock_split_3}
				comp_vector := vector_mayor(clock_stored, vectortemp)
				if comp_vector {
					fmt.Print("Numero de soldados en el sector: " + nombre_sector + " Base: " + nombre_base + " es de " + diccionario_valores[key_diccionario] + " Reloj vectorial: ")
					fmt.Print(clock_stored)
					fmt.Println(" Obtenido desde el " + res_split[2] + "(Se aplicó Monotonic Reads)")
				} else {
					fmt.Print("Numero de soldados en el sector: " + nombre_sector + " Base: " + nombre_base + " es de " + res_split[0] + " Reloj vectorial: ")
					fmt.Print(vectortemp)
					fmt.Println(" Obtenido desde el " + res_split[2])
					diccionario_relojes[key_diccionario] = vectortemp
					diccionario_valores[key_diccionario] = res_split[0]
				}
			} else {
				fmt.Println("Valor no estaba en memoria")
				res, err := serviceCliente.GetSoldados(context.Background(),
					&pb.CommandGet{
						NombreSector: nombre_sector,
						NombreBase:   nombre_base,
					})
				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				res_split := strings.Split(res.Body, "-")
				clock_split := strings.Split(res_split[1], "/")
				clock_split_1, err := strconv.Atoi(clock_split[0])
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				clock_split_2, err := strconv.Atoi(clock_split[1])
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				clock_split_3, err := strconv.Atoi(clock_split[2])
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				var vectortemp [3]int = [3]int{clock_split_1, clock_split_2, clock_split_3}
				if res_split[0] == "No se encontro ningun valor que haga match" {
					fmt.Println(res_split[0])
					continue
				} else {
					fmt.Print("Numero de soldados en el sector: " + nombre_sector + " Base: " + nombre_base + " es de " + res_split[0] + " Reloj vectorial: ")
					fmt.Print(vectortemp)
					fmt.Println(" Obtenido desde el " + res_split[2])
					diccionario_relojes[key_diccionario] = vectortemp
					diccionario_valores[key_diccionario] = res_split[0]
				}
			}
		} else {
			fmt.Println("Comando inválido")
		}
	}
}

func vector_mayor(vector1 [3]int, vector2 [3]int) bool {
	var suma1 = 0
	var suma2 = 0
	for i := 0; i < 3; i++ {
		suma1 += vector1[i]
		suma2 += vector2[i]
	}
	if suma1 >= suma2 {
		return true
	} else {
		return false
	}
}
